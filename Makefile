gen-proto:
	protoc -I pkg/plugins/proto/ pkg/plugins/proto/plugin.proto --go_out=plugins=grpc:pkg/plugins/proto/