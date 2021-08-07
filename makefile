run:
	go run .

build:
	go build . && ./go-rest-sample


docker-build:
	docker build -t droidx/gorestsample .


docker-sync:
	docker build -t droidx/gorestsample .
	docker push droidx/gorestsample

