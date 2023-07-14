.PHONY: build test clean

PORT = 3000

build:
	go build -o candy_server ./cmd/candy/

clean:
	@rm -rf log1.txt log2.txt log3.txt log4.txt true_log1.txt true_log2.txt true_log3.txt true_log4.txt candy_server

test: test1 test2 test3 test4

test1:
	@touch true_log1.txt
	@echo '{"change":0,"thanks":"Thank you!"}' > true_log1.txt
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 46, "candyType": "YR", "candyCount": 2}' http://127.0.0.1:$(PORT)/buy_candy > log1.txt
	diff true_log1.txt log1.txt

test2:
	@touch true_log2.txt
	@echo '{"change":5,"thanks":"Thank you!"}' > true_log2.txt
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": 1}' http://127.0.0.1:$(PORT)/buy_candy > log2.txt
	diff true_log2.txt log2.txt

test3:
	@touch true_log3.txt
	@echo '{"error":"oops... no such candies or invalid candies count :("}' > true_log3.txt
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AAdf", "candyCount": 1}' http://127.0.0.1:$(PORT)/buy_candy > log3.txt
	diff true_log3.txt log3.txt

test4:
	@touch true_log4.txt
	@echo '{"error":"oops... no such candies or invalid candies count :("}' > true_log4.txt
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": -1}' http://127.0.0.1:$(PORT)/buy_candy > log4.txt
	diff true_log4.txt log4.txt

.DEFAULT_GOAL: build