# Makefile

APP_NAME = cleanout

build:
	go build -o $(APP_NAME) .

run:
	go run main.go

clean:
	rm -f $(APP_NAME)
