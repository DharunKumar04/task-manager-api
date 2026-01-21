.PHONY: docker-up docker-down docker-clean

docker-up:
	docker network create task-network || true
	docker run --name postgres-17 \
		--network task-network \
		-e POSTGRES_USER=taskuser \
		-e POSTGRES_PASSWORD=taskpass123 \
		-e POSTGRES_DB=taskdb \
		-p 5432:5432 \
		-v postgres_data:/var/lib/postgresql/data \
		-d postgres:17
	docker run --name pgadmin \
		--network task-network \
		-e PGADMIN_DEFAULT_EMAIL=admin@example.com \
		-e PGADMIN_DEFAULT_PASSWORD=admin \
		-p 5050:80 \
		-v pgadmin_data:/var/lib/pgadmin \
		-d dpage/pgadmin4

docker-down:
	docker stop postgres-17 pgadmin || true
	docker rm postgres-17 pgadmin || true

docker-clean:
	docker volume rm postgres_data pgadmin_data || true