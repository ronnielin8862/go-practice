package diverseString

import (
	"fmt"
	"math"
	"strings"
)

// 如果需要 admin 內的字串轉大寫的各种排列組合，像 [Admin, ADmin, AdMin,…. ADMIN] 你會如何實作
func diverseString(admin string) (list []string) {
	if len(admin) == 0 {
		return nil
	}

	admin = strings.ToLower(admin)

	for i, _ := range admin {
		if i == 0 {
			list = append(list, admin)
			list = append(list, strings.ToUpper(admin[:1])+admin[1:])
			continue
		}
		var upList []string
		for _, c := range list {
			upList = append(upList, c[0:i]+strings.ToUpper(admin[i:i+1])+c[i+1:])
		}
		list = append(list, upList...)
	}
	return list
}

// Parni version
func generatePermutations(input string) {
	n := len(input)
	total := int(math.Pow(2, float64(n)))

	for i := 0; i < total; i++ {
		result := ""
		for j := 0; j < n; j++ {
			bit := (i >> j) & 1
			char := input[j]

			if bit == 0 {
				result += string(char)
			} else {
				result += strings.ToUpper(string(char))
			}
		}
		fmt.Println(result)
	}
}
