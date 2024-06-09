# Путь до бинарных файлов
LOCAL_BIN=$(CURDIR)/bin

# Установка бинарных файлов различных пакетов.
# Плагины пакетов можно установить глобально выполнив команду go install [package].
# Но лучше сохранить их в каталоге проекта, зафиксировав используемые версии,
# чтобы не нарушить обратную совместимость при коллективной работе над проектом.
install:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.32.2
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.20.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.20.0
clean:
	rm -rf $(LOCAL_BIN)
# Скачивание proto файлов для того, что бы IDE не ругалась:
	curl https://raw.githubusercontent.com/googleapis/googleapis/974ad5bdfc9ba768db16b3eda2850aadd8c10a2c/google/api/annotations.proto --create-dirs -o proto/google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/974ad5bdfc9ba768db16b3eda2850aadd8c10a2c/google/api/http.proto --create-dirs -o proto/google/api/http.proto
	curl https://raw.githubusercontent.com/bufbuild/protovalidate/main/proto/protovalidate/buf/validate/validate.proto --create-dirs -o proto/buf/validate/validate.proto
	curl https://raw.githubusercontent.com/bufbuild/protovalidate/main/proto/protovalidate/buf/validate/expression.proto --create-dirs -o proto/buf/validate/expression.proto
	curl https://raw.githubusercontent.com/bufbuild/protovalidate/main/proto/protovalidate/buf/validate/priv/private.proto --create-dirs -o proto/buf/validate/priv/private.proto