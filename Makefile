build:
	protoc --go_out=./proto --go-grpc_out=./proto *.proto
push:
	git add .
	git commit -m "update"
	git push -u origin master