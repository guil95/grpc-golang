proto-generate: ## Generate protocol buffer file from proto
	docker run --rm \
		--name rickymorty-proto \
		-u $(id -u):$(id -u) \
		-v `pwd`:/code \
		rvolosatovs/protoc:3.2 \
			--proto_path=/code/api \
			--go_out=plugins=grpc,paths=source_relative:/code/api/genpb /code/api/*.proto