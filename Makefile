github:
	docker compose -f ./docker/test/docker-compose.yml up --build --exit-code-from backend --abort-on-container-exit
test:	
	docker compose -f ./docker/test/docker-compose.yml up --force-recreate --build backend
dev:
	docker compose -f ./docker/dev/docker-compose.yml up --force-recreate --build
infrastructure:
	docker compose -f ./docker/infrastructure/docker-compose.yml up --build -d
deploy:
	docker compose -f ./docker/deploy/docker-compose.yml up --force-recreate --build 

down: 
	docker compose -f ./docker/deploy/docker-compose.yml down 
	docker compose -f ./docker/dev/docker-compose.yml down 
	docker compose -f ./docker/test/docker-compose.yml down 
	docker compose -f ./docker/infrastructure/docker-compose.yml down 

doc:
	swag init -g be/internal/controller/router.go --output be/apidocs/
