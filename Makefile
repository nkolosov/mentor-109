gen_proto:
	docker run --rm -v "$(pwd):/work" docker.citik.ru/base/uber-prototool:latest chown -R "$(id -u)":"$(id -g)" "/work/category/internal/grpc/gen"