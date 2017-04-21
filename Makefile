deps:
	which godep || go get github.com/tools/godep
	godep restore

save:
	godep save

build:
	go build ./main.go

run:
	./main