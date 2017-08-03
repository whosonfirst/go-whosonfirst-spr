# go-whosonfirst-spr

Go tools for working Who's On First "standard places responses" (SPR)

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

Too soon. Way way too soon. Move along. Nothing should be considered "stable" yet. If you want to follow along, please consult:

https://github.com/whosonfirst/go-whosonfirst-spr/issues/1

## Interface

_Please finish writing me..._

```
type StandardPlacesResult interface {
	Id() int64
	ParentId() int64
	Name() string
	Placetype() string
	Country() string
	Repo() string
	Path() string
	URI() string
	IsCurrent() bool
	IsCeased() bool
	IsDeprecated() bool
	IsSuperseded() bool
	IsSuperseding() bool
	SupersededBy() []int64
	Supersedes() []int64
}
```

## Tools

### wof-feature-to-spr

Serialize a Who's On First `Feature` as a [WOFStandardPlacesResult](whosonfirst/whosonfirst.go) thingy, which implements the `StandardPlacesResult` interface.

_Please finish writing me..._

```
./bin/wof-feature-to-spr /usr/local/data/whosonfirst-data/data/420/561/633/420561633.geojson | python -mjson.tool
{
    "mz:is_ceased": 1,
    "mz:is_current": 0,
    "mz:is_deprecated": 0,
    "mz:is_superseded": 0,
    "mz:uri": "https://whosonfirst.mapzen.com/data/420/561/633/420561633.geojson",
    "wof:country": "US",
    "wof:id": 420561633,
    "wof:name": "Super Bowl City",
    "wof:parent_id": 85865899,
    "wof:path": "420/561/633/420561633.geojson",
    "wof:placetype": "microhood",
    "wof:repo": "whosonfirst-data",
    "wof:superseded_by": [],
    "wof:supersedes": []
}

./bin/wof-feature-to-spr /usr/local/data/whosonfirst-data/data/856/326/09/85632609.geojson | python -mjson.tool
{
    "mz:is_ceased": 0,
    "mz:is_current": -1,
    "mz:is_deprecated": 0,
    "mz:is_superseded": 0,
    "mz:is_superseding": 1,
    "mz:uri": "https://whosonfirst.mapzen.com/data/856/326/09/85632609.geojson",
    "wof:country": "BA",
    "wof:id": 85632609,
    "wof:name": "Bosnia and Herzegovina",
    "wof:parent_id": 102191581,
    "wof:path": "856/326/09/85632609.geojson",
    "wof:placetype": "country",
    "wof:repo": "whosonfirst-data",
    "wof:superseded_by": [],
    "wof:supersedes": [
        1108955785
    ]
}
```

## Background

_Please write me..._

* https://code.flickr.net/2008/08/19/standard-photos-response-apis-for-civilized-age/
* https://code.flickr.net/2008/08/25/api-responses-as-feeds/

## See also

* https://github.com/whosonfirst/go-whosonfirst-geojson-v2
