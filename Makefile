default: install

build:
	go build

install: build
	go install

docker-build:
	docker build -t auribuo/apisense:latest .

docker-run: docker-build
	docker run -it --rm -p 23232:23232 -p 8080:8080 --name apisense auribuo/apisense:latest

clean: clean-build clean-cfg

clean-build:
	go clean -i

clean-cfg:
	rm -r "$HOME"/apisense
	rm -r "$HOME"/.config/apisense