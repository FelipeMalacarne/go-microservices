FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /app/frontendApp ./cmd/web 

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/frontendApp /app/frontendApp

CMD ["/app/frontendApp"]
