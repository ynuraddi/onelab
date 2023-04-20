FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o app ./cmd

FROM alpine
WORKDIR /app
COPY --from=builder /app/app .

LABEL maintainers = "ynuraddi"
LABEL version = "1.0"

CMD ["./app"]