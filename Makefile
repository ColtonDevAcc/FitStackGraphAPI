## https://github.com/golang-migrate/migrate#cli-usage <====== you need this inorder to migrate with maker file
## for ios its ====> arch -x86_64 brew install golang-migrate
mock:
	mockery --all --keeptree

migrate:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@35.223.54.178:5432/postgres?sslmode=disable up 

rollback:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@35.223.54.178:5432/postgres?sslmode=disable down

drop:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@35.223.54.178:5432/postgres?sslmode=disable drop

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir postgres/migrations $name

run:
	go run cmd/graphqlserver/*.go

generate: 
	go generate ./.. 