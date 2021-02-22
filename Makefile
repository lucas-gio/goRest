dockerProduction:
	docker build ./ --build-arg app_env=production

dockerDevelopment:
	docker build ./

run:
	docker run -i -t -p 8080:8080 gorest

all-production: dockerProduction run
all-dev: dockerDevelopment run