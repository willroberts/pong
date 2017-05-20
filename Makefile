test:
	go vet *.go
	golint *.go
	errcheck *.go
	go test
