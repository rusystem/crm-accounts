

  	protoc --go_out=pkg/gen --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pkg/gen --go-grpc_opt=paths=source_relative proto/company.proto

  	protoc --go_out=pkg/gen --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pkg/gen --go-grpc_opt=paths=source_relative proto/user.proto

  	protoc --go_out=pkg/gen --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pkg/gen --go-grpc_opt=paths=source_relative proto/sections.proto