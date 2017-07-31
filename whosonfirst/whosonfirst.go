package whosonfirst

import (
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	wof "github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-spr"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

type WOFStandardPlacesResult struct {
	spr.StandardPlacesResult `json:",omitempty"`
	WOFId                    int64   `json:"wof:id"`
	WOFParentId              int64   `json:"wof:parent_id"`
	WOFName                  string  `json:"wof:name"`
	WOFPlacetype             string  `json:"wof:placetype"`
	WOFCountry               string  `json:"wof:country"`
	WOFRepo                  string  `json:"wof:repo"`
	WOFPath                  string  `json:"wof:path"`
	MZURI                    string  `json:"mz:uri"`
	WOFSupersededBy          []int64 `json:"wof:superseded_by"`
	WOFSupersedes            []int64 `json:"wof:supersedes"`
	MZIsCurrent              int     `json:"mz:is_current"`
	MZIsCeased               int     `json:"mz:is_ceased"`
	MZIsDeprecated           int     `json:"mz:is_deprecated"`
	MZIsSuperseded           int     `json:"mz:is_superseded"`
}

func NewSPRFromFeature(f geojson.Feature) (spr.StandardPlacesResult, error) {

	err := feature.EnsureWOFFeature(f.Bytes())

	if err != nil {
		return nil, err
	}

	id := wof.Id(f)
	parent_id := wof.ParentId(f)
	name := wof.Name(f)
	placetype := wof.Placetype(f)
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

	is_current := 0
	is_ceased := 0
	is_deprecated := 0
	is_superseded := 0

	_current, _ := wof.IsCurrent(f)

	if _current {
		is_current = 1
	}

	if wof.IsCeased(f) {
		is_ceased = 1
	}

	if wof.IsDeprecated(f) {
		is_deprecated = 1
	}

	if wof.IsSuperseded(f) {
		is_superseded = 1
	}

	superseded_by := wof.SupersededBy(f)
	supersedes := wof.Supersedes(f)

	spr := WOFStandardPlacesResult{
		WOFId:           id,
		WOFParentId:     parent_id,
		WOFPlacetype:    placetype,
		WOFName:         name,
		WOFCountry:      country,
		WOFRepo:         repo,
		WOFPath:         path,
		MZURI:           uri,
		MZIsCurrent:     is_current,
		MZIsCeased:      is_ceased,
		MZIsDeprecated:  is_deprecated,
		MZIsSuperseded:  is_superseded,
		WOFSupersedes:   supersedes,
		WOFSupersededBy: superseded_by,
	}

	return &spr, nil
}

func (spr *WOFStandardPlacesResult) Id() int64 {
	return spr.WOFId
}

func (spr *WOFStandardPlacesResult) ParentId() int64 {
	return spr.WOFParentId
}

func (spr *WOFStandardPlacesResult) Name() string {
	return spr.WOFName
}

func (spr *WOFStandardPlacesResult) Country() string {
	return spr.WOFCountry
}

func (spr *WOFStandardPlacesResult) Repo() string {
	return spr.WOFRepo
}

func (spr *WOFStandardPlacesResult) Path() string {
	return spr.WOFPath
}

func (spr *WOFStandardPlacesResult) URI() string {
	return spr.MZURI
}

func (spr *WOFStandardPlacesResult) IsCurrent() bool {

	if spr.MZIsCurrent == 1 {
		return true
	}

	return false
}

func (spr *WOFStandardPlacesResult) IsCeased() bool {

	if spr.MZIsCeased == 1 {
		return true
	}

	return false
}

func (spr *WOFStandardPlacesResult) IsDeprecated() bool {

	if spr.MZIsDeprecated == 1 {
		return true
	}

	return false
}

func (spr *WOFStandardPlacesResult) IsSuperseded() bool {

	if spr.MZIsSuperseded == 1 {
		return true
	}

	return false
}

func (spr *WOFStandardPlacesResult) SupersededBy() []int64 {
	return spr.WOFSupersededBy
}

func (spr *WOFStandardPlacesResult) Supersedes() []int64 {
	return spr.WOFSupersedes
}
