logs:
	docker-compose logs -f

default:
	@echo "=============building Local API============="
	docker build -f Dockerfile -t gorest:latest .

up: default
	@echo "=============starting locally============="
	docker-compose up -d

delete:
	docker rmi gorest:latest

test:
	go test -v -cover ./...

down:
	docker-compose down

clean: down
	@echo "=============cleaning up============="
	rm -f gorest
	docker system prune -f
	docker volume prune -f

run: up logs