setup-local-env-files:
	cp -R ./app/.env.example ./app/.env
	
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

air-docker:
	docker compose exec app air -c .air.toml

# airの導入方法: https://github.com/air-verse/air?tab=readme-ov-file#installation
air-local:
	cd app && air -c .air.toml