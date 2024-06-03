# build-go image
FROM golang:bullseye as build-go
WORKDIR /build

RUN apt-get update && apt-get install --assume-yes upx
COPY . /build
RUN make go-build
RUN upx --best --lzma us-soccer-go

# runtime image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN mkdir /app
RUN mkdir /app/bin
RUN mkdir /data
COPY --from=build-go /build/us-soccer-go /app

# runtime params
WORKDIR /app
ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

EXPOSE 8080

CMD ["./us-soccer-go"]