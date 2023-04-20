# Link shortener

### Run
`docker-compose up`

### Test with coverage
`make test-report`

### gRPC schema

See `internal/delivery/grpc/v1/schema/shortener.proto`

### HTTP schema

`POST /create` - create new shortening

Request:
```json
{
  "url": "https://github.com/"
}
```

Response:
```json
{
  "short": "QZIDeu54HA"
}
```


`GET /restore?shortened={short}` - get shortening by URL

Response:
```json
{
  "original": "https://example.org/"
}
```




