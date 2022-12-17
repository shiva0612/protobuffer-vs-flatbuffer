
domain:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. --go_opt=module=shiva --go-grpc_opt=module=shiva  proto/*.proto
test:
	go test ./bench/... -run=Benchmark -bench=. -benchtime=2s -count=5 -benchmem
