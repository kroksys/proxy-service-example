# Proxy Service
Very simple HTTP proxy which forwards all the requests to a given domain.

Logs to file all requests and responses that go through it.

## Usage

### Help
```bash
go run . proxy -h
```

### Run with defaults
```bash
go run . proxy
```

### Run with CLI flags
```bash
go run . proxy --listen 0.0.0.0:5000 --target http://httpforever.com --log anotherlog.txt
```

### Run with config
```bash
go run . proxy -c config.yaml
```

## Test
I tested only proxy part, but it would be great to test config file reading and loging to file as an improvement.
```bash
go test ./src/proxy
```