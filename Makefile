include .env

backend-ssh:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh

# DB関連
## マイグレーション
db-migrate:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go run migrate/migrate.go"


# ローカル開発用
# go library install
## 複数のライブラリを指定する場合は、name="xxx yyy" のように""で囲んで実行すること
go-get:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go get ${name}"
## テスト
test:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go test -v ./..."
lint:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "staticcheck ./..."