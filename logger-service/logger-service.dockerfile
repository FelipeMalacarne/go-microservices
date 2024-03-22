FROM alpine:3.8

WORKDIR /app

COPY ./loggerApp /app

CMD ["/app/loggerApp"]
