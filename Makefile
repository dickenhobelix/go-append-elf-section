.PHONY: build build_base clean run

build_base:
	go build -o _main

build_base_tinygo:
	tinygo build -o _main

append:
	objcopy --add-section payload=payload_file _main main

build: build_base append

build_tinygo: build_base_tinygo append

clean:
	rm -f _main main

run: build
	./main

run_tinygo: build_tinygo
	./main
