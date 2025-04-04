GOFILES = (main.go tokenizer.go)

make:
	go run .

build:
	go build GOFILES

clean:
	rm main