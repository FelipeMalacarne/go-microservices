FROM alpine:3.8

WORKDIR /app

COPY ./mailApp /app
COPY templates /app/templates

CMD ["/app/mailApp"]
