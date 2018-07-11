VERSION ?= latest
all: dep build docker clean

dep:
	dep ensure -vendor-only

build:
	env GOOS=linux GOARCH=amd64 go build .

docker:
	docker build -t $(REGISTRY)/statist:$(VERSION) .

clean:
	rm statist
