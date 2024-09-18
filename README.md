# user-service

## Menjalankan aplikasi dengan docker

 - dev
	```
	$ docker-compose up -d --build
	```
 - prod
	```
	$ docker build -t user-service .
	$ docker run -d --name user-service -e <ENV> -p 3000:3000 user-service 
	```
	envoiremnts silahkan lihat di file docker-compose.yml

## Prerequisites
- Golang(>=1.20) - Download and Install [Golang](https://golang.org/)
```
  clone this project
  $ cd user-service
  $ go mod download
  $ go build 
  $ ./user-service
```
