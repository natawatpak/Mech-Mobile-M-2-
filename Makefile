local-backend:
	go run main/server.go

local-graphql:
	go run backend/server.go

local-frontend:
	cd frontend;yarn serve

local-frontend-mech:
	cd frontend\ mech;yarn serve

local-DB:
	docker run --name postgres -e POSTGRES_PASSWORD=Eauu0244 -p 5432:5432 -d postgres
