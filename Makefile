.SILENT:

default: build

build:
	go build -o .bin/weather cmd/main.go

test: build
	./.bin/weather | jq .

prod: build
	mkdir -p ~/.config/waybar/scripts/
	\cp -f .bin/weather ~/.config/waybar/scripts/
