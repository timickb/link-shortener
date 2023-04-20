# Link shortener

### Storage type flag
`app --storage=postgres` - uses PostgreSQL as a storage

`app --storage=memory` - uses in-memory storage

### Pull Docker image
`docker pull timickb/linkshortener`

The image is rebuilt and pushed on each `master` commit

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




