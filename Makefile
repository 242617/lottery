SERVICE_NAME ?= 242617/lottery:1.0.0

# Debug
.PHONY: build
build:
	go build \
		-o bin/app \
		cmd/main/main.go

.PHONY: run
run: build
	./bin/app


# Docker
docker\:build:
	docker build \
		-t ${SERVICE_NAME} \
		-f Dockerfile \
		.

docker\:run:
	docker run \
		-it --rm \
		${SERVICE_NAME}

docker\:push:
	docker push ${SERVICE_NAME}


# Node
node:
	docker run -it \
		-p 8070:8070 \
		-p 30303:30303 \
		ethereum/client-go \
			--http \
			--http.addr=localhost \
			--http.port=8070 \
			--syncmode=light \
			--nousb

node\:check:
	curl \
		-X POST \
		-H "content-type: application/json" \
		--data '{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":67}' \
		localhost:8070
