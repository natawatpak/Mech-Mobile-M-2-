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
	go run serverLocal.go

backup:
	cd graph/model && ls && cp models_gen.go backup

dockerUp:
	docker run --name postgres -e POSTGRES_PASSWORD=Eauu0244 -p 5432:5432 -d postgres


dockerClear:
	docker kill $$(docker ps -q)
	docker rm $$(docker ps -a -q)

deploy:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main
	zip deployment.zip main

client:
	genqlient graph/genqlient.yaml