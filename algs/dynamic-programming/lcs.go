package dynamic_programming

// LongestCommonSubSequence 最长公共子序列的长度
func LongestCommonSubSequence(seq1, seq2 []string) int32 {
	return lcs(seq1, seq2, len(seq1), len(seq2))
}

func lcs(seq1, seq2 []string, m, n int) int32 {
	if m == 0 || n == 0 {
		return 0
	} else if seq1[m-1] == seq2[n-1] {
		return 1 + lcs(seq1, seq2, m-1, n-1)
	} else {
		l1 := lcs(seq1, seq2, m-1, n)
		l2 := lcs(seq1, seq2, m, n-1)
		return max(l1, l2)
	}
}

func LongestCommonSubSequenceNonRecur(seq1, seq2 []string) int32 {
	seq1Len := len(seq1)
	seq2Len := len(seq2)
	var l = make([][]int32, seq1Len+1)
	for i := range l {
		l[i] = make([]int32, seq2Len+1)
	}

	for i := 0; i <= seq1Len; i++ {
		for j := 0; j <= seq2Len; j++ {
			if i == 0 || j == 0 {
				l[i][j] = 0
			} else if seq1[i-1] == seq2[j-1] {
				l[i][j] = l[i-1][j-1] + 1
			} else {
				l[i][j] = max(l[i-1][j], l[i][j-1])
			}
		}
	}

	return l[seq1Len][seq2Len]
}
