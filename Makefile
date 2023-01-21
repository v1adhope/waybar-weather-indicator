.SILENT:

prod: build
	cp .bin/weather ~/.config/waybar/scripts/weather

run: build
	./.bin/weather

build:
	go build -o .bin/weather
