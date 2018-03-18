package algs

// MergeSort stable sort algorithm
// https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F
func MergeSort(outOrderSeq []int32) []int32 {
	if len(outOrderSeq) < 2 {
		return outOrderSeq
	}

	mid := len(outOrderSeq) / 2
	leftSeq := MergeSort(outOrderSeq[0:mid])
	rightSeq := MergeSort(outOrderSeq[mid:])

	return merge(leftSeq, rightSeq)
}

// merge two sorted array
func merge(seq1, seq2 []int32) []int32 {
	i, j := 0, 0
	len1, len2 := len(seq1), len(seq2)

	var seq []int32

	for i < len1 && j < len2 {
		if seq1[i] <= seq2[j] {
			seq = append(seq, seq1[i])
			i++
		} else {
			seq = append(seq, seq2[j])
			j++
		}
	}

	for i < len1 {
		seq = append(seq, seq1[i])
		i++
	}

	for j < len2 {
		seq = append(seq, seq2[j])
		j++
	}

	return seq
}
