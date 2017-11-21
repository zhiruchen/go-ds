package advanced

/*
	RecurReverseString 递归的反转一个字符串列表
	example:
	input: s: ["1","2","3"]
	递归的  output = RecurReverseString(["2","3"])+["1"]
           output = RecurReverseString(["3"]) + ["2"] + ["1"]
		   output = ["3"] + ["2"] + ["1"]
 */
func RecurReverseString(input []string) []string {
	if len(input) == 1 {
		return input
	}

	return append(RecurReverseString(input[1:]), input[0])
}
