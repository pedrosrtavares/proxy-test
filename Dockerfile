FROM golang:alpine as builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]
