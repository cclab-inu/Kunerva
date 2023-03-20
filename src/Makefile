CURDIR=$(shell pwd)

.PHONY: build
build:
	# for build_image -> CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o Kunerva main.go
	go build -o Kunerva main.go

.PHONY: clean
clean:
	rm -f Kunerva go.sum

.PHONY: image
image:
	cd $(CURDIR)/../build; $(MAKE)
