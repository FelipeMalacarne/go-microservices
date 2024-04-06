FROM alpine:3.8

WORKDIR /app

COPY ./mailApp /app

CMD ["/app/mailApp"]
