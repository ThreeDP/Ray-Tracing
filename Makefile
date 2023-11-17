COMPOSE = docker-compose.yml
IMAGES = $(docker ps -qa)
DIR_VOLUMES = /nfs/homes/dapaulin/project-volumes/postgresql

all: build up

build:
	mkdir -p $(DIR_VOLUMES)
	docker-compose -f $(COMPOSE) build

up:
	docker-compose -f $(COMPOSE) up -d

down:
	docker-compose -f $(COMPOSE) down

clean:
	docker-compose -f $(COMPOSE) down --rmi all

fclean: clean
	rm -rf $(DIR_VOLUMES)

.PHONY: all build up down fclean