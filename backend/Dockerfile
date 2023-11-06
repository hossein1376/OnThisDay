FROM golang:1.21.3-alpine3.18

WORKDIR /otd

COPY go.mod go.sum ./
RUN go mod download

COPY . .
EXPOSE 4000

RUN go build -o app && chmod +x app

CMD ./app