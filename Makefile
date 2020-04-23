default:
	@echo "build"
	docker build -t still_up .

up: default
	@echo "start docker-compose"
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f stillUp
	docker system prune -f
	docker volume prune -f