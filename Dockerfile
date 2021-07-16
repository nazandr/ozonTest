FROM golang:alpine AS builder
RUN apk add --no-cache git

WORKDIR /app/

COPY . .
RUN go mod download

RUN CGO_ENABLED=0

RUN go build -o ./out/shortUrl ./cmd/shortUrl/

FROM alpine:3.9 

COPY --from=builder /app/out/shortUrl /app/shortUrl
COPY ./configs /configs

EXPOSE 8080

CMD ["/app/shortUrl"]
