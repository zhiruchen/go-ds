package basic

import "testing"

type tVal int32

func (v tVal) Compare(v2 Comparator) int {
	vv := v2.(tVal)
	if v < vv {
		return -1
	} else if v == vv {
		return 0
	} else {
		return 1
	}
}

func genTestValList(vs ...int32) []tVal {
	var tvs []tVal
	for _, v := range vs {
		tvs = append(tvs, tVal(v))
	}

	return tvs
}

func TestLinkedListV2_MergeSort(t *testing.T) {
	cases := []struct {
		input  []tVal
		output string
	}{
		{
			input:  genTestValList(1),
			output: "1",
		},
		{
			input:  genTestValList(1234, 5678, 0, -1, 9, 2, 10),
			output: "-1->0->2->9->10->1234->5678",
		},
		{
			input:  genTestValList(5, 4, 3, 2, 1),
			output: "1->2->3->4->5",
		},
		{
			input:  genTestValList(6, 7, 8, 9, 1000),
			output: "6->7->8->9->1000",
		},
		{
			input:  genTestValList(6, 7, 8, 9, 1000, 0),
			output: "0->6->7->8->9->1000",
		},
	}

	for _, cc := range cases {
		ll := NewLinkedListV2()
		for _, v := range cc.input {
			ll.AppendToTail(v)
		}

		ll.MergeSort()
		strRep := ll.String()
		t.Logf("str rep:%s\n", strRep)
		if cc.output != strRep {
			t.Errorf("expected: %s, get: %s\n", cc.output, strRep)
		}
	}
}

func TestMergeTwoSortedLinkedList(t *testing.T) {
	cases := []struct {
		l1     []tVal
		l2     []tVal
		output string
	}{
		{
			l1:     genTestValList(1, 2, 3),
			l2:     genTestValList(2, 3, 4),
			output: "1->2->2->3->3->4",
		},
		{
			l1:     genTestValList(1234, 5678),
			l2:     genTestValList(0, 2, 9, 10),
			output: "0->2->9->10->1234->5678",
		},
		{
			l1:     genTestValList(10000),
			l2:     genTestValList(100000000),
			output: "10000->100000000",
		},
	}

	for _, cc := range cases {
		ll1 := NewLinkedListV2()
		for _, v := range cc.l1 {
			ll1.AppendToTail(v)
		}

		ll1.MergeSort()

		ll2 := NewLinkedListV2()
		for _, v := range cc.l2 {
			ll2.AppendToTail(v)
		}

		ll2.MergeSort()

		sortedLL := MergeTwoSortedLinkedList(ll1, ll2)
		strRep := sortedLL.String()
		t.Logf("str rep:%s\n", strRep)

		if cc.output != strRep {
			t.Errorf("expected: %s, get: %s\n", cc.output, strRep)
		}
	}
}
