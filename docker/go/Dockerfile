FROM golang:1.15.5-alpine3.12

RUN apk --no-cache add htop git

RUN addgroup -S www-data && adduser -S www-data -G www-data

COPY --chown=www-data . /go/src/messenger

WORKDIR /go/src/messenger

RUN go build -o /bin/messenger && chown www-data:www-data /bin/messenger

USER www-data

CMD /bin/messenger server