FROM golang:1.13

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shh cmd/*.go

EXPOSE 8080

ENTRYPOINT ["/app/shh"]