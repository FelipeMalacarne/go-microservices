FROM alpine:latest

WORKDIR /app

COPY ./authApp /app

CMD ["/app/authApp"]
