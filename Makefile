app-bash:
	docker compose exec app bash

up:
	docker-compose up -d

up-with-build:
	docker-compose up -d --build

down:
	docker-compose down

log:
	tail -f app/src/app.log

air:
	docker compose exec app air -c .air.toml