build:
	go build -ldflags="-s -w" goctl.go
	mv goctl $(GOPATH)/bin/gqgo
	gqgo template clean
	gqgo template init
	$(if $(shell command -v upx), upx goctl)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o gqgo goctl.go
	$(if $(shell command -v upx), upx goctl-darwin)
	mv gqgo $(GOPATH)/bin/gqgo

win:
	GOOS=windows go build -ldflags="-s -w" -o gqgo.exe goctl.go
	$(if $(shell command -v upx), upx goctl.exe)
	mv gqgo.exe $(GOPATH)/bin/gqgo.exe

linux:
	GOOS=linux go build -ldflags="-s -w" -o gqgo goctl.go
	$(if $(shell command -v upx), upx goctl-linux)
	mv gqgo $(GOPATH)/bin/gqgo

image:
	docker build --rm --platform linux/amd64 -t kevinwan/goctl:$(version) .
	docker tag kevinwan/goctl:$(version) kevinwan/goctl:latest
	docker push kevinwan/goctl:$(version)
	docker push kevinwan/goctl:latest
	docker build --rm --platform linux/arm64 -t kevinwan/goctl:$(version)-arm64 .
	docker tag kevinwan/goctl:$(version)-arm64 kevinwan/goctl:latest-arm64
	docker push kevinwan/goctl:$(version)-arm64
	docker push kevinwan/goctl:latest-arm64
