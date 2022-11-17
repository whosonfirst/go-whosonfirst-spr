package sort

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"sort"
)

func init() {
	ctx := context.Background()
	RegisterSorter(ctx, "name", NewNameSorter)
}

type NameSorter struct {
	Sorter
}

func NewNameSorter(ctx context.Context, uri string) (Sorter, error) {
	s := &NameSorter{}
	return s, nil
}

func (s *NameSorter) Sort(ctx context.Context, results spr.StandardPlacesResults, follow_on_sorters ...Sorter) (spr.StandardPlacesResults, error) {

	lookup := make(map[string][]spr.StandardPlacesResult)

	for _, s := range results.Results() {

		_results, ok := lookup[s.Name()]

		if !ok {
			_results = make([]spr.StandardPlacesResult, 0)
		}

		_results = append(_results, s)
		lookup[s.Name()] = _results

	}

	names := make([]string, 0)

	for n, _ := range lookup {
		names = append(names, n)
	}

	sort.Strings(names)

	sorted := make([]spr.StandardPlacesResult, 0)

	for _, n := range names {

		for _, s := range lookup[n] {
			sorted = append(sorted, s)
		}
	}

	count_follow_on := len(follow_on_sorters)

	if count_follow_on == 0 {

		sorted_results := &SortedStandardPlacesResults{
			results: sorted,
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

	last_name := ""

	doNextSort := func(name string) error {

		_results, _ := tmp[name]

		name_results := &SortedStandardPlacesResults{
			results: _results,
		}

		name_sorted, err := next_sorter.Sort(ctx, name_results, other_sorters...)

		if err != nil {
			return fmt.Errorf("Failed to apply next sorter to name '%s', %w", name, err)
		}

		for _, name_s := range name_sorted.Results() {
			final = append(final, name_s)
		}

		return nil
	}

	for _, s := range sorted {

		name := s.Name()

		if name != last_name {

			if last_name != "" {

				err := doNextSort(last_name)

				if err != nil {
					return nil, fmt.Errorf("Failed to perform next sort for %s, %w", name, err)
				}
			}

			last_name = name
		}

		_results, ok := tmp[name]

		if !ok {
			_results = make([]spr.StandardPlacesResult, 0)
		}

		_results = append(_results, s)
		tmp[name] = _results
	}

	err := doNextSort(last_name)

	if err != nil {
		return nil, fmt.Errorf("Failed to perform next sort for %s, %w", last_name, err)
	}

	sorted_results := &SortedStandardPlacesResults{
		results: final,
	}

	return sorted_results, nil
}
