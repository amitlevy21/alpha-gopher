build:
	docker-compose build

dev: build
	docker-compose up

lint:
	cd web/alpha-gopher-frontend && \
	./node_modules/eslint/bin/eslint.js --ext .js,.vue src

lint-fix:
	cd web/alpha-gopher-frontend && \
	./node_modules/eslint/bin/eslint.js --ext .js,.vue src --fix
