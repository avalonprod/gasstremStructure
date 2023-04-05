FROM golang:1.19

EXPOSE 8000

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build src/cmd/app/main.go

CMD ["./main"]