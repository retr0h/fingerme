build:
	@echo "+ $@"
	sudo docker build -t retr0h/containume:latest .

push:
	@echo "+ $@"
	docker push retr0h/containume:latest

run:
	@echo "+ $@"
	sudo docker run --rm -it \
		-u fingerd \
		-p 79:1079 \
		retr0h/containume:latest
