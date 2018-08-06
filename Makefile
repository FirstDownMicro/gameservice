buildlocal:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/Nanixel/FirstDownMicro/gameservice	\
		proto/gameservice.proto
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/Nanixel/FirstDownMicro/gameservice	\
		proto/gameservice.proto
	docker build -t gameservice .
runlocal:
	go run *.go --server_address :50051 --registry mdns --broker nats --broker_address :4222
run:
	docker run --net="host" \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-e MICRO_BROKER=nats \
		-e MICRO_BROKER_ADDRESS=:4222 gameservice