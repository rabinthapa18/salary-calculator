package pkg

func CheckPairQuantity(n int, parts []string) bool {
	return len(parts) == 4+n*2
}
