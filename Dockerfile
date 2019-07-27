FROM golang:1.12.7-alpine3.10

ARG APP_NAME=agungdpcms
ARG LOG_DIR=/${APP_NAME}/logs
ARG BUILD_PACKAGES="git curl"

RUN mkdir -p ${LOG_DIR}

WORKDIR /usr/app

ENV SRC_DIR=/usr/app/

ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

COPY . $SRC_DIR

RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && go mod download \
  && apk del $BUILD_PACKAGES \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o bin .

EXPOSE 8000

ENTRYPOINT ["sh", "-c", "./bin"]