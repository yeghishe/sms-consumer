default: run

clean:
	rm -rf bin

deps:
	go get github.com/tools/godep
	godep restore	

build: clean deps
	go build -o bin/sms-consumer

run: build
	bin/sms-consumer

docker: build
	docker build -t yeghishe/sms-consumer .
