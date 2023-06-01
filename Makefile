export DSN_MIGRATIONS="./repository/sqlite/migrations"

test:
	go test -v ./...

migrate-create:
	@read -p 'Введите название новой миграции: ' value; \
	migrate create -seq -ext=.sql -dir=$(DSN_MIGRATIONS) $$value