package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res, _ := versionCompare("1.8.9-v.45", "1.8.9-v.91")

	if res == 1 {
		fmt.Println("v1 is greater than v2")
	} else if res == -1 {
		fmt.Println("v2 is greater than v1")
	} else {
		fmt.Println("v1 is equal to v2")
	}

}

func versionCompare(v1, v2 string) (int, error) {
	v1List := strings.Split(v1, ".")
	v2List := strings.Split(v2, ".")
	for index := 0; index < len(v1List); index++ {
		if len(v2List) == index {
			return 1, nil
		}
		v1SubList := strings.Split(v1List[index], "-")
		v2SubList := strings.Split(v2List[index], "-")
		version1, err := strconv.ParseInt(v1SubList[0], 0, 64)
		if err != nil {
			return 0, err
		}
		version2, err := strconv.ParseInt(v2SubList[0], 0, 64)
		if err != nil {
			return 0, err
		}
		if version1 < version2 {
			return -1, nil
		}
		if version1 > version2 {
			return 1, nil
		}
	}
	if len(v1List) < len(v2List) {
		return -1, nil
	}
	return 0, nil
}
