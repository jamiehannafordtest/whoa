default: image

whoa: whoa.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

image: whoa Dockerfile
	docker build -t rackerlabs/whoa .

upload: image
	docker push rackerlabs/whoa
