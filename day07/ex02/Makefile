SRC = http://localhost:6060/pkg/day07/

dvi:
	@~/go/bin/godoc -http=:6060

zip-docs:
	@curl -s -o index.html $(SRC)
	@zip docs.zip index.html

extract-docs:
	@unzip docs.zip

.PHONY: dvi zip-docs extract-docs