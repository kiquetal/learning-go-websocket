### Create volume
	
	docker volume create postgres-volume

### Run docker

     docker run --name pgdev -e POSTGRES_PASSWORD=secretpassword -d -p 5432:5432 -v ${PWD}/resources/volume-pg:/var/lib/postgresql/data  postgres

### Run buffalo for migrations

 wget https://github.com/gobuffalo/cli/releases/download/v0.18.14/buffalo_0.18.14_Linux_x86_64.tar.gz
 tar -xvzf buffalo_0.18.14_Linux_x86_64.tar.gz
 sudo mv buffalo /usr/local/bin/buffalo

### Install soda-cli

  go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest


### Just run the project run.sh

 user: admin@example.com
 password: password
