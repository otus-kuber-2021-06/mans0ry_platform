FROM golang:1.15-alpine as builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base
COPY . /src
WORKDIR /src

RUN go build -o /go/bin/go-server .

FROM alpine as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
RUN mkdir /app
COPY ./homework.html /app/.
WORKDIR /app
COPY --from=builder /go/bin/go-server /src/go-server

RUN addgroup -S -g 1001 developer && adduser -S developer -G developer --uid 1001
RUN chown -Rvc developer:developer /app
USER developer

EXPOSE 8000
ENTRYPOINT ["/src/go-server"]