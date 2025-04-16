# High-Performance TCP Server

A high-performance TCP server built with Go and gnet, supporting high-concurrency connections.

## Features

- High-performance TCP server using gnet
- Multi-core support
- Docker containerization
- Automated CI/CD pipeline with GitHub Actions
- Auto-deployment to VPS

## Local Development

1. Install Go 1.21 or later
2. Clone the repository
3. Install dependencies:
```bash
go mod download
```
4. Run the server:
```bash
go run main.go -port 9000
```

## Docker Deployment

Build and run with Docker:

```bash
docker build -t mmo-go-server .
docker run -p 9000:9000 mmo-go-server
```

## CI/CD Setup

To enable CI/CD pipeline, set the following secrets in your GitHub repository:

- `DOCKER_USERNAME`: Your Docker Hub username
- `DOCKER_PASSWORD`: Your Docker Hub password
- `VPS_HOST`: Your VPS IP address
- `VPS_USERNAME`: SSH username for VPS
- `VPS_SSH_KEY`: SSH private key for VPS access

## Testing

To test the TCP server, you can use tools like netcat:

```bash
echo "Hello" | nc localhost 9000
```

## Performance Testing

For performance testing, you can use tools like hey or wrk:

```bash
wrk -t12 -c400 -d30s tcp://localhost:9000
``` 