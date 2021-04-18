include .env
export

project-init:
	go mod init github.com/brandoncate-personal/monitor-cli

build-image:
	docker build -t brandoncate/monitor:v0.0.1 .

release-init:
	docker run --rm --privileged \
	-v $PWD:/go/src/github.com/user/repo \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-w /go/src/github.com/user/repo \
	goreleaser/goreleaser init

release-gen:
	docker run --rm --privileged \
	-v $PWD:/go/src/github.com/user/repo \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-w /go/src/github.com/user/repo \
	-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
	goreleaser/goreleaser --snapshot --skip-publish --rm-dist
