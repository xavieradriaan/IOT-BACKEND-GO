FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app main.go

FROM golang:1.23
WORKDIR /app
RUN apt-get update && apt-get install -y sqlite3
COPY --from=builder /app/app .
COPY .env .env
COPY init_users.sql init_users.sql

EXPOSE 8000
CMD ["sh", "-c", "sqlite3 /app/users.db < /app/init_users.sql && exec ./app"]
