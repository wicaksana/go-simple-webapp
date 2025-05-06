APPLICATION_NAME ?= go-hello-app
PORT ?= 8900

build:
	docker build -t ${APPLICATION_NAME} .
run:
	docker run -p ${PORT}:${PORT} -d --name ${APPLICATION_NAME} ${APPLICATION_NAME}
stop:
	docker stop ${APPLICATION_NAME}