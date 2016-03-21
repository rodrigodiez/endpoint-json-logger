# endpoint-json-logger
Light web server that decodes and logs JSON requests. Useful for debugging webhooks

## Run it
```bash
docker run --rm -p "8000:8000" rodrigodiez/endpoint-json-logger
```

## Test it
```bash
curl -X POST -d '[{"username":"xyz","password":"xyz"}]'  localhost:8000
```
