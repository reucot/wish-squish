run: build start

start: 
	sudo ./main

build: 
	go build cmd/server/main.go