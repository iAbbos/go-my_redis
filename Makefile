include .env

build:
	if [ -f "./bin/${BINARY_NAME}" ]; then \
  		rm "./bin/${BINARY_NAME}"; \
  		echo "deleted ${BINARY_NAME}"; \
	fi
	@echo "Building binary..."
	go build -o "./bin/${BINARY_NAME}" ${CMD_DIR}

run: build
	"./bin/${BINARY_NAME}"
	@echo "App is started..."

stop:
	@echo "Stopping app..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "App stopped..."
