init:
	cd go/command && go mod tidy

serve:
	go run command/main.go