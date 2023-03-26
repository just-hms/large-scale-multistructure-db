FROM golang:1.19.3 as builder
WORKDIR /app

RUN apt update && \
	apt upgrade -y

COPY go.mod go.sum ./
RUN go mod download && \
	go mod verify && \
	go mod tidy

COPY . .	
RUN mkdir -p build && \
	go build -o build ./...

FROM ubuntu:22.04
WORKDIR /app

RUN apt update && \
	apt upgrade -y && \
	apt-get install -y netcat

COPY --from=builder /app/build/cmd .
COPY --from=builder /app/docker/deploy/be/entrypoint.sh .

ENV GIN_MODE=release

CMD ["sh", "./entrypoint.sh"]
