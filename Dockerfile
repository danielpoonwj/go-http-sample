# builder stage
FROM golang:1.12-alpine3.9 AS builder

# Install go get dependencies
RUN apk add --no-cache git

WORKDIR /src
COPY . .

RUN go build -o ./build/server

# application stage
FROM alpine:3.9

WORKDIR /app
COPY --from=builder /src/build/server ./server
CMD ["./server"]
