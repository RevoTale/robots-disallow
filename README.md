# robots-disallow
Docker container to response `User-agent: * Disallow: /` to each request.

## Features
- Serves `User-agent: * Disallow: /` robots.txt content for all requests except `/status`
- Provides a `/status` endpoint that returns service information in JSON format

## Endpoints
- `GET /` (or any path except `/status`) - Returns robots.txt content
- `GET /status` - Returns service status and information in JSON format

## Usage
```bash
# Run with Docker
docker run -p 8080:80 ghcr.io/revotale/robots-disallow

# Build and run locally
go build
./robots-disallow
```
