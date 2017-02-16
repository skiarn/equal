GOOS=windows GOARCH=amd64 go build -o equal.exe main.go;
GOOS=linux GOARCH=amd64 go build -o equal main.go;
GOOS=darwin GOARCH=386 go build -o mac_equal main.go;
