install:
	cp ./bin/task /usr/bin/

build:
	go build  -o ./bin/task

small-build:
	go build -ldflags="-s -w" -o ./bin/task

clean:
	rm -rf ./bin/task

