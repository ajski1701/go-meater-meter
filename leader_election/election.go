package leader_election

import "os"

func checkKubernetes() bool {
	podName := os.Getenv("POD_NAME")
	return len(podName) > 0
}
