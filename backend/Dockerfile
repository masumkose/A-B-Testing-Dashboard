
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ab-testing-backend .


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/ab-testing-backend .

EXPOSE 8080


CMD [ "./ab-testing-backend" ]

