FROM golang:1.19.3 as build
WORKDIR /app

RUN apt update && \
	apt upgrade -y && \
	apt-get install -y netcat

COPY go.mod go.sum ./
RUN go mod download && \
	go mod verify && \
	go mod tidy

COPY . .
RUN mkdir -p build && \
	go build -o build ./...


CMD ["./docker/test/entrypoint.sh"]
