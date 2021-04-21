build:
	protoc --go_out=./proto --go-grpc_out=./proto *.proto
	git add .
	git commit -m "update"
	git push -u origin master
	git tag v${VERSION}
	git push origin master --tags
	curl https://pkg.go.dev/fetch/github.com/mophos/proto-refer@v${VERSION}