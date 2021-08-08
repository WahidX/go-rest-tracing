run:
	go run .

build:
	go build . && ./go-rest-sample


docker-build:
	docker build -t droidx/gorestsample .


docker-sync:
	docker build -t droidx/gorestsample .
	docker push droidx/gorestsample
	
watch:
	$(eval PACKAGE_NAME=$(shell head -n 1 go.mod | cut -d ' ' -f2))
	docker run -it --rm -w /go/src/$(PACKAGE_NAME) -v $(shell pwd):/go/src/$(PACKAGE_NAME) -p 8000:8000 cosmtrek/air
