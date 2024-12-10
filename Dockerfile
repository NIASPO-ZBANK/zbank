FROM golang:1.23.1

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o zbank ./cmd/zbank/zbank.go

EXPOSE 8080

CMD ["./zbank"]
