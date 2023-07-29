FROM golang:1.19-alpine
USER nobody
WORKDIR /tabrss
COPY . .
RUN go build -o tabrss .
CMD ["./tabrss"]