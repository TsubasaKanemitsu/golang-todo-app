## 初期設定
```bash
# postgresqlのDB用のDockerの立ち上げ
make setup-db
# DDLの投入
make gooseup
# build
make build
# サーバの立ち上げ
make setup-example-server
```

## sample
- ADD
```bash
./todo-cli todoservice add-todo-task --body '{
          "asignee":"Kanemitsu",                                  
          "contents": "Qui assumenda ipsum consequatur nisi dolorem.",
          "end_date": "1990-01-16",
          "label": "vuls",
          "start_date": "1977-11-04",
          "status": "not started",
          "title": "Quia error rerum sed."
       }'
```
- GET
```bash
./todo-cli todoservice get-todo-task -id 1
```

- GET List
```bash
./todo-cli todoservice get-todo-task-list
```

- UPDATE
```bash
./todo-cli todoservice update-todo-task --body '{
      "asignee": "Kanemitsu",
      "contents": "Qui assumenda ipsum consequatur nisi dolorem.",
      "end_date": "1991-01-16",
      "label": "Ea et id.",
      "start_date": "1977-11-04",
      "status": "not started",
      "title": "Quia error rerum sed."
   }' --id 2
```

- DELETE
```bash
./todo-cli todoservice delete-todo-task -id 1
```
