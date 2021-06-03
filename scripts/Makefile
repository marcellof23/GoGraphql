stopdocker:
	docker stop mariadb-ktp-dbs
removedocker:
	docker rm mariadb-ktp-dbs
runmariadb:
	docker run -d -p 5000:3306 --name mariadb-ktp-dbs mariadb-ktp
execdocker:
	docker exec -it mariadb-ktp-dbs bash 
exdocker:
	docker exec -it mariadb-ktp-dbs mysql -u root --password=123123123 
migrateup:
	migrate -path ../migrations -database "mysql://root:123123123@tcp(localhost:5000)/ktp_db" up
migratedown:
	migrate -path ../migrations -database "mysql://root:123123123@tcp(localhost:5000)/ktp_db" down
run:
	go run ../server.go
.PHONY: mariadb createdb dropdb migrateup migratedown runmariadb