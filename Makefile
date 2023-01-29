test:	
	docker-compose -f ./docker/test/docker-compose.yml up --build 

deploy:
	docker-compose -f ./docker/deploy/docker-compose.yml up --build 
