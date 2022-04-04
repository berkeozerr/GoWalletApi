gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb 
	
gen2:
	protoc --proto_path=proto proto/*.proto  --go_out=pb/ 
	protoc --proto_path=proto proto/*.proto  --go-grpc_out=pb/
	protoc --proto_path=proto proto/*.proto   --openapiv2_out=swagger/
	protoc -I . --grpc-gateway_out ./pb/  \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    proto/wallet_processor.proto 

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

.PHONY: clean gen server client test cert 