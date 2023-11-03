run:
	go run cmd/cv-api.go

build:
	go build -o bin/api cmd/cv-api.go

compile:
	echo "Compiling for Linux (arm64, amd64) and Windows (amd64, arm)"
	GOOS=linux GOARCH=arm64 go build -o bin/api-linux-arm64 cmd/cv-api.go
	GOOS=linux GOARCH=amd64 go build -o bin/api-linux-amd64 cmd/cv-api.go
	GOOS=windows GOARCH=arm go build -o bin/api-windows-arm cmd/cv-api.go
	GOOS=windows GOARCH=amd64 go build -o bin/api-windows-amd64 cmd/cv-api.go