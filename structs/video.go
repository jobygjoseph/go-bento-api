package structs

import "time"

// Video struct
// type Video struct {
// 	AlternateTeaseImages    string `json:"title"`
// 	AssociatedVideoPlaylist string `json:"title2"`
// 	ClosedCaptioning        string `json:"title4"`
// 	DateAvailable           string `json:"title5"`
// 	DateBroadcast           string `json:"title6"`
// 	DateCreated             string `json:"title7"`
// 	DateModified            string `json:"title8"`
// 	DatePublished           string `json:"title9"`
// 	DefaultAssociation      string `json:"title10"`
// 	Description             string `json:"title11"`
// 	DocumentTracing         string `json:"title12"`
// 	Duration                string `json:"title13"`
// 	Flag                    string `json:"title15"`
// 	HasCaptions             string `json:"title16"`
// 	Headline                string `json:"title17"`
// 	Hidden                  string `json:"title18"`
// 	ID                      string `json:"_id"`
// 	LegacyData              string `json:"title20"`
// 	LiveVideoStatus         string `json:"title21"`
// 	MpxMetadata             string `json:"title22"`
// 	Playable                string `json:"title24"`
// 	PrimaryImage            string `json:"title25"`
// 	Publisher               string `json:"title26"`
// 	RelatedURL              string `json:"title27"`
// 	Searchable              string `json:"title28"`
// 	Sentiment               string `json:"title29"`
// 	Source                  string `json:"title30"`
// 	Taxonomy                string `json:"title31"`
// 	TeaseImage              string `json:"title32"`
// 	Type                    string `json:"title37"`
// 	Unibrow                 string `json:"title33"`
// 	URL                     string `json:"title34"`
// 	VideoAssets             string `json:"title35"`
// }

type Video struct {
	AutoCuration bool `bson:",omitempty"`
	Expires      time.Time
	ID           string `bson:"_id"`
	NativeAd     bool
	VideoType    string
}
