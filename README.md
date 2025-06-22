# SimpleTimeService

## Introduction
SimpleTimeService is a minimal microservice built in Golang using the Gin web framework.  
When accessed via the `/` route, it returns a JSON response with the `current timestamp` and the `IP address of the visitor`.

The container is designed to run securely using a `non-root user` and requires `no special setup` to build or run.

## Run Using Docker
### Step 1: Pull the Image
```
docker pull shilpa2704/particle41:v1.0
```

### Step 2: Run the Container
```
docker run -p 8080:8080 shilpa2704/particle41:v1.0
```
You can access the service at
http://localhost:8080

## Run Using Go (Locally)
```
go run cmd/main.go
```
This starts the server on localhost:8080

## Sample Response
```
{
  "timestamp": "2025-06-22T15:21:44+08:00",
  "ip": "127.0.0.1"
}
```
