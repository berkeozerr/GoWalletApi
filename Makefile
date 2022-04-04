gen:
	cd proto; buf generate
	
cert:
	cd cert; ./gen.sh; cd ..    

clean:
	rm pb/*.go
	rm swagger/*

server:
	go run cmd/server/main.go -port 8080

test:
	go test -cover -race ./...
rest:
	go run cmd/server/main.go -port 8081 -type rest 

.PHONY: clean gen server test cert 