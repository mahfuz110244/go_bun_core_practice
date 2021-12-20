FROM scratch
COPY go-echo-bun-crud-service ./
COPY config.ini ./
EXPOSE 3000

ENTRYPOINT ["./go-echo-bun-crud-service"]