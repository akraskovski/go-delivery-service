FROM golang:1.14.1 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN make build

FROM alpine AS production
COPY --from=builder /app/main /
COPY --from=builder /app/configs /

EXPOSE 8000

ENTRYPOINT ["./main", "-config-path", "deliveryservice.toml"]
