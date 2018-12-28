package main

import (
	"fmt"

	"github.com/go-bento-api/store"
)

func main() {

	APIstore := store.CreateStore()
	//fmt.Println(APIstore.Video.FindByID("mmvo16656963832"))
	fmt.Println(APIstore.Video.FindByID("mmvo16656963832"))

	// Schema
	// fields := graphql.Fields{
	// 	"hello": &graphql.Field{
	// 		Type: graphql.String,
	// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 			return "world", nil
	// 		},
	// 	},
	// }
	// rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	// schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	// schema, _ := graphql.NewSchema(schemaConfig)

	// h := handler.New(&handler.Config{
	// 	Schema:   &schema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// })

	// http.Handle("/graphql", h)
	// http.ListenAndServe(":8080", nil)
}
