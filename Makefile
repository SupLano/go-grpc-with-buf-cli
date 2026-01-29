lint-break:
	buf breaking --against 'github.com/supLano/go-grpc-proto/api'

lint-proto:
	buf lint


generate-proto:
	buf generate --template buf.gen.yaml