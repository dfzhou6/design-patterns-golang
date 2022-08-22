package chain

import "testing"

func TestFilterList_Filter(t *testing.T) {
	f := &FilterList{}
	f.Add(&WordFilter{})
	f.Add(&LenFilter{})

	if err := f.Filter("hello felix"); err != nil {
		t.Errorf("filter failed, err: %v", err)
	}
}
