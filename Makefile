goa-gen-linux:
	go install goa.design/goa/v3/cmd/goa@v3
	goa gen github.com/TsubasaKanemitsu/golang-todo-app/backend/design

goa-example:
	goa example github.com/TsubasaKanemitsu/golang-todo-app/backend/design
	go build ./cmd/todo && go build ./cmd/todo-cli/

setup-example-server:
	./todo

setup-db:
	docker compose up -d psql
down-db:
	docker compose down