proto: 
	buf generate

sqlc:
	sqlc generate --experimental

.PHONY: proto sqlc
