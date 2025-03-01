FROM golang:1.23.3

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

WORKDIR /app

RUN go build -o /main main.go

CMD ["/main"]