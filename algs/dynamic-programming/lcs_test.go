package dynamic_programming

import (
	"strings"
	"testing"
)

func TestLongestCommonSubSequence(t *testing.T) {
	testCases := []struct {
		seq1   []string
		seq2   []string
		length int32
	}{
		{
			seq1:   strings.Split("A B C", " "),
			seq2:   strings.Split("B C D", " "),
			length: 2,
		},
		{
			seq1:   strings.Split("A B C D G H", " "),
			seq2:   strings.Split("A E D F H R", " "),
			length: 3,
		},
		{
			seq1:   strings.Split("A G G T A B", " "),
			seq2:   strings.Split("G X T X A Y B", " "),
			length: 4,
		},
	}

	for _, cc := range testCases {
		length := LongestCommonSubSequence(cc.seq1, cc.seq2)
		if length != cc.length {
			t.Errorf("expect length: %d, get: %d\n", cc.length, length)
		}
	}
}

func TestLongestCommonSubSequenceNonRecur(t *testing.T) {
	testCases := []struct {
		seq1   []string
		seq2   []string
		length int32
	}{
		{
			seq1:   strings.Split("A B C", " "),
			seq2:   strings.Split("B C D", " "),
			length: 2,
		},
		{
			seq1:   strings.Split("A B C D G H", " "),
			seq2:   strings.Split("A E D F H R", " "),
			length: 3,
		},
		{
			seq1:   strings.Split("A G G T A B", " "),
			seq2:   strings.Split("G X T X A Y B", " "),
			length: 4,
		},
	}

	for _, cc := range testCases {
		length := LongestCommonSubSequenceNonRecur(cc.seq1, cc.seq2)
		if length != cc.length {
			t.Errorf("expect length: %d, get: %d\n", cc.length, length)
		}
	}
}
