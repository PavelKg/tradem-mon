help:
	@echo "Usage:"
	@echo " "
	@echo "	make <command>" [arguments]
	@echo " "
	@echo "The commands are:"
	@echo " "
	@echo "	start_server		Run backend API server"
	@echo "	restart_server		Restart backend API server"
	@echo "	unit_test_server	Start backend API server unit test"
	@echo "	github_test_server	Start backend API server github test"
	@echo "	build			Build backend API server"

start_server:
	docker-compose up -d trm-api

restart_server:
	docker-compose stop trm-api
	docker-compose rm -f
	docker-compose up -d trm-api

create_db:
	docker-compose up -d db

remove_db:
	docker-compose stop db
	docker-compose rm -f
	docker volume rm tradem-mon-api_trm-api-postgres

unit_test_server:
	docker-compose run trm-api go test ./...

github_test_server:
	cd ./src/cmd && go test ./...

build:
	 docker-compose run --no-deps trm-api go build -o build/trm-api cmd/main.go

swagger:
	swagger init



.PHONY: start_server restart_server github_test_server unit_test_server build
