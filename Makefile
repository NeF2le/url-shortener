generate_auth_service:
	protoc -I ./auth_service \
			--go_out=. \
			--go-grpc_out=. \
			./auth_service/api/auth_service.proto

generate_shortener_service:
	protoc -I ./shortener_service \
			--go_out=. \
			--go-grpc_out=. \
			./shortener_service/api/shortener_service.proto

