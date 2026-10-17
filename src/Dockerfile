# 1. Docker file directory 
FROM ubuntu:latest
FROM golang:1.26-alpine AS builder
# 2. Dependecy copy
WORKDIR /app 
COPY go.mod ./
# 3. Dependecy download 
RUN go mod download
# 4.directory code copy 
COPY . . 
# 5. Build
RUN go build -o main .
# 6. Run
EXPOSE 8000
CMD ["./main"]


