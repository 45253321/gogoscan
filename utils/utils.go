package utils

// 致敬python的itertools.product
func ItertoolsProduct(first ,second []string) *[][2]string{

	result := make([][2]string, 0, 0)
	for _, f := range first{
		for _, s := range second{
			result = append(result, [2]string{f, s})
		}
	}
	return &result
}

