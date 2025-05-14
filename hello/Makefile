DOCKERHUB_USERNAME ?= muarwi
APPLICATION_NAME ?= go-hello-app
PORT ?= 8900

build:
	docker build -t ${DOCKERHUB_USERNAME}/${APPLICATION_NAME} .
run:
	docker run -p ${PORT}:${PORT} -d --name ${APPLICATION_NAME} ${DOCKERHUB_USERNAME}/${APPLICATION_NAME}
stop:
	docker stop ${APPLICATION_NAME}
push:
	docker login -u ${DOCKERHUB_USERNAME}
	docker push ${DOCKERHUB_USERNAME}/${APPLICATION_NAME}