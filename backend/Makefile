run:
	go run *.go

main:
	go run server.go

test:
	go run server_test.go

gen:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

server:
	go run main.go

backup:
	cd graph/model && ls && cp models_gen.go backup