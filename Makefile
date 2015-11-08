deps:
	go get "github.com/jeffail/gabs"

wwms:
	go fmt api.go
	go fmt bin/wwms.go
	go build -o bin/wwms bin/wwms.go

echo:
	go fmt api.go
	go fmt bin/echo.go
	go build -o bin/echo bin/echo.go

examples: wwms echo
