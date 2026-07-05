FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o go-htmx ./cmd/go-htmx


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-htmx .
COPY --from=builder /app/cmd/go-htmx/index.html .

EXPOSE 8000

CMD ["./go-htmx"]