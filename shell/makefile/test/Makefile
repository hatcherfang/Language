LOG_SERVICE_PATH := $(shell cd ..; pwd)
LOG_SERVICE_CURDIR := $(shell  pwd)
.PHONY: print print1 touchfile build moveBuild
all: print print1 build touchfile moveBuild 

build: 
	@make -C $(LOG_SERVICE_CURDIR)/test1
touchfile : 
	@touch a.txt; touch b.txt
print : touchfile 
	@echo $(LOG_SERVICE_PATH) >> a.txt
	@echo $(LOG_SERVICE_CURDIR) >> a.txt
print1 : touchfile 
	@echo "print1" >> b.txt
	@echo $(LOG_SERVICE_PATH) >> b.txt

moveBuild: build
	@rm -rf $(LOG_SERVICE_CURDIR)/build 
	@mv -f $(LOG_SERVICE_CURDIR)/test1/build ./
cptest:
	cp -rf ./build ./test1/
clean : 
	@rm -rf a.txt; rm -rf b.txt
	@rm -rf $(LOG_SERVICE_CURDIR)/build 
	@make -C $(LOG_SERVICE_CURDIR)/test1 clean
