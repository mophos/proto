VERSION=1.0.11
publish:
	protoc --go_out=./proto --go-grpc_out=./proto *.proto
	git add .
	git commit -m "update"
	git push -u origin master
	git tag v${VERSION}
	git push origin master --tags
	curl https://proxy.golang.org/github.com/mophos/proto-refer/@v/v${VERSION}.info
build:
	protoc --go_out=./proto --go-grpc_out=./proto *.proto