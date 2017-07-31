package whosonfirst

import (
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	wof "github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
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
	path          string  `json:"wof:path"`
	uri           string  `json:"mz:uri"`
	superseded_by []int64 `json:wof:superseded_by"`
	supersedes    []int64 `json:wof:supersedes"`
	is_current    bool    `json:"mz:is_current"`
	is_ceased     bool    `json:"mz:is_ceased"`
	is_deprecated bool    `json:"mz:is_deprecated"`
	is_superseded bool    `json:"mz:is_superseded"`
}

func NewSPRFromFeature(f geojson.Feature) (spr.StandardPlacesResult, error) {

	err := feature.EnsureWOFFeature(f.Bytes())

	if err != nil {
		return nil, err
	}

	id := wof.Id(f)
	parent_id := wof.ParentId(f)
	name := wof.Name(f)
	country := wof.Country(f)
	repo := wof.Repo(f)

	path, err := uri.Id2RelPath(id)

	if err != nil {
	   return nil, err
	}

	uri, err := uri.Id2AbsPath("https://whosonfirst.mapzen.com/data", id)

	if err != nil {
	   return nil, err
	}

	is_current, _ := wof.IsCurrent(f)

	is_ceased := wof.IsCeased(f)
	is_deprecated := wof.IsDeprecated(f)
	is_superseded := wof.IsSuperseded(f)

	// FIX ME
	superseded_by := make([]int64, 0)
	supersedes := make([]int64, 0)

	spr := WOFStandardPlacesResult{
		id:            id,
		parent_id:     parent_id,
		name:          name,
		country:       country,
		repo:          repo,
		path:          path,
		uri:           uri,
		is_current:    is_current,
		is_ceased:     is_ceased,
		is_deprecated: is_deprecated,
		is_superseded: is_superseded,
		supersedes:    supersedes,
		superseded_by: superseded_by,
	}

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
	return spr.path
}

func (spr *WOFStandardPlacesResult) URI() string {
     	  return spr.uri
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
