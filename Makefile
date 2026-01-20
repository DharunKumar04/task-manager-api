.PHONY: docker-up docker-down docker-logs

docker-up:
	docker run --name postgres-17 \
		-e POSTGRES_USER=taskuser \
		-e POSTGRES_PASSWORD=taskpass123 \
		-e POSTGRES_DB=taskdb \
		-p 5432:5432 \
		-d \
		postgres:17

docker-down:
	docker stop postgres-17 && docker rm postgres-17

docker-logs:
	docker logs -f postgres-17