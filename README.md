# simple clean architecture
A simple web api in golang with clean architecture

## Run application
```
$ make build
$ ./bin/simple
$ curl "localhost:8000/create" \
	-H "Content-Type: application/json" \
	-d '{"user_id": "001", "first_name": "Taro", "last_name": "Yamada"}'
{"message":"Update success"}
$ curl "localhost:8000/get?user_id=001"
{"FirstName":"Taro","LastName":"Yamada"}
```

## Run tests
```
$ make lint
$ make test
```