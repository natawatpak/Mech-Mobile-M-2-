local-backend:
	go run main/server.go

local-graphql:
	go run backend/server.go

local-frontend:
	cd frontend;yarn && yarn serve

yarn-local-frontend:
	cd frontend;yarn

local-frontend-mech:
	cd frontend\ mech;yarn && yarn serve

yarn-local-frontend-mech:
	cd frontend\ mech;yarn

local-DB:
	docker run --name postgres -e POSTGRES_PASSWORD=Eauu0244 -p 5432:5432 -d postgres
