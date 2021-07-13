VERSION=1.3.24
publish:
	rm -rf grpc/*
	protoc --go_out=./grpc --go-grpc_out=./grpc ./src/*.proto
	git add .
	git commit -m "update version ${VERSION}"
	git push
	git tag v${VERSION}
	git push origin master --tags
	curl https://proxy.golang.org/github.com/mophos/proto/@v/v${VERSION}.info
build:
	rm -rf grpc/*
	protoc --go-grpc_out=./grpc --go_out=./grpc ./src/*.proto