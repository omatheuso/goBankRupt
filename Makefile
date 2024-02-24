build:
	@go build -o bin/goBankRupt
run: build
	@./bin/goBankRupt
test:
	@go test -v ./...
