build:
	go build -o reqx ./cmd/app/main.go

run: 
	./reqx

start:
	go run ./cmd/app/main.go