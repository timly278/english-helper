# docker build . -t name:tag e.g: go1.21.2:latest

run: 
	go run ./app/main.go

build: 
	go build -o english ./app/main.go 

.PHONY: run build

# Remove the file from the Git index (replace filename with the actual filename)
# git rm --cached filename