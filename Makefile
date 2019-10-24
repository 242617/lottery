
.PHONY: build
build:
	go build \
		-o lottery \
		cmd/lottery/main.go

.PHONY: utils\:private2address
utils\:private2address:
	go build \
		-o private2address \
		cmd/private2address/main.go

.PHONY: utils\:keystore2private
utils\:keystore2private:
	go build \
		-o keystore2private \
		cmd/keystore2private/main.go
