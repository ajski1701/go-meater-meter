# go-meater-meter

## Configuration
### Container Environment Variables
| Environment Variable | Description | Default |
| ----------- | ----------- | ----------- |
| CONFIG_FILE      | Location of the application's config file       | /app/config.ini |
### config.ini
```ini
[api-authentication]
email = "<meater_cloud_email>" #OPTIONAL IF USING TOKEN AUTHENICATION
password = "<meater_cloud_password>" #OPTIONAL IF USING TOKEN AUTHENICATION
token = "<meater_cloud_api_token>" #EMAIL/PASSWORD OPTIONS ARE IGNORED IF USED

[app-config]
poll_rate = 30 #in seconds

[influxdb]
url = "<influxdb_url>"
token = "<influxdb_token>"
org = "<influxdb_org>"
bucket = "<influxdb_bucket>"
```

## Running in Docker
```
docker run -dit --name meater-meter -v <configFilePath>\config.ini:/app/config.ini quay.io/aj1701/go-meater-meter:latest
```

## References
### API Spec
https://github.com/apption-labs/meater-cloud-public-rest-api
