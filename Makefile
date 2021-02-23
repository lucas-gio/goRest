dockerBuild:
	docker build ./ -t gorest

dockerProduction: dockerBuild --build-arg app_env=production

dockerDevelopment: dockerBuild

run:
	docker run -i -t -p 8080:8080 gorest:latest -e GOREST_CONFIG="/home/pc/config.cfg"

all-prod: dockerProduction run
all-dev: dockerDevelopment