run:
	$(MAKE) services
	$(MAKE) app 

services:

	cd service && go run ./cmd/server/main.go --port 8080 

frontend:
	cd app && pnpm run start

install:
	cd app && pnpm install
	cd service && go mod download