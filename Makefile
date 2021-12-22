.PHONY: prp
prp:
	@protoc -I . sample.proto --go_out=plugins=grpc:server/pb --go_opt=paths=source_relative --grpc-gateway_out=logtostderr=true:server/pb

.PHONY: dev
dev:
	docker-compose -f docker/docker-compose.yml up --build

.PHONY: proto
proto:
	protoc -I ../../go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ \
	-I../../go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/ \
	--grpc-gateway_out=logtostderr=true:chat --grpc-gateway_opt=paths=source_relative -I . ./chat.proto \
	&& protoc -I/usr/local/include \
	-I ../../go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	-I../../go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/ \
	--go_out=./chat --go_opt=paths=source_relative --go-grpc_out=./chat --go-grpc_opt=paths=source_relative -I . ./chat.proto


.PHONY: compile
compile:
	protoc -Iproto --grpc-gateway_out=logtostderr=true:./server/pb --go_out=plugins=grpc:server/pb --go_opt=paths=source_relative proto/*.proto

.PHONY: envoy
envoy:
	protoc -I/usr/local/include -I.  -I../../go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	--include_imports \
	--include_source_info \
	--descriptor_set_out=./envoy/proto.pb \
	./chat.proto

.PHONY: api
api:
	protoc -I . --openapiv2_out ./openapiv2 \
	-I ../../go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ \
	-I../../go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/ \
    --openapiv2_opt logtostderr=true \
     --openapiv2_opt use_go_templates=true \
    ./chat.proto

