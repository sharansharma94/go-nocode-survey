### Commands 

GET
```bash 
 curl localhost:8080/forms
```

POST
```bash  curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","gender":"male","age":30}' \
  http://localhost:8080/users

```