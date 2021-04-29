VERSION=1.3.4
publish:
	rm -rf services/*
	protoc --go_out=./services --go-grpc_out=./services ./src/*.proto
	git add .
	git commit -m "update version ${VERSION}"
	git push
	git tag v${VERSION}
	git push origin master --tags
	curl https://proxy.golang.org/github.com/mophos/proto/@v/v${VERSION}.info
build:
	rm -rf services/*
	protoc --go_out=./services --go-grpc_out=./services ./src/*.proto