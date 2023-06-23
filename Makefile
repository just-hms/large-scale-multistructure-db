github:
	docker compose -f ./docker/test/docker-compose.yml up --build --exit-code-from backend --abort-on-container-exit
test:	
	docker compose -f ./docker/test/docker-compose.yml up --force-recreate --build backend
dev:
	docker compose -f ./docker/dev/docker-compose.yml up --force-recreate --build
infrastructure:
	docker compose -f ./docker/infrastructure/docker-compose.yml up --build -d
deploy:
	docker context use replica1
	docker compose -f ./docker/deploy/docker-compose-replica.yml up --force-recreate --build -d
	docker context use replica2
	docker compose -f ./docker/deploy/docker-compose-replica.yml up --force-recreate --build -d 
	docker context use primary
	docker compose -f ./docker/deploy/docker-compose.yml up --force-recreate --build -d
	docker context use default

down: 
	docker compose -f ./docker/dev/docker-compose.yml down 
	docker compose -f ./docker/test/docker-compose.yml down 
	docker compose -f ./docker/infrastructure/docker-compose.yml down 

deploy-down:
	docker context use replica1
	docker compose -f ./docker/deploy/docker-compose-replica.yml down
	docker context use replica2
	docker compose -f ./docker/deploy/docker-compose-replica.yml down
	docker context use primary
	docker compose -f ./docker/deploy/docker-compose.yml down
	docker context use default

doc:
	swag init -g be/internal/controller/router.go --output be/apidocs/

docker-contexts:
	docker context create replica1 --docker host=ssh://root@172.16.5.43
	docker context create replica2 --docker host=ssh://root@172.16.5.47
	docker context create primary --docker host=ssh://root@172.16.5.42
