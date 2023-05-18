FROM golang:1.19

WORKDIR /app

COPY . .

RUN go build -o bin/mano-server ../cmd/main.go

EXPOSE 9090

CMD ["./mano-server"]
