package types

import (
	"fmt"

	"github.com/go-bento-api/structs"
	"github.com/graphql-go/graphql"
)

var Video = graphql.NewObject(graphql.ObjectConfig{
	Name: "Video",
	Fields: graphql.Fields{
		"autoCuration": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				fmt.Println(params)
				vid := structs.Video(params.Source)
				// customFields := video.customFields;
				// if vid.AutoCuration != nil {
				// 	return vid.AutoCuration
				// }

				// if (!Array.isArray(customFields) || customFields.length === 0) {
				// 	return null;
				// }

				// const { value: autoCuration } = customFields.find(ele => ele.key === 'autoCuration') || {};
				// return !isNil(autoCuration) ? autoCuration : null;
				return structs.Video{}, nil
			},
		},
		"expires": &graphql.Field{
			Type: graphql.DateTime,
		},
		"ID": &graphql.Field{
			Type: graphql.String,
		},
		"nativeAd": &graphql.Field{
			Type: graphql.Boolean,
		},
		"videoType": &graphql.Field{
			Type: graphql.String,
		},
	},
})
