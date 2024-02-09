
FROM golang:1.20.2

WORKDIR /app

COPY go.mod go.sum ./

ADD template static ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /webserver

EXPOSE 9090

CMD ["/webserver"]