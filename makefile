include .env

.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yaml up --build

.PHONY: docker-down
docker-down: ## Stop docker containers and clear artefacts.
	docker-compose -f docker-compose.yaml down
	docker system prune 

.PHONY: postman-ci-cd-test
postman-local-test:
	postman login --with-api-key ${POSTMAN_API_KEY}
	postman collection run "postman/collections/calculator.json"
	