build:
	docker-compose build

dev: build
	docker-compose -f docker-compose.dev.yml up -d
	docker exec -t alpha-gopher-client bash -c "/app/vue_htop.sh &"

prod: build
	docker-compose up -d
	docker exec -t alpha-gopher-client bash -c "/app/vue_htop.sh &"

lint:
	cd web/alpha-gopher-frontend && \
	./node_modules/eslint/bin/eslint.js --ext .js,.vue src

lint-fix:
	cd web/alpha-gopher-frontend && \
	./node_modules/eslint/bin/eslint.js --ext .js,.vue src --fix
