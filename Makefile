all: run

run:
	go run main.go logon.go

push:
	git push git@github.com:RB-PRO/avtoto.git

pull:
	git pull git@github.com:RB-PRO/avtoto.git

pushW:
	git push https://github.com/RB-PRO/avtoto.git

pullW:
	git pull https://github.com/RB-PRO/avtoto.git
