.PHONY: build clean test test1 test2 test3 test4 test5

PORT = 8080

build:
	go build ./candy-server/cmd/candy

clean:
	@rm -rf log1.txt log2.txt log3.txt log4.txt log5.txt true_log1.txt true_log2.txt true_log3.txt true_log4.txt true_log5.txt candy

test: test1 test2 test3 test4 test5

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
	@echo '{"error":"invalid input data"}' > true_log3.txt
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AAdf", "candyCount": 1}' http://127.0.0.1:$(PORT)/buy_candy > log3.txt
	diff true_log3.txt log3.txt

test4:
	@touch true_log4.txt
	@echo '{"error":"invalid input data"}' > true_log4.txt
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": -1}' http://127.0.0.1:$(PORT)/buy_candy > log4.txt
	diff true_log4.txt log4.txt

test5:
	@touch true_log5.txt
	@echo '{"error":"You need 5 more money!"}' > true_log5.txt
	curl -XPOST -H "Content-Type: application/json" -d '{"money": 10, "candyType": "AA", "candyCount": 1}' http://127.0.0.1:$(PORT)/buy_candy > log5.txt
	diff true_log5.txt log5.txt

.DEFAULT_GOAL: build