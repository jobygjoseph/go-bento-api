package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/go-bento-api/types"
	"github.com/go-bento-api/store"
)

VideoSchema := &graphql.Field{
	Type: types.Video,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
			Description: "id of the video (e.g. `mmvo191641155794` or `mmvo225272387839`)",
		}
	}
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		vidID := params.Args["id"].(string)
		return store.Video.FindByID(vidID), nil
	}
}