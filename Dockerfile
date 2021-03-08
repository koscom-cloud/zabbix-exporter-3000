FROM golang

LABEL maintainer="rzrbld <razblade@gmail.com>"

ENV GOPATH /go
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org

WORKDIR /app
COPY ./ /app

RUN  go build main.go && cp main /go/bin/ze3000

FROM alpine

# EXPOSE 8080
RUN mkdir /app && chmod 777 /app
WORKDIR /app

COPY --from=0 /go/bin/ze3000 /app/ze3000

# RUN  \
#      apk add --no-cache ca-certificates 'curl>7.61.0' 'su-exec>=0.2' && \
#      echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

CMD ["/app/ze3000"]
