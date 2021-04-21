build:
	protoc --go_out=./proto --go-grpc_out=./proto *.proto
	git add .
	git commit -m "update"
	git push -u origin master
	git tag v${VERSION}
	git push origin master --tags
	latest="$(git describe --tags $(git rev-list --tags --max-count=1))"
	curl https://proxy.golang.org/github.com/moph-gateway/his-proto/@v/$latest.info