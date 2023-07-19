SRC = http://localhost:6060/pkg/day07/ex02/

doc:
	@~/go/bin/godoc -http=:6060

zip-docs:
	curl -o index.html -J $(SRC)
	#curl -J $(SRC)
	#curl -s -o . $(SRC)
#	@zip docs.zip index.html
#	@rm index.html

extract-doc:
	@unzip docs.zip

clean:
	@rm ./index.html

.PHONY: doc zip-docs extract-doc clean