working_dir := $(PWD)

lint:
	cd ${working_dir}/src && golangci-lint run --build-tags integration

unit:
	cd ${working_dir}/src && go test -cover ./... && go fmt ./...

integration:
	cd ${working_dir}/src && go test -cover -tags=integration ./...

build-app:
	docker build -t greeting .

up:
	docker-compose up -d

down:
	docker-compose down