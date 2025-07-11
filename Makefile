migrateup:
	migrate -path DB/migrations -database "postgresql://myfoodsapp:mysecretpassword@localhost:5432/foodsapp?sslmode=disable" -verbose up

migratedown:
	migrate -path DB/migrations -database "postgresql://myfoodsapp:mysecretpassword@localhost:5432/foodsapp?sslmode=disable" -verbose down

sqlc:
	sqlc generate

run:
	go run main.go

.PHONY: migrateup migratedown sqlc run