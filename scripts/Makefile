APP_NAME = cleanout
ROOT_DIR = ..
MAIN_FILE = $(ROOT_DIR)/main.go

build:
	cd $(ROOT_DIR) && go build -o $(APP_NAME) .

run:
	cd $(ROOT_DIR) && go run $(MAIN_FILE)

clean:
	cd $(ROOT_DIR) && rm -f $(APP_NAME)
