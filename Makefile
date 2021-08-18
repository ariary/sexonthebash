before.build:
	go mod download && go mod vendor

build.sexonthebash:
	@echo "build in ${PWD}";go build -o sexonthebash cmd/sexonthebash/main.go

build.shellonthebeach:
	@echo "build in ${PWD}";go build -o shellonthebeach cmd/shellonthebeach/main.go
