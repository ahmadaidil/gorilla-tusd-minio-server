FROM golang:alpine
ENV WORKDIR /go/src/github.com/ahmadaidil/gorilla-tusd-minio-server
ENV SERVER_PORT :3000
COPY . ${WORKDIR}
WORKDIR ${WORKDIR}
RUN apk add --no-cache git \
    && go get -d -v \
    && go build -o /go/bin/gorilla-tusd-minio-server ./main.go \
    && mkdir -p /srv/tusd-hooks \
    && mkdir -p /srv/tusd-data \
    && mv template /srv/tusd-data \
    && rm -r /go/src/* \
    && apk del git
WORKDIR /srv/tusd-data
EXPOSE 3000
ENTRYPOINT [ "/go/bin/gorilla-tusd-minio-server", "--hooks-dir", "/srv/tusd-hooks" ]