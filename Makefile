
IMAGE := promotion_api

run:
	go run main.go

dockerize:
	@docker build -t ${IMAGE} .	

run-container:
	@docker run -it -d --net=host  ${IMAGE}
	
