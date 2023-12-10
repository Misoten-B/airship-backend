local:
	go generate
	go run main.go

lint:
	echo "This is experimental feature"
	docker compose -f compose.yml -f compose.lint.yml up