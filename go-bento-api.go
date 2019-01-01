package main

import (
	"net/http"

	"github.com/go-bento-api/schemas"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	//APIstore := store.CreateStore()
	//fmt.Println(APIstore.Video.FindByID("mmvo16656963832"))
	// fmt.Println(APIstore.Video.FindByID("mmvo16656963832"))

	description := "![NBC News](http://sslnodeassets.nbcnews.com/cdnassets/projects/site-images/nbcnews-logo-white.png \"NBC NEWS\")The Graphiql IDE for the Bento API"

	// Schema
	fields := graphql.Fields{"Video": schemas.VideoSchema}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields, Description: description}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, _ := graphql.NewSchema(schemaConfig)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
