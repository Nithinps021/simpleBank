#build stage
FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN apk add curl
RUN apk add make
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz 
RUN go build -o main main.go 

#final image
FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY --from=builder /app/db/ ./db
COPY ./app.env .

RUN mv migrate /usr/bin

CMD ["/app/main"]