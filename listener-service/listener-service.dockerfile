FROM alpine:latest

WORKDIR /app

COPY ./listenerApp /app

CMD ["/app/listenerApp"]
