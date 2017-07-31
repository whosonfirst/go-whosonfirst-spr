package whosonfirst

import (
       "errors"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/geojson"
	"github.com/whosonfirst/go-whosonfirst-spr"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

type WOFStandardPlacesResult struct {
	spr.StandardPlacesResult
	id            int64   `json:"wof:id"`
	parent_id     int64   `json:"wof:parent_id"`
	name          string  `json:"wof:name"`
	country       string  `json:"wof:country"`
	repo          string  `json:"wof:repo"`
	path          string  `json:"xx:path"`
	uri           string  `json:"xx:uri"`
	superseded_by []int64 `json:wof:superseded_by"`
	supersedes    []int64 `json:wof:supersedes"`
	is_current    bool    `json:"mz:is_current"`
	is_ceased     bool    `json:"mz:is_ceased"`
	is_deprecated bool    `json:"mz:is_deprecated"`
	is_superseded bool    `json:"mz:is_superseded"`
}

func NewSPRFromFeature(f geojson.Feature) (spr.StandardPlacesResult, error) {

     return nil, errors.New("please write me...")

	spr := WOFStandardPlacesResult{}

	return &spr, nil
}

func (spr *WOFStandardPlacesResult) Id() int64 {
	return spr.id
}

func (spr *WOFStandardPlacesResult) ParentId() int64 {
	return spr.parent_id
}

func (spr *WOFStandardPlacesResult) Name() string {
	return spr.name
}

func (spr *WOFStandardPlacesResult) Country() string {
	return spr.country
}

func (spr *WOFStandardPlacesResult) Repo() string {
	return spr.repo
}

func (spr *WOFStandardPlacesResult) Path() string {

	rel_path := uri.Id2RelPath(spr.Id())
	return rel_path
}

func (spr *WOFStandardPlacesResult) URI() string {

	abs_path := uri.Id2AbsPath("https://whosonfirst.mapzen.com/data", spr.Id())
	return abs_path
}

func (spr *WOFStandardPlacesResult) IsCurrent() bool {
	return spr.is_current
}

func (spr *WOFStandardPlacesResult) IsCeased() bool {
	return spr.is_ceased
}

func (spr *WOFStandardPlacesResult) IsDeprecated() bool {
	return spr.is_deprecated
}

func (spr *WOFStandardPlacesResult) IsSuperseded() bool {
	return spr.is_superseded
}

func (spr *WOFStandardPlacesResult) SupersededBy() []int64 {
	return spr.superseded_by
}

func (spr *WOFStandardPlacesResult) Supersedes() []int64 {
	return spr.supersedes
}
