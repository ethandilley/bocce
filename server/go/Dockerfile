FROM golang:1.20-alpine as builder

WORKDIR /app

# TODO: Change if we have unit tests
COPY . .

RUN go mod tidy
RUN go build -o bocce-server .

FROM alpine:latest

COPY --from=builder /app/bocce-server /usr/local/bin/

EXPOSE 8000

CMD ["bocce-server"]
