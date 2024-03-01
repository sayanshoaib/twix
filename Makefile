 gql:
	go get github.com/99designs/gqlgen/internal/imports@v0.17.10
	go get github.com/99designs/gqlgen@v0.17.10
	go run github.com/99designs/gqlgen generate

ent:
	go generate ./ent
