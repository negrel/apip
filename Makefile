# Import config
cnf ?= .env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

# DOCKER TASKS
# Build the container
build:
	docker build -t $(APP_NAME) .

# Build the container without caching
build-nc:
	docker build --no-cache -t $(APP_NAME) .

# Run container on configured port
run:
	docker run -i -t --rm --env-file=".env" -p=$(PORT):$(PORT) --name=$(APP_NAME) $(APP_NAME)

# Run container on configured port
up: build run

# Stop and remove a running container
stop:
	docker stop $(APP_NAME); docker rm $(APP_NAME)


# OTHER TASKS
# Compile go to binary
bin: clean
	mkdir -p $(BIN_OUTPUT)
	$(GO_BUILD_ENV) go build -v -o $(BIN_OUTPUT) .

# Clean the compiled go binary
clean:
	rm -rf $(BIN_OUTPUT)

# Deploy to heroku
heroku: build
	heroku container:push web -a $(HEROKU_APP)
