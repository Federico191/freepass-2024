include app.env

mysql:
	docker run --name mysql15 -p ${DATABASE_PORT}:${DATABASE_PORT} -e MYSQL_ROOT_PASSWORD=${DATABASE_PASSWORD} -d mysql:8.2

createdb:
	docker exec -it mysql15 mysql -e "CREATE DATABASE ${DATABASE_NAME};" -u ${DATABASE_USERNAME} -p

dropdb:
	docker exec -it mysql15 mysql -e "DROP DATABASE ${DATABASE_NAME};" -u ${DATABASE_USERNAME} -p

migrateup :
	 migrate -database "mysql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:${DATABASE_PORT})/${DATABASE_NAME}?charset=utf8mb4&parseTime=True&loc=Local" -path db/migration up

migratedown :
	migrate -database "mysql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:${DATABASE_PORT})/${DATABASE_NAME}" -path db/migration down

.PHONY: mysql createdb dropdb