init:
	cd go/command && go mod tidy
tidy:
	cd go/command && go mod tidy
serve:
	go run command/main.go