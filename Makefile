github:
	docker-compose -f ./docker/test/docker-compose.yml up --build --exit-code-from backend --abort-on-container-exit
test:	
	docker-compose -f ./docker/test/docker-compose.yml up --build backend
dev:
	docker-compose -f ./docker/dev/docker-compose.yml up --build
deploy:
	docker-compose -f ./docker/deploy/docker-compose.yml up --build 


down: 
	docker-compose -f ./docker/deploy/docker-compose.yml down 
	docker-compose -f ./docker/dev/docker-compose.yml down 
	docker-compose -f ./docker/test/docker-compose.yml down 

