FROM scratch
COPY go-contact-service ./
COPY config.ini ./
EXPOSE 1326

ENTRYPOINT ["/go-echo-bun-crud-service"]