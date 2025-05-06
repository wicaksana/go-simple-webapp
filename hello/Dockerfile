FROM golang:1.23.4 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
# CGO_ENABLED=0 produces a static binary suitable for scratch/alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8900
CMD ["/app/server"]