package sort

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"sort"
)

func init() {
	ctx := context.Background()
	RegisterSorter(ctx, "placetype", NewPlacetypeSorter)
}

type byPlacetype []spr.StandardPlacesResult

func (s byPlacetype) Len() int {
	return len(s)
}

func (s byPlacetype) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byPlacetype) Less(i, j int) bool {

	i_pt, err := placetypes.GetPlacetypeByName(s[i].Placetype())

	if err != nil {
		return false
	}

	j_pt, err := placetypes.GetPlacetypeByName(s[j].Placetype())

	if err != nil {
		return false
	}

	return placetypes.IsDescendant(i_pt, j_pt)
}

type PlacetypeSorter struct {
	Sorter
}

func NewPlacetypeSorter(ctx context.Context, uri string) (Sorter, error) {
	s := &PlacetypeSorter{}
	return s, nil
}

func (s *PlacetypeSorter) Sort(ctx context.Context, results spr.StandardPlacesResults, follow_on_sorters ...Sorter) (spr.StandardPlacesResults, error) {

	to_sort := results.Results()
	sort.Sort(byPlacetype(to_sort))

	count_follow_on := len(follow_on_sorters)

	if count_follow_on == 0 {

		sorted_results := &SortedStandardPlacesResults{
			results: to_sort,
		}

		return sorted_results, nil
	}

	next_sorter := follow_on_sorters[0]
	var other_sorters []Sorter

	if count_follow_on > 1 {
		other_sorters = follow_on_sorters[1:]
	}

	tmp := make(map[string][]spr.StandardPlacesResult)
	final := make([]spr.StandardPlacesResult, 0)

	last_placetype := ""

	doNextSort := func(pt string) error {

		_results, _ := tmp[pt]

		pt_results := &SortedStandardPlacesResults{
			results: _results,
		}

		pt_sorted, err := next_sorter.Sort(ctx, pt_results, other_sorters...)

		if err != nil {
			return fmt.Errorf("Failed to apply next sorter to placetype '%s', %w", pt, err)
		}

		for _, pt_s := range pt_sorted.Results() {
			final = append(final, pt_s)
		}

		return nil
	}

	for _, s := range to_sort {

		pt := s.Placetype()

		if pt != last_placetype {

			if last_placetype != "" {

				err := doNextSort(last_placetype)

				if err != nil {
					return nil, fmt.Errorf("Failed to perform next sort for %s, %w", pt, err)
				}
			}

			last_placetype = pt
		}

		_results, ok := tmp[pt]

		if !ok {
			_results = make([]spr.StandardPlacesResult, 0)
		}

		_results = append(_results, s)
		tmp[pt] = _results
	}

	err := doNextSort(last_placetype)

	if err != nil {
		return nil, fmt.Errorf("Failed to perform next sort for %s, %w", last_placetype, err)
	}

	sorted_results := &SortedStandardPlacesResults{
		results: final,
	}

	return sorted_results, nil
}
