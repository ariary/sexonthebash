build.sexonthebash:
	@echo "build in ${PWD}";go build -o sexonthebash cmd/sexonthebash/main.go

run:
	./run.sh
