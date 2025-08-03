generate_auth_service:
	protoc -I ./auth_service \
			--go_out=. \
			--go-grpc_out=. \
			./auth_service/api/auth_service.proto

