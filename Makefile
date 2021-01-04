CONTAINER_API=book-shelf-server_api_1
CONTAINER_DB=book-shelf-server_db_1

up:
	docker-compose up
exec:
	docker exec -it $(CONTAINER_DB) /bin/bash -c 'mysql -u user -p'
stop:
	docker-compose stop
