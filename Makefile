VERSION = $$(cat VERSION)
LDFLAGS = -X 'main.Version=$(VERSION)'

# builds a fresh binary in $cwd
supd: clean
	go build -ldflags="$(LDFLAGS)"

# installs a binary on your system
install: clean
	go install -ldflags="$(LDFLAGS)"

.PHONY: clean
clean:
	rm -f supd
