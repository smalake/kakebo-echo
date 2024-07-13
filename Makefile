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

test:
	@$(EXEC_APP) go mod tidy -go=1.21
	@$(EXEC_APP) go test -v -coverprofile=cover.out ./...
	@$(EXEC_APP) go tool cover -html=cover.out -o cover.html

gen-api:
	./tools/codegen.sh

gen-mock:
	sh tools/mockgen.sh $(src)