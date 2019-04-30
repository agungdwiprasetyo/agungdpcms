FROM golang:1.11.1-alpine3.8

WORKDIR /go/src/github.com/agungdwiprasetyo/agungdpcms

ENV SRC_DIR=/go/src/github.com/agungdwiprasetyo/agungdpcms

ENV BUILD_PACKAGES="git curl"

ADD . $SRC_DIR

RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && curl https://glide.sh/get | sh \
  && glide install \
  && apk del $BUILD_PACKAGES \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o bin .

COPY .env $SRC_DIR.env

EXPOSE 8000

ENTRYPOINT ["sh", "-c", "./bin"]