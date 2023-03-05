.SILENT:

prod: build
	rm /home/rat/.config/waybar/scripts/weather
	cp .bin/weather ~/.config/waybar/scripts/weather

run: build
	./.bin/weather

build:
	go build -o .bin/weather
