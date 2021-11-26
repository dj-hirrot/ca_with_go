BINARY=engine

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	docker-compose up -d

stop:
	docker-compose down

remove:
	docker-compose down -v

dbclean:
	rm -rf ./docker/db/mysql/data/*

gswag:
	swag init -g src/infrastructure/router.go

.PHONY: clean run stop remove gswag dbclean
