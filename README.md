# go-meater-meter

## Configuration - config.ini
```ini
[api-authentication]
email = "<meater_cloud_email>"
password = "<meater_cloud_password>"

[app-config]
poll_rate = 30 #in seconds

[influxdb]
url = "<influxdb_url>"
token = "<influxdb_token>"
org = "<influxdb_org>"
bucket = "<influxdb_bucket>"
```

## References
### API Spec
https://github.com/apption-labs/meater-cloud-public-rest-api