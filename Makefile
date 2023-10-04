.PHONY: build
build:
	go run main.go && \
 	tectonic build/cheatsheet.tex

.PHONY: watch
watch:
	find . | grep -v git | grep -v build | entr -c -r make build
