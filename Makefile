TAGS := "auth event private pattern setting health"

up:
	docker-compose up -d
stop:
	docker-compose stop
down:
	docker-compose down
logs:
	docker-compose logs

gen-api:
	./tools/codegen.sh ${TAGS}