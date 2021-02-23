dockerProduction:
	docker build ./ -t gorest --build-arg app_env=production

dockerDevelopment:
	docker build ./ -t gorest

run:
	docker run -i -t -p 8080:8080 gorest:latest

all-production: dockerProduction run
all-dev: dockerDevelopment run