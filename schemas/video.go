package schemas

import (
	"github.com/go-bento-api/store"
	"github.com/go-bento-api/types"
	"github.com/graphql-go/graphql"
)

var VideoSchema = &graphql.Field{
	Type: types.Video,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.ID),
			Description: "id of the video (e.g. `mmvo191641155794` or `mmvo225272387839`)",
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//APIstore := params.Context.Value("APIstore")
		APIstore := store.CreateStore()
		vidID := params.Args["id"].(string)
		//vidStore := &store.Video{}
		//return vidStore.FindByID(vidID), nil
		return APIstore.Video.FindByID(vidID), nil
	},
}
