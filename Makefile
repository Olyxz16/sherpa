include .env

.PHONY: diff migrate

diff:
	cd infrastructure/persistence && atlas migrate diff --to file://schema.sql --dev-url docker://postgres/17/dev

migrate:
	cd infrastructure/persistence && atlas migrate apply --dir file://migrations --url postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable

upload-artifacts:
	@if ["$(VERSION)" = ""] ; then \
		echo "ERROR : VERSION variable not set !" || exit 1 ; \
	fi
	docker build -t sherpa-server -f Dockerfile .
	docker tag sherpa-server ghcr.io/olyxz16/sherpa/server:$(VERSION)
	docker push ghcr.io/olyxz16/sherpa/server:$(VERSION)
	docker tag sherpa-server ghcr.io/olyx16/sherpa/server:latest
	docker push ghcr.io/olyxz16/sherpa/server:latest
	docker build -t sherpa-migrations -f Dockerfile.migrate .
	docker tag sherpa-migrations ghcr.io/olyxz16/sherpa/migrations:$(VERSION)
	docker push ghcr.io/olyxz16/sherpa/migrations:$(VERSION)
	docker tag sherpa-migrations ghcr.io/olyx16/sherpa/migrations:latest
	docker push ghcr.io/olyxz16/sherpa/migrations:latest
