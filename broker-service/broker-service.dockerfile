FROM alpine:latest

WORKDIR /app

COPY ./brokerApp /app

CMD ["/app/brokerApp"]
