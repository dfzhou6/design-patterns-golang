package chain

import "fmt"

type IFilter interface {
	Filter(content string) error
}

type FilterList struct {
	filterList []IFilter
}

func (f *FilterList) Add(filter IFilter) {
	f.filterList = append(f.filterList, filter)
}

func (f *FilterList) Filter(content string) error {
	var err error

	for _, filter := range f.filterList {
		if err = filter.Filter(content); err != nil {
			return err
		}
	}

	return nil
}

type WordFilter struct {
}

func (*WordFilter) Filter(content string) error {
	fmt.Println("word filter", content)
	return nil
}

type LenFilter struct {
}

func (*LenFilter) Filter(content string) error {
	fmt.Println("len filter", content)
	return nil
}
