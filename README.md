# endpoint-json-logger
Light web server that decodes JSON payloads from requests and logs them. Useful for debugging webhooks

## Run it
```bash
docker run --rm -p "8000:8000" rodrigodiez/endpoint-json-logger
```

## Test it
```bash
curl -X POST -d '[{"username":"xyz","password":"xyz"}]'  localhost:8000
```
