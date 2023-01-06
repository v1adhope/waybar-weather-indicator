.SILENT:

run: build
	./.bin/weather

build:
	go build -o .bin/weather
