project-init:
	go mod init github.com/brandoncate-personal/monitor-cli

build-image:
	docker build -t brandoncate/monitor:v0.0.1 .