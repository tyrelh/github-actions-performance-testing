build:
	export GOOS=linux && \
	export GOARCH=amd64 && \
	export CGO_ENABLED=0 && \
	go build -o test-action-go-docker .
	docker build -t test-action-go-docker:0.0.2 --target production --platform linux/amd64 .

push: build
	docker tag test-action-go-docker:0.0.2 superflux/test-action-go-docker:0.0.2
	docker push superflux/test-action-go-docker:0.0.2