# Makefile

build-containers:
	docker build --rm -t go-db:latest -f ./build/Dockerfile.go-db ./src


run-container:
	docker run -p 8080:8080 -d go-db:latest