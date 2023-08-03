FROM golang:1.20 as builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd


FROM scratch as release
LABEL author="ghilbut@gmail.com"

EXPOSE 8080 9095

ENV GOMAXPROCS=1

COPY --from=builder /app/server /usr/local/bin/

CMD ["server"]
