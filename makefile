build:
	env GO111MODULE=on env GOOS=linux env GOARCH=arm env CGO_ENABLED=0 go build  -o servo/cmd/main servo/main.go
login:
	docker login -u ${DOCKER_USER} -p ${DOCKER_PASS}
docker/build:
	docker build -t halosaka/driver-servo -f Dockerfile ./servo/
docker/push:
	docker push halosaka/driver-servo
local/deploy:
	ls -la
	make build
	make docker/build
	make docker/push
ci/deploy:
	make login
	make local/deploy
