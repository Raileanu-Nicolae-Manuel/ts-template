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
proto: clean proto-go proto-ts

.PHONY: clean
clean:
	rm -rf backend/proto/*
	rm -rf frontend/src/service/api/generate/*