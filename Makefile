export DSN_MIGRATIONS="./repostiroy/sqlite/migrations"

test:
	go test -v ./...

migrate-create:
	@read -p 'Введите название новой миграции: ' value; \
	migrate create -seq -ext=.sql -dir=$(DSN_MIGRATIONS) $$value