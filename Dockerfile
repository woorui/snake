FROM golang:1.12-alpine AS builder

WORKDIR /usr/src/app

COPY . .
RUN go build -v

FROM alpine:3.5

RUN apk --no-cache add ca-certificates

WORKDIR /usr/local/bin

COPY --from=builder /usr/src/app/snake .

CMD ["./snake"]