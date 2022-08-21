build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o release/go-marketplace

docker:
	docker build -t sailor1921/go-marketplace .
	
test:
	go test -v .