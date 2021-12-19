.PHONY: prp
prp:
	@protoc -I .  call.proto --go_out=call --go-grpc_out=call --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative --js_out=import_style=commonjs:client/src/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:client/src/proto

.PHONY: dev
dev:
	docker-compose -f docker/docker-compose.yml up --build