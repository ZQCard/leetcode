package implement_strstr

import (
	"strings"
)

func implementStrstr(haystack string, needle string) int {
	haystackLength := len(haystack)
	needleLength := len(needle)

	if needleLength == 0{
		return 0
	}

	if haystackLength == 0 || haystackLength < needleLength{
		return -1
	}
	for i := 0; i < haystackLength; i++ {
		if haystack[i] == needle[0]{
			if i+needleLength > haystackLength {
				return -1
			}

			temp := haystack[i: i+needleLength]
			if temp == needle{
				return i
			}
		}
	}
	return -1
}

func implementStrstrStardLib(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}