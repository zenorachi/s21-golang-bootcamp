MODULE_0 = ex00
EXECUTABLE_0 = readDB

MODULE_2 = ex02
EXECUTABLE_2 = compareFS

ex00:
	@cd $(MODULE_0) && go build -o $(EXECUTABLE_0)
	@mv $(MODULE_0)/$(EXECUTABLE_0) ./

ex02:
	@cd $(MODULE_2) && go build -o $(EXECUTABLE_2)
	@mv $(MODULE_2)/$(EXECUTABLE_2) ./

clean:
	@rm -rf $(EXECUTABLE_0)
	@rm -rf $(EXECUTABLE_2)

.PHONY: ex00 ex02