postgres:
	docker run --name sushi_roll_postgres -e "POSTGRES_PASSWORD=postgres" -p 5432:5432 -v $HOME/.psql/:/var/lib/postgres -d postgres

createuser:
	docker exec -it sushi_roll_postgres echo "CREATE USER sushi WITH PASSWORD 'roll'" | psql -U postgres

createdb:
	docker exec -it sushi_roll_postgres createdb --username=postgres --owner=sushi sushi_roll_db 

dropdb:
	docker exec -it sushi_roll_postgres dropdb sushi_roll_db

.PHONY: postgres createuser createdb dropdb
