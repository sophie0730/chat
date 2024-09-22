FROM golang:1.23-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN cd cmd && go build -o main .

FROM alpine:latest
WORKDIR /chat
COPY --from=builder /build/cmd/main .
EXPOSE 8000
CMD ["./main"]