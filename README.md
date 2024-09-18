# chi-onion-arch
Boilerplate aplikasi go dengan framework chi, dengan pendekatan clean/onion architecture

## Menjalankan aplikasi
```
$ cd chi-onion-arch
$ go run main.go
```
## Menjalankan aplikasi dengan docker

 - dev
	```
	$ docker-compose up -d --build
	```
 - prod
	```
	$ docker build -t chi-onion-arch .
	$ docker run -d --name chi-onion-arch -e <ENV> -p 3000:3000 chi-onion-arch 
	```
	envoiremnts silahkan lihat di file docker-compose.yml

## Prerequisites
- Golang(>=1.20) - Download and Install [Golang](https://golang.org/)
```
  clone this project
  $ cd chi-onion-arch
  $ go mod download
  $ go build 
  $ ./chi-onion-arch
```
