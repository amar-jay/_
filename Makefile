run:
	start-service
	start-frontend

start-service:
	cd service && go run ./cmd/server/main.go --port 8080 

start-frontend:
	cd app && pnpm run start

install:
	cd app && pnpm install
	cd service && go mod download