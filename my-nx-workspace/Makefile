
build-api:
	nx run api:build
	npx ncc build dist/apps/api/main.js -o dist/apps/api-ncc/
	docker build -f apps/api/Dockerfile .

