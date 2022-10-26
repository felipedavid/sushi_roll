postgres:
	mkdir $$HOME/.psql
	docker run --name sushi_roll_postgres -e "POSTGRES_PASSWORD=postgres" -p 5432:5432 -v $$HOME/.psql/:/var/lib/postgres -d postgres

createuser:
	docker exec -it sushi_roll_postgres psql -U postgres -c "CREATE USER sushi WITH PASSWORD 'roll'"

createdb:
	docker exec -it sushi_roll_postgres createdb --username=postgres --owner=sushi sushi_roll_db 

dropdb:
	docker exec -it sushi_roll_postgres dropdb sushi_roll_db

sonar:
	go test ./cmd/web -coverprofile=coverage.out && \
		docker run --rm -e SONAR_HOST_URL="http://labens.dct.ufrn.br/sonarqube" -e SONAR_LOGIN="$(SONAR_TOKEN)" -v "$(PWD):/usr/src" sonarsource/sonar-scanner-cli && \
		rm ./coverage.out


dsn = "postgresql://postgres:postgres@localhost:5432/sushi_roll_db?sslmode=disable"

migrateup:
	migrate -path db/migration -database $(dsn) -verbose up

migratedown:
	migrate -path db/migration -database $(dsn) -verbose down

run:
	go run ./cmd/web

.PHONY: postgres createuser createdb dropdb migrateup migratedown run sonar
