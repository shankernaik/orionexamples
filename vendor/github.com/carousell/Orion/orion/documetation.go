/*
Package orion is a small lightweight framework written around grpc with the aim to shorten time to build microservices

Source code for Orion can be found at https://github.com/carousell/Orion

It is derived from 'Framework' a small microservices framework written and used inside https://carousell.com, It comes with a number of sensible defaults such as zipkin tracing, hystrix, live reload of configuration, etc.


Why Orion

Orion uses protocol-buffers definitions (https://developers.google.com/protocol-buffers/docs/proto3) using gRPC and orion proto plugin as base for building services.

Using proto definitions as our service base allows us to define clean contracts that everyone can understand and enables auto generation of client code.

You define your services as a proto definition, for example
	service SimpleService{
		rpc Echo (EchoRequest) returns (EchoResponse){
		}
	}
	message EchoRequest {
		string msg = 1;
	}
	message EchoResponse {
		string msg = 1;
	}
The above definition represents a service named 'SimpleService' which accepts 'EchoRequest' and returns 'EchoResponse' at 'Echo' endpoint.

After you have generated the code from protoc (using grpc and orion plugin), you need to implement the server interface generated by gRPC, orion uses this service definition and enables HTTP/gRPC calls to be made to the same service implementation.

How do i use it

Lets go through the example defined at https://github.com/carousell/Orion/tree/master/example/simple , It covers the minimum required implementation for orion. It has the following structure
	.
	├── cmd
	│   ├── client
	│   │   └── client.go
	│   └── server
	│       └── server.go
	├── service
	│   └── service.go
	└── simple_proto
		├── generate.sh
		└── simple.proto

First we define the proto 'simple_proto/simple.proto' as

	syntax = "proto3";

	package simple_proto;

	service SimpleService{
		rpc Echo (EchoRequest) returns (EchoResponse){
		}
	}

	message EchoRequest {
		string msg = 1;
	}

	message EchoResponse {
		string msg = 1;
	}
The above definition represents a service named 'SimpleService' which accepts 'EchoRequest' and returns 'EchoResponse' at 'Echo' endpoint.

Now we can execute 'generate.sh' which contains
	protoc -I . simple.proto --go_out=plugins=grpc:. --orion_out=.
Running this command generates the following file in the simple_proto directory:
	.
	├── generate.sh
	├── simple.pb.go
	├── simple.proto
	└── simple.proto.orion.pb.go
This contains:
	All the protocol buffer code to populate, serialize, and retrieve our request and response message types (simple.pb.go)
	An interface type (or stub) for clients to call with the methods defined in the SimpleService (simple.pb.go)
	An interface type for servers to implement, also with the methods defined in the SimpleService (simple.pb.go)
	Registration function for Orion (simple.proto.orion.pb.go)

Whats Incuded

Orion comes included with.
	Hystrix (http://github.com/afex/hystrix-go)
	Zipkin (http://github.com/opentracing/opentracing-go)
	NewRelic (http://github.com/newrelic/go-agent)
	Prometheus (http://github.com/grpc-ecosystem/go-grpc-prometheus)
	Pprof (https://golang.org/pkg/net/http/pprof/))
	Configuration (http://github.com/spf13/viper)
	Live Configuration Reload (http://github.com/carousell/Orion/utils/listenerutils)
	And much more...

Getting Started

First follow the install guide at https://github.com/carousell/Orion/blob/master/README.md
*/
package orion

// This comment block (re)generates the documentation.
//go:generate godoc2ghmd -ex -file=README.md github.com/carousell/Orion/orion
//go:generate godoc2ghmd -ex -file=handlers/README.md github.com/carousell/Orion/orion/handlers
//go:generate godoc2ghmd -ex -file=modifiers/README.md github.com/carousell/Orion/orion/modifiers
//go:generate godoc2ghmd -ex -file=helpers/README.md github.com/carousell/Orion/orion/helpers
