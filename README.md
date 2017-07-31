# go-whosonfirst-spr

Go tools for working Who's On First "standard places responses" (SPR)

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

Too soon. Way way too soon. Move along.

## Tools

### wof-feature-to-spr

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
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-geojson-v2