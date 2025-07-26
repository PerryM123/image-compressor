app-bash:
	docker compose exec app bash

up:
	docker-compose up -d

up-with-build:
	docker-compose up -d --build

down:
	docker-compose down
