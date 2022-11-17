package sort

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"sort"
)

func init(){
	ctx := context.Background()
	RegisterSorter(ctx, "inception", NewInceptionSorter)
}

type byInception []spr.StandardPlacesResult

func (s byInception) Len() int {
	return len(s)
}

func (s byInception) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byInception) Less(i, j int) bool {

	i_inception := s[i].Inception()
	j_inception := s[j].Inception()

	if i_inception.String() == "" {
		return false
	}

	if j_inception.String() == "" {
		return true
	}

	is_before, err := i_inception.Before(j_inception)

	if err != nil {
		return false
	}

	return is_before
}

type InceptionSorter struct {
	Sorter
}

func NewInceptionSorter(ctx context.Context, uri string) (Sorter, error) {
	s := &InceptionSorter{}
	return s, nil
}

func (s *InceptionSorter) Sort(ctx context.Context, results spr.StandardPlacesResults) (spr.StandardPlacesResults, error) {

	to_sort := results.Results()
	sort.Sort(byInception(to_sort))

	sorted := &SortedStandardPlacesResults{
		results: to_sort,
	}

	return sorted, nil
}
