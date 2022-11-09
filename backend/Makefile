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
	go run server.go

backup:
	cd graph/model && ls && cp models_gen.go backup

dockerBuildUP:
	docker-compose up --build

dockerDown:
	docker-compose down

dockerUp:
	docker run --name postgres -e POSTGRES_PASSWORD=Eauu0244 -p 5432:5432 -d postgres


dockerClear:
	docker kill $(docker ps -q)
	docker rm $(docker ps -a -q)