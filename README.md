# go-meater-meter

## Configuration
### Meater API Authentication (DRAFT)
```ini
email = <username>
password = <password>
```
### Postgres Authentication (DRAFT)
```ini
server = <server address>
port = <server_port>
username = <username>
password = <password>
database_name = <database_name>
```

## References
### API Spec
https://github.com/apption-labs/meater-cloud-public-rest-api
### lib/pq
https://github.com/lib/pq/blob/master/example/listen/doc.go
#### Connection string format
`postgresql://<username>@<server>:<port>/<database_name>`