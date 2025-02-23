package iteration

func Repeat(value string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += value
	}

	return repeated
}
