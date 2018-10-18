# Gorilla Tusd Minio Server

### Upload file server with gorilla mux, tusd and minio store
---
Get / clone
```cmd
$ go get github.com/ahmadaidil/gorilla-tusd-minio-server
$ cd $GOPATH/src/github.com/gorilla-tusd-minio-server
```
---
Get Dependencies:
```cmd
$ go get -d -v
```
---
Pull docker minio and run
```
$ docker pull minio/minio
$ docker run -p 9000:9000 \
  -e "MINIO_ACCESS_KEY=MINIO_ACCESS" \
  -e "MINIO_SECRET_KEY=MINIO_SECRET" \
  minio/minio server /data
```
Output
```Output
Endpoint:  http://172.17.0.2:9000  http://127.0.0.1:9000
AccessKey: MINIO_ACCESS 
SecretKey: MINIO_SECRET 

Browser Access:
   http://172.17.0.2:9000  http://127.0.0.1:9000

Command-line Access: https://docs.minio.io/docs/minio-client-quickstart-guide
   $ mc config host add myminio http://172.17.0.2:9000 MINIO_ACCESS MINIO_SECRET
```
* Copy AccessKey, SecretKey and Endpoint.
* Go to http://127.0.0.1:9000, login with AccessKey and SecretKey.
* Create bucket with name: "my-bucket" or anything else.
---
Run
```cmd
$ go run main.go
```
---
Build and run
```cmd
$ go build
$ ./gorilla-tusd-minio-server
```
---
Docker build
```cmd
docker build -t <image-name:tag> .
```
---
&copy; 2018 Ahmad Aidil