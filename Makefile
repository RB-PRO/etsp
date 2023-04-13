
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

build-config:
	go env GOOS GOARCH

build-linux:
	set GOARCH=amd64
	set GOOS=linux
	go build cmd/Georgy/main.go

build-windows:
	set GOARCH=amd64
	set GOOS=windows
	go build cmd/Georgy/main.go

go-car:
	set GOARCH=amd64
	set GOOS=linux
	go env GOOS GOARCH
	go build -o main ./cmd/Georgy/main.go
	scp main Login Password Token root@194.87.107.129:go/etsp/