package main

import (
	"context"
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/ajski1701/go-meater-meter/authentication"
	"github.com/ajski1701/go-meater-meter/config"
	"github.com/ajski1701/go-meater-meter/devices"
	"github.com/ajski1701/go-meater-meter/influxdb"
	"github.com/ajski1701/go-meater-meter/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/klog"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	client *clientset.Clientset
)

func getAuthenticationToken() string {
	authenticationDebug := config.LoadConfigIni().Section("debug").Key("disable_authentication").String()
	sessionToken := ""
	if authenticationDebug == "true" {
		sessionToken = ""
	} else {
		sessionToken = authentication.GetAuth(config.LoadConfigIni())
	}
	return sessionToken
}

func getPollRate() int {
	pollRateStr := config.LoadConfigIni().Section("app-config").Key("poll_rate").String()
	pollRateInt, err := strconv.Atoi(pollRateStr)
	if err != nil {
		panic(err)
	}
	return pollRateInt
}

func submitInfluxData(devices []models.Devices, sessionToken string, client influxdb2.Client) {
	for _, device := range devices {
		tags := map[string]string{
			"device_id": device.Id,
			"cook_id":   device.Cook.Id,
			"cook_name": device.Cook.Name,
		}
		fields := map[string]interface{}{
			"device_internal_temperature": device.Temperature.Internal,
			"device_ambient_temperature":  device.Temperature.Ambient,
			"cook_target_temperature":     device.Cook.Temperature.Target,
			"cook_peak_temperature":       device.Cook.Temperature.Peak,
			"cook_elapsed_time":           device.Cook.Time.Elapsed,
			"cook_remaining_time":         device.Cook.Time.Remaining,
			"cook_state":                  device.Cook.State,
			"updated_at":                  device.Updated_At,
		}
		klog.Info("Writing influxdb data for ", device.Id)
		go influxdb.WriteData(client, tags, fields)
	}
}

func getAndSubmit() {
	for {
		klog.Info("Authenticating to Meater Cloud API")
		sessionToken := getAuthenticationToken()
		klog.Info("Querying Meater Cloud Device API")
		devices := devices.GetDevices(sessionToken)
		influxdbClient := influxdb.GetInfluxClient()
		pollRate := getPollRate()

		submitInfluxData(devices, sessionToken, influxdbClient)
		time.Sleep(time.Duration(pollRate) * time.Second)
	}
}

func main() {
	klog.Info("Starting go-meater-meter application")

	var (
		leaseLockName      string
		leaseLockNamespace string
		podName            = os.Getenv("POD_NAME")
	)
	flag.StringVar(&leaseLockName, "lease-name", "", "Name of lease lock")
	flag.StringVar(&leaseLockNamespace, "lease-namespace", "default", "Name of lease lock namespace")
	flag.Parse()

	if leaseLockName == "" {
		klog.Fatal("Missing lease-name flag")
	}
	if leaseLockNamespace == "" {
		klog.Fatal("Missing lease-namespace flag")
	}

	config, err := rest.InClusterConfig()
	client = clientset.NewForConfigOrDie(config)

	if err != nil {
		klog.Fatalf("Failed to get kubeconfig")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lock := getNewLock(leaseLockName, podName, leaseLockNamespace)
	runLeaderElection(lock, ctx, podName)
}

func getNewLock(lockname, podname, namespace string) *resourcelock.LeaseLock {
	return &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      lockname,
			Namespace: namespace,
		},
		Client: client.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: podname,
		},
	}
}

func runLeaderElection(lock *resourcelock.LeaseLock, ctx context.Context, id string) {
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: true,
		LeaseDuration:   15 * time.Second,
		RenewDeadline:   10 * time.Second,
		RetryPeriod:     2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(c context.Context) {
				go getAndSubmit()
			},
			OnStoppedLeading: func() {
				klog.Info("No longer the leader, staying inactive")
			},
			OnNewLeader: func(current_id string) {
				if current_id == id {
					klog.Info("Still the leader")
					return
				}
				klog.Info("Newly elected leader is ", current_id)
			},
		},
	})
}
