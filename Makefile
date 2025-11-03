_PHONY: watch

watch:
	SERVER_ADDR=127.0.0.1:8000 go tool templ generate --watch --proxy="http://127.0.0.1:8000" --cmd="go run ."
