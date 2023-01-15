generate:
	protoc -I apis/v1alpha1 apis/v1alpha1/*.proto --go_out=grpc --go_opt=paths=source_relative --go-grpc_out=grpc --go-grpc_opt=paths=source_relative