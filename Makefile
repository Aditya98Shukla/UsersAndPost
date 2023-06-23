
step_1:
	go mod init example

# define download location for 3rd party modules
step_2:
	go env -w GOMODCACHE=$(PWD)/.deps/pkg/mod

# downloading modules in order to import gin
step_3:
	go mod tidy

run:
	go run main.go

getUsers:
	curl localhost:8082/users

getPosts:
	curl localhost:8082/posts

getBoth:
	curl localhost:8082/usersAndPost/$(ID)

build:
	go build -o app main.go