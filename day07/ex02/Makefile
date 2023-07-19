SRC = http://localhost:6060/pkg/day07/

doc:
	@~/go/bin/godoc -http=:6060

zip-docs:
	@curl -s -o index.html $(SRC)
	@zip docs.zip index.html

extract-doc:
	@unzip docs.zip

.PHONY: doc zip-docs extract-doc