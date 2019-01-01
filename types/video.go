package types

import (
	"github.com/graphql-go/graphql"
)

var Video = graphql.NewObject(graphql.ObjectConfig{
	Name: "Video",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.String,
		},
		"VideoType": &graphql.Field{
			Type: graphql.String,
		},
	},
})
