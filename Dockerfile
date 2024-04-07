FROM golang:1.20-alpine

ENV DATABASE_HOST=localhost
ENV DATABASE_USER=postgres
ENV DATABASE_PASSWORD=gemmq123456
ENV DATABASE_NAME=kopoksu
ENV DATABASE_PORT=5432

ENV ADMIN_EMAIL=admin@kangpoksu.com
ENV ADMIN_PASSWORD=adminkangpoksu

ENV MASTER_EMAIL=master@kangpoksu.com
ENV MASTER_PASSWORD=masterkangpoksu

ADD . /go/src/app
WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

CMD go run cmd/kopoksu/kopoksu.go
