run:
	go run ./cmd/api

migrateup:
	migrate -path=./migrations -database=$$SUSHI_ROLL_DB_DSN up

migratedown:
	migrate -path=./migrations -database=$$SUSHI_ROLL_DB_DSN down
