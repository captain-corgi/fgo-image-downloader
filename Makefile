all: tidy vendor build run

build:
	go build -o fid cmd/fid/main.go

run:
	./fid

clean:
	rm ./fid
	rm ./fid.exe
	rm -f vendor

tidy:
	go mod tidy

vendor:
	go mod vendor

.PHONY: coverage
# coverage:
# 	go test \
# 		-race -covermode=atomic -timeout 30s \
# 		-coverprofile=coverage/coverage.out \
# 		github.com/captain-corgi/fid/pkg/iptrans