module github.com/inidaname/mosque/mosques-service

go 1.24.3

require (
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/inidaname/mosque v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.7.4
	github.com/patrickmn/go-cache v2.1.0+incompatible
	golang.org/x/crypto v0.38.0
	google.golang.org/grpc v1.72.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/inidaname/mosque => ../
