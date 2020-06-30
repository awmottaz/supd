VERSION = $$(cat VERSION)
HASH = $$(git rev-parse HEAD)
LDFLAGS = -X 'github.com/awmottaz/supd/cmd.version=$(VERSION)' -X 'github.com/awmottaz/supd/cmd.hash=$(HASH)'

# builds a fresh binary in $cwd
supd: clean
	go build -ldflags="$(LDFLAGS)"

# installs a binary on your system
install: clean
	go install -ldflags="$(LDFLAGS)"

.PHONY: clean
clean:
	rm -f supd
