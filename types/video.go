package types

import (
	"github.com/graphql-go/graphql"
)

var Video = graphql.NewObject(graphql.ObjectConfig{
	Name: "Video",
	Fields: graphql.Fields{
		// "autoCuration": &graphql.Field{
		// 	Type: graphql.Boolean,
		// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// 		// fmt.Println(params)
		// 		vid := (params.Source).(structs.Video)
		// 		autoCuration := vid.AutoCuration
		// 		if autoCuration {
		// 			return autoCuration, nil
		// 		}

		// 		// customFields := video.customFields;

		// 		// if (!Array.isArray(customFields) || customFields.length === 0) {
		// 		// 	return null;
		// 		// }

		// 		// const { value: autoCuration } = customFields.find(ele => ele.key === 'autoCuration') || {};
		// 		// return !isNil(autoCuration) ? autoCuration : null;
		// 		return structs.Video{}, nil
		// 	},
		// },
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
