ARG GOLANG_TAG
FROM golang:${GOLANG_TAG} as debug


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



ARG APP_NAME
WORKDIR /go/src/${APP_NAME}


COPY ./app .
RUN ls

ENV GOPROXY https://proxy.golang.org
RUN go mod download




RUN go build -o app


