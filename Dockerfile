FROM golang:1.25.6 AS bulilder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:stable-slim

WORKDIR /app

COPY --from=bulilder /app/main .

EXPOSE 8080

CMD ["./main"]