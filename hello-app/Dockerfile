FROM golang:1.14.9-alpine AS builder

RUN mkdir /build
ADD go.mod go.sum hello.go /build/
workdir /build	

RUN go build

FROM alpine
run adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/helloworld /app/
WORKDIR /app
CMD ["./helloworld"]
