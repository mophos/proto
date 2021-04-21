build:
	protoc --go_out=./proto --go-grpc_out=./proto *.proto
publish:
	git add .
	git commit -m "update"
	git push -u origin master
	git tag v${VERSION}
	git push origin master

	latest="$(git describe --tags $(git rev-list --tags --max-count=1))"
	curl https://proxy.golang.org/github.com/mophos/proto-refer/@v/$latest.info