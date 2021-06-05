require-%:
	@if [ "${${*}}" = "" ]; then \
        	echo "Parameter '$*' is required."; \
        	exit 1; \
    	fi

new-go: require-day require-year
	@echo "Creating new file structure for day" $(day)"..."
	mkdir -p $(year)/day-$(day)/bin; \
	touch $(year)/day-$(day)/main.go; \
	touch $(year)/day-$(day)/input;

build: require-day require-year
	go build -o $(year)/day-$(day)/bin/main $(year)/day-$(day)/main.go

run: require-day require-year
	go run $(year)/day-$(day)/main.go

