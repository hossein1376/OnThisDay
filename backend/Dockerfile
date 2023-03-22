FROM golang:1.20.2-alpine3.17

WORKDIR /otd

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 4000

RUN go build -o app
RUN chmod +x app

CMD ["./app"]