package dynamic_programming

// LengthOfLongestNoDescSeq 序列seq中最长的非递减序列的长度
func LengthOfLongestNoDescSeq(seq []int32) int32 {
	seqLen := len(seq)
	if seqLen <= 1 {
		return 1
	}

	var lenStates = make([]int32, seqLen+1)
	lenStates[1] = 1

	for i := 1; i < seqLen; i++ {
		if seq[i] >= seq[i-1] && lenStates[i]+1 > lenStates[i+1] {
			lenStates[i+1] = lenStates[i] + 1
		} else {
			lenStates[i+1] = lenStates[i]
		}
	}

	return lenStates[seqLen]
}
