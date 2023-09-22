proto:
	protoc pkg/auth/pb/auth.proto --go_out=plugins=grpc:.