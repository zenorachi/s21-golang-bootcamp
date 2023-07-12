EXECUTABLE = stats

build:
	go build

run: build
	./$(EXECUTABLE)

help:
	@./$(EXECUTABLE) -help

test:
	@cd tests && go test -v

clean:
	@rm -rf $(EXECUTABLE)