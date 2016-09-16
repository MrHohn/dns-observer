OUT_DIR = build
IMAGE = gcr.io/zihongz-kubernetes-codelab/dns-observer
TAG = 0.1

# Rules for building the real image for deployment to gcr.io

compile: dns-observer/
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o $(OUT_DIR)/dns-observer dns-observer/dns-observer.go

go: compile

clean:
	rm -rf build

build: go
	docker build -t ${IMAGE}:$(TAG) .

docker: build

push: docker
	gcloud docker push ${IMAGE}:$(TAG)
