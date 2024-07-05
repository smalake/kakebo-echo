TAGS := "auth event private pattern setting health"
EXEC_APP := docker-compose exec -T golang

up:
	docker-compose up -d
stop:
	docker-compose stop
down:
	docker-compose down
logs:
	docker-compose logs

gen-api:
	./tools/codegen.sh

gen-mock:
	sh tools/mockgen.sh $(src)