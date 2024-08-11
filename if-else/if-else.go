package if_else

func If_Else(x int) string {
	if x > 5 {
		return "X is greater than 5"
	} else if x > 2 && x < 5 {
		return "X is less than 5 but greater than 2"
	} else {
		return "X is less than 5"
	}
}
