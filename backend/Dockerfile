FROM golang:1.20.2-alpine3.17

RUN apk --update add redis

WORKDIR /otd

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 4000

RUN go build -o app
RUN chmod +x app

CMD redis-server & ./app