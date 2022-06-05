# go-meater-meter

## Configuration - config.ini
```ini
[api-authentication]
email = "<meater_cloud_password>"
password = "<meater_cloud_password>"

[app-config]
poll_rate = 30 #in seconds

[influxdb]
url = "<influxdb_url>"
token = "<influxdb_token>" #INFLUXDB_TOKEN
org = "<org>"
bucket = "<bucket>"
```

## References
### API Spec
https://github.com/apption-labs/meater-cloud-public-rest-api