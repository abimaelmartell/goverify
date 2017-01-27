.PHONY: run build_docker run_docker


run:
	go build . && ./goverify

build_docker:
	docker build -t goverify .

run_docker:
	docker run -p 8080:8080 -t goverify