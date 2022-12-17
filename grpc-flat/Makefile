# flatc --binary api/fbs/req.fbs cmd/client/ip.json
# flatc --binary api/fbs/res.fbs cmd/server/op.json

gen:

	flatc --go --grpc --gen-object-api --go-module-name shiv -o api/ api/fbs/*.fbs
test:
	go test ./cmd/bench/... -run=Benchmark -bench=. -benchtime=2s -count=5 -benchmem

