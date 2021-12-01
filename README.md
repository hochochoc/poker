# Poker

## Prerequisites
- go 1.16

## Install dependencies
```go mod download```

## Development
```
# directly
go run main.go <hand1> <hand2>

# docker
docker build -t <image_name> .
docker run <image_name> <hand1> <hand2>
```
## Test
```go test .```
