prune:
	rm -rf ~/workflow
	rm -rf ./bin
	rm /usr/local/bin/workflow

build:
	go build -o ./bin/workflow ./src/main.go

install:
	chmod +x ./bin/workflow
	mv ./bin/workflow /usr/local/bin
	rm -rf ./bin
	mkdir ~/workflow