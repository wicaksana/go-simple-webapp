go mod init wicaksana.id/hello

docker build -t go-simple-webapp .

docker run -p 8900:8900 go-simple-webapp

curl localhost:8900/hello