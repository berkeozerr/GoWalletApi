# GoWalletApi
Go Wallet Api provides bitcoin wallet functions such as random mnemonic , HD SegWit address and multi-sig P2SH bitcoin address generation.
Project consist of 3 unary RPC API and provides a reverse-proxy server which translates a RESTful HTTP API into gRPC using gRPC-Gateway.
## 1. Generating random mnemonic words following BIP39 standard
```
[GET] .../v1/wallet/random_mnemonic
```
### Response:
```json
{
  "seedPhrase": "string",
  "path": "string"
}
```
## 2. Generate a Hierarchical Deterministic (HD) Segregated Witness (SegWit) bitcoin address from a given seed and path
```
[POST] .../v1/wallet/generate_hd_seg_with
```
### Request Body:
```json
{
  "seedPhrase": "string",
  "path": "string"
}
```
### Response:
```json
{
  "generatedAddress": "string"
}
```
## 3. Generate an n-out-of-m Multisignature (multi-sig) Pay-To-Script-Hash (P2SH) bitcoin address, where n, m and public keys can be specified
### Request Body:
```json
{
  "n": 0,
  "m": 0,
  "publicKeys": [
    "string"
  ]
}
```
### Response:
```json
{
  "generatedAddress": "string"
}

```
## Setup Dev Environment
- ### Install ***protoc***:
``` 
brew install protobuf 
```
- ### Install ***protoc-gen-go***, ***protoc-gen-go-grpc***, ***protoc-gen-grpc-gateway*** and ***protoc-gen-openapiv2***:
``` 
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```
## How to Run 
| Command     | Description | 
| :---        | :----:   |  
| make server | Run Server  | 
| make rest  | Run Rest Server  |
| make clean   | Cleans pb and swagger folder |
| make gen  | Code generation from proto pb   |
| make test  | Run Tests|
| make cert  | Generate SSL/TLS certificates|
