init:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/google/wire/cmd/wire@latest
	
local:
	go generate
	go run main.go

lint:
	echo "This is experimental feature"
	docker compose -f compose.yml -f compose.lint.yml up