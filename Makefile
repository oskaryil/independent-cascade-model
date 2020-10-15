test: 
	go test ./...

build-mac:
	make clean
	go build -o build/icm_mac

build-linux:
	make clean
	env GOOS=linux GOARCH=amd64 go build -o build/icm_linux

clean:
	rm -rf build/*

run:
	make build
	./icm -f ./android.csv