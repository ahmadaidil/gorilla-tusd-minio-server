FROM golang:alpine
ENV WORKDIR /go/src/github.com/ahmadaidil/gorilla-tusd-minio-server
ENV SERVER_PORT :8000
COPY . ${WORKDIR}
WORKDIR ${WORKDIR}
RUN apk add --no-cache git \
    && go get -d -v \
    && go build -o /go/bin/gorilla-tusd-minio-server \
    && rm -rf /go/src/* \
    && apk del git
EXPOSE 8000
ENTRYPOINT [ "/go/bin/gorilla-tusd-minio-server" ]