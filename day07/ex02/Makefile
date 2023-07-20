SRC = http://localhost:6060/pkg/day07/ex02/

install:
	go get golang.org/x/tools/cmd/godoc

doc:
	@~/go/bin/godoc -http=:6060

zip-doc:
	@curl -s -o index.html $(SRC)
	@zip docs.zip index.html
	@rm index.html

extract-doc:
	@unzip docs.zip

clean:
	@rm ./index.html

.PHONY: doc zip-doc extract-doc clean
.DEFAULT_GOAL: doc