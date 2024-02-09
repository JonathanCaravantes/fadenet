
FROM golang:1.22

WORKDIR /app
COPY . .



RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /webserver

EXPOSE 9090

CMD ["/webserver"]