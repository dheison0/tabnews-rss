FROM golang:1.19-alpine
WORKDIR /tabrss
COPY . .
RUN go build -o tabrss .
CMD ["./tabrss"]