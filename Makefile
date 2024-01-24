ifeq ($(OS),Windows_NT)
EXT =.exe
else
EXT =
endif

all : charcode-converter$(EXT)

charcode-converter$(EXT) : ./cmd/charcode-converter/main.go
	go build -o charcode-converter$(EXT) ./cmd/charcode-converter

.PHONY: clean

clean:
	rm charcode-converter$(EXT)
