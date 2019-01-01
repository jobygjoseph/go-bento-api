package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-bento-api/schemas"
	"github.com/go-bento-api/store"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	APIstore := store.CreateStore()
	//fmt.Println(APIstore.Video.FindByID("mmvo16656963832"))
	// fmt.Println(APIstore.Video.FindByID("mmvo16656963832"))

	description := "![NBC News](http://sslnodeassets.nbcnews.com/cdnassets/projects/site-images/nbcnews-logo-white.png \"NBC NEWS\")\n\nThe Graphiql IDE for Bento API"

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

	handleGraphqlRequest := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), "APIstore", APIstore)
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		Query := r.FormValue("query")
		//Variables := r.FormValue("variables")
		OperationName := r.FormValue("operationName")
		params := graphql.Params{
			RequestString: Query,
			//VariableValues: Variables,
			OperationName: OperationName,
			Context:       ctx,
			Schema:        schema}
		result := graphql.Do(params)
		if len(result.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
		}
		rJSON, _ := json.Marshal(result)
		fmt.Fprintf(w, string(rJSON))
		//w.Write(string(rJSON))
	}

	//http.HandleFunc()
	http.Handle("/graphql", h)
	http.HandleFunc("/", handleGraphqlRequest)
	http.ListenAndServe(":8080", nil)
}
