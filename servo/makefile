build:
	env GOOS=linux env GOARCH=arm env CGO_ENABLED=0 go build -o ./cmd/main main.go
	docker build -t makki0205/test .

push:
	make build
	docker push makki0205/test