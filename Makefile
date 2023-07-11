goa-gen-linux:
	go install goa.design/goa/v3/cmd/goa@v3
	goa gen github.com/TsubasaKanemitsu/golang-todo-app/design

goa-example:
	goa example github.com/TsubasaKanemitsu/golang-todo-app/design
	go build ./cmd/todo && go build ./cmd/todo-cli/
setup-example-server:
	./todo
help-example-server:
	./todo-cli --help
setup-db:
	docker compose up -d psql
down-db:
	docker compose down
psql:
	docker compose exec psql psql -U todo -d testdb

sqlboiler:
	sqlboiler psql --config ./sqlboiler.toml --pkgname dbmodels --wipe --no-hooks --struct-tag-casing camel --output ./pkg/dbmodels
gooseup:
	goose -dir migration/ddl postgres "user=todo dbname=testdb password=test sslmode=disable" up
