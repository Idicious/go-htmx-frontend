build:
	go build -o bin/main main.go

run:
	go run main.go

tailwind:
	cd web && npx tailwindcss -i styles/tailwind.css -o public/style.css
