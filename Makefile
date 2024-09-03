all: getpack proto run
.PHONY: getpack
getpack:
	@echo "Getting package..."
	@go get -u
	@echo "Done."
.PHONY: run
run:
	@echo "Running server"
	@air
# Clean proto file
.PHONY: clean
clean:
	@echo "Cleaning proto files..."
	@rm -rf /pb/*.go

.PHONY: proto
proto: clean
	@echo "Generating proto files..."
	@protoc proto/*.proto --go_out=./ --go-grpc_out=./

