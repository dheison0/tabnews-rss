FROM golang:1.19-alpine
RUN apk update && apk add gcc musl-dev
WORKDIR /tabrss
COPY . .
RUN go build -o tabrss .
CMD ["./tabrss"]