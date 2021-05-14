FROM golang:1.16 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app
ENV CGO_ENABLED=0
RUN CGO_ENABLE=0 GOOS=linux go build -o app cmd/server/main.go
FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]

## Use below command to build docker image
# docker build -t prod-ready-api .