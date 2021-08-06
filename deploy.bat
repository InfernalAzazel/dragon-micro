go env -w GOOS=linux
go build -o ./build/dragon-micro main.go
go env -w GOOS=windows