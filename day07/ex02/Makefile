SRC = http://localhost:6060/pkg/day07/ex02/

install:
	go get golang.org/x/tools/cmd/godoc

doc:
	@~/go/bin/godoc -http=:6060

zip-docs:
	@curl -s -o . $(SRC)
	@zip docs.zip index.html
	@rm index.html

extract-doc:
	@unzip docs.zip

clean:
	@rm ./index.html

.PHONY: doc zip-docs extract-doc clean
.DEFAULT_GOAL: doc