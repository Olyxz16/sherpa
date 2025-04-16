include .env

.PHONY: diff migrate

diff:
	cd infrastructure/persistence && atlas migrate diff --to file://schema.sql --dev-url docker://postgres/17/dev

migrate:
	cd infrastructure/persistence && atlas migrate apply --dir file://migrations --url postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable
