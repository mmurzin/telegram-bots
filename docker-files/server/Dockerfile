FROM golang:1.19 AS build-stage

WORKDIR /app
COPY .env ./
WORKDIR /app/go
COPY ./app/go.mod ./app/go.sum ./
COPY ./app/*.go ./
RUN go mod download

RUN CGO_ENABLED=0 COOS=linus go build -o telegram-bots


CMD ["./telegram-bots"]