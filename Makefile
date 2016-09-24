all: clean build

clean:
	rm -f charon

build: genproto
	go build .

syncbuild: syncproto genproto
	go build .

genproto:
	./proto/gen_go.sh

syncproto:
	cd proto && git pull origin master

init:
	git submodule update --init

install: init genproto
	glide install

docker-build:
	docker build -t charon .

docker-push:
	docker tag charon:latest 096202052535.dkr.ecr.us-west-2.amazonaws.com/charon:latest
	docker push 096202052535.dkr.ecr.us-west-2.amazonaws.com/charon:latest
