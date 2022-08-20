# build stage
FROM golang:1.19-alpine3.15 AS builder 
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz
RUN chmod +x start.sh

# run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main . 
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY db/migrations ./migrations

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]