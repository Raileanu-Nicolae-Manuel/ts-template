.PHONY: proto-go
proto-go:
	protoc --proto_path=proto \
		--go_out=backend \
		--go-grpc_out=backend \
		proto/*.proto

.PHONY: proto-ts
proto-ts:
	npm run proto:generate

.PHONY: proto
proto: proto-go proto-ts