all:
	GOOS=linux GOARCH=amd64 go build -o out/dickins-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o out/dickins-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o out/dickins-windows-amd64