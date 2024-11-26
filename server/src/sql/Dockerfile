FROM alpine:3.17.0

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.23.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /migration
ADD ./schema ./schema
ADD goose.sh .
RUN chmod +x ./goose.sh

ENTRYPOINT ["./goose.sh"]