FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN go build -o server server/server.go

CMD ["./server"]
