module github.com/nexarise/microservices-factory/order

go 1.24

require (
	github.com/nexarise/microservices-factory/shared v0.0.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require github.com/google/uuid v1.6.0 // indirect

replace github.com/nexarise/microservices-factory/shared => ../shared
