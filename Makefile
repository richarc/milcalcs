BINARY_NAME=Milcalcs.app
APP_NAME=Milcalc
VERSION=0.1.0

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appVersion ${VERSION} -name ${APP_NAME} -release

## run: builds and runs the application
run:
	go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...