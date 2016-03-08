export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

test:
	@go test -v mincss.go mincss_test.go

fmt:
	go fmt *.go
