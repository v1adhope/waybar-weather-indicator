.SILENT:

prod: build
	cp .bin/weather ~/.config/waybar/scripts/weather

test: build
	cp .bin/weather ~/.config/waybar/scripts/weatherTest

run: build
	./.bin/weather

build:
	go build -o .bin/weather
