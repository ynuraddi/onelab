migration_up: confirm
	@echo 'Running up migrations...'
	migrate -path=./internal/repository/postgre/migrations -database=postgres://onelab:onelab@localhost:5432/onelab?sslmode=disable up

run/db:
	@echo 'Running up container with database...'
	docker run --rm --name psg -e POSTGRES_USER=onelab -e POSTGRES_PASSWORD=onelab -e POSTGRES_DATABASE=onelab -d -p 5432:5432 postgres
	sleep 5 && cat onelab_dump.sql | docker exec -i psg psql -U onelab -d onelab

stop/db:
	@echo 'Dumping up container with database...'
	docker exec -i psg pg_dump -U onelab onelab > onelab_dump.sql
	@echo 'Stopping up container with database...'
	docker stop psg

run/app:
	go run ./cmd

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]