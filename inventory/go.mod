module github.com/nexarise/microservices-factory/inventory

go 1.24

require (
	github.com/nexarise/microservices-factory/shared v0.0.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

replace github.com/nexarise/microservices-factory/shared => ../shared
