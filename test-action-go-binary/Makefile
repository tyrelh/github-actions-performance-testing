build:
	export GOOS=linux && \
	export GOARCH=amd64 && \
	export CGO_ENABLED=0 && \
	go build -o test-action-go-binary .

push: build
	aws s3 cp test-action-go-binary s3://giftbit-deployment-artifacts/test-action-go-binary --profile infrastructure-admin-dev