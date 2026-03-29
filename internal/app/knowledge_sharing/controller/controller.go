package controller

import "strconv"

func GetPage(pageStr string) (page int) {

	page, _ = strconv.Atoi(pageStr)

	if page <= 0 {
		page = 1
	}

	return
}
