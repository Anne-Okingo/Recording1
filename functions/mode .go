package function

import (
	"fmt"
	"strconv"
)

func Mode(n []int) {
	result := ""
	sliced := []string{}
	for i := 0; i < len(n); i++ {
		if n[i] != ',' {
			result += strconv.Itoa((n[i]))
		} else if result != "" {
			sliced = append(sliced, result)
			result = ""
		}
	}
	if result != "" {
		sliced = append(sliced, result)
	}
	fmt.Println(sliced)

	result2 := ""
	for _, ch := range sliced {
		result2 += result2 + string(ch)
	}
	fmt.Println(result2)

	repeat := 1
	result4 := ""
	for i := 1; i < len(result2); i++ {
		if result2[i] == result2[i-1] {
			repeat++
		} else {
			repeat = 1
		}
		// fmt.Println(result2[i])
		// fmt.Println(repeat)
		result4 += string(repeat)
		fmt.Println(repeat)
		// }
	}

	// for i := 0; i <len(result4); i++{
	// 	for j := 0; j < len(result4)-i-1; j++{
	// 		if result4[j+1] > result4[j]{
	// 			result4[j] =result4[j+1]
	// 		}
	// 	}
	// }
}
