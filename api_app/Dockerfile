FROM golang:alpine3.12 as debug


ENV GO111MODULE=on

# installing git
RUN apk update && apk upgrade && \
  apk add --no-cache git \
  dpkg \
  gcc \
  git \
  musl-dev

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH



WORKDIR /go/src/work
COPY ./src /go/src/work/

ENV GOPROXY https://proxy.golang.org
RUN go mod download


RUN go build -o app

###########START NEW IMAGE###################

FROM alpine:3.9 as prod
COPY --from=debug /go/src/work/app /
CMD ./app