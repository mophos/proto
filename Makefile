VERSION=1.2.1
publish:
	protoc --go_out=./proto --go-grpc_out=./proto ./src/*.proto
	git add .
	git commit -m "update version ${VERSION}"
	git push
	git tag v${VERSION}
	git push origin master --tags
	curl https://proxy.golang.org/github.com/mophos/proto/@v/v${VERSION}.info
build:
	protoc --go_out=./proto --go-grpc_out=./proto ./src/*.proto