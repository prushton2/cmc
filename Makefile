GOFILES = main.go tokenizer.go
PROGRAM = main

make:
	go run .

main:
	go build -o $(PROGRAM) $(GOFILES)

clean:
	rm main

