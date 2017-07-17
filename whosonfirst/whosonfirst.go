package whosonfirst

import (
       "github.com/whosonfirst/go-whosonfirst-geojson-v2/geojson"
       "github.com/whosonfirst/go-whosonfirst-spr"
)

type WOFStandardPlacesResult struct {
     spr.StandardPlacesResult
     id	int64 `json:"wof:id"`
     parent_id int64 `json:"wof:parent_id`"
     name string `json:"wof:name"`
     country string `json:"wof:country"`
     repo string `json:"wof:repo"`
     path string `json:"xx:path"`
     uri string `json:"xx:uri"`     
}

func NewSPRFromFeature(f geojson.Feature) (spr.StandardPlacesResult, error){

}
