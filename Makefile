.PHONY: build run compile

# Set the Golang binary name
BIN_NAME = app

# Output directory
OUTPATH = bin

# Set the path to your main.go file
MAIN_FILE = src/core/main.go

# Compile for all OS and arch types
OS_ARCH = "darwin/amd64 linux/amd64 windows/amd64"

build:
	go generate 
	go build -o $(OUTPATH)/$(BIN_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

compile:
	$(foreach GOOS_ARCH, $(OS_ARCH), \
		GOOS=$(firstword $(subst /, ,$(GOOS_ARCH))) \
		GOARCH=$(lastword $(subst /, ,$(GOOS_ARCH))) \
		go build -o $(OUTPATH)/$(BIN_NAME)-$(GOOS_ARCH) $(MAIN_FILE);)

