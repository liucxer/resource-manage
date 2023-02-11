all: resource-manage_mac resource-manage_linux_x86_64 resource-manage_linux_aarch64 resource-manage_windows

resource-manage_mac:
	mkdir -p build
	go build -o ./build/resource-manage_mac

resource-manage_linux_x86_64:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/resource-manage_linux_x86_64

resource-manage_linux_aarch64:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/resource-manage_linux_aarch64

resource-manage_windows:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/resource-manage_windows

clean:
	\rm -rf ./build/
