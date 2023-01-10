
PROJECTNAME=$(shell basename "$(PWD)")

all: run

run:
	go run cmd/Georgy/main.go

push:
	git push git@github.com:RB-PRO/etsp.git

pull:
	git pull git@github.com:RB-PRO/etsp.git

pushW:
	git push https://github.com/RB-PRO/etsp.git

pullW:
	git pull https://github.com/RB-PRO/etsp.git

build:
	set GOARCH=amd64
	set GOOS=linux
	go build cmd/Georgy/main.go