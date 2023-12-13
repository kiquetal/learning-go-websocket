### Create volume
	
	docker volume create postgres-volume

### Run docker

     docker run --name pgdev -e POSTGRES_PASSWORD=secretpassword -d -p 5432:5432 -v ${PWD}/resources/volume-pg:/var/lib/postgresql/data  postgres

