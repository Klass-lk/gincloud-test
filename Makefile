.PHONY: build clean build-testprojectFunction

build: clean
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
	mkdir -p bin/
	zip bin/testproject.zip bootstrap
	rm bootstrap

build-testprojectFunction:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
	mkdir -p $(ARTIFACTS_DIR)
	cp bootstrap $(ARTIFACTS_DIR)/

clean:
	rm -rf bin/
	rm -f bootstrap
	rm -f testproject.zip