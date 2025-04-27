run: build
	@./bin/task-cli $(ARGS)
#make run ARGS="list done"

build:
	@go build -o bin/task-cli cmd/main.go

test:
	go test cmd