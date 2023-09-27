FROM golang:1.21 AS build

RUN mkdir app
COPY .. /app

WORKDIR /app

RUN go mod download
RUN go mod verify

RUN go build -o server cmd/main.go

EXPOSE 3000/tcp

CMD ["./server"]