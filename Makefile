logs:
	docker-compose logs -f

default:
	@echo "=============building Local API============="
	docker build -f Dockerfile -t gorest:latest .

up:
	@echo "=============starting locally============="
	docker-compose up -d

delete:
	docker rmi gorest

test:
	go test -v -cover ./...

down:
	docker-compose down

clean: down
	@echo "=============cleaning up============="
	docker rmi $(docker images -q)

rmall:
	docker system prune -f -a

ls:
	docker image ls

run: up logs