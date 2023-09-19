entrypoint := ./cmd/qms-cli/main.go
build_dir=./build
build_name=qms-cli

all: clean build

run:
	go run $(entrypoint)

clean:
	rm -rf $(build_dir)

prebuild: clean
	echo $(entrypoint)/$(build_dir)
	mkdir -p $(build_dir)

build-win_amd64: prebuild
	GOOS=windows GOARCH=amd64 go build -o $(build_dir)/$(build_name)-win64.exe $(entrypoint)

build-linux_amd64: prebuild
	GOOS=linux GOARCH=amd64 go build -o $(build_dir)/$(build_name)-linux64 $(entrypoint)
