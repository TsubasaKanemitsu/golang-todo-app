## sample
- ADD
```bash
./todo-cli todoservice add-todo-task --body '{
          "asignee": "Kanemitsu",                                  
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