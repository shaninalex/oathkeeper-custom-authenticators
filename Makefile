start:
	docker compose up -d --build

down:
	docker compose down

restart: down start