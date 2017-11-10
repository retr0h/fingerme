build:
	@echo "+ $@"
	sudo docker build -t containume:latest .

run:
	@echo "+ $@"
	sudo docker run --rm -it \
		-u fingerd \
		-p 79:1079 \
		containume:latest
