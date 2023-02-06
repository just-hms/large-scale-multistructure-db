github:
	docker-compose -f ./docker/test/docker-compose.yml up --build --exit-code-from backend --abort-on-container-exit
test:	
	docker-compose -f ./docker/test/docker-compose.yml up --build backend
deploy:
	docker-compose -f ./docker/deploy/docker-compose.yml up --build 
