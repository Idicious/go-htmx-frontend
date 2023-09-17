BIN=bin

build-server:
	go mod download
	go build -ldflags "-s -w" -o $(BIN)/main main.go

build-web:
	cd web && npx tailwindcss -i styles/tailwind.css -o public/style.css

build: build-server build-web

tailwind:
	cd web && npx tailwindcss -i styles/tailwind.css -o public/style.css --watch
