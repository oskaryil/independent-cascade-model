test: 
	go test ./...

build:
	make clean
	go build

clean:
	rm -f icm

run:
	make build
	./icm -f ./android.csv