BINARY=engine

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	docker-compose up -d

stop:
	docker-compose down

remove:
	docker-compose down -v

.PHONY: clean run stop remove
