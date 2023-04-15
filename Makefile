proto: 
	rm -f gen/pb/*.go
	protoc --proto_path=proto --go_out=gen/pb --go_opt=paths=source_relative \
	--go-grpc_out=gen/pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

sqlc:
	sqlc generate --experimental

.PHONY: proto sqlc
