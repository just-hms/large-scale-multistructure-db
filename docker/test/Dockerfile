FROM golang:1.19.3 as build
WORKDIR /app

RUN apt update && \
	apt upgrade -y

COPY go.mod go.sum ./
RUN go mod download && \
	go mod verify && \
	go mod tidy

CMD /wait

COPY . .
RUN mkdir -p build && \
	go build -o build ./... && \
	go test ./...
