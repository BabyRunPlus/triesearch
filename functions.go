package triesearch

import (
	"sort"
)

// calcCutNum
/**
 *  @Author: rym 2022-11-16 17:45:36
 *  @Description: 计算字符串需要截取的长度
 *  @param str
 *  @param length
 *  @return strLen
 */
func calcCutNum(str string, length int) (strLen, cutLen int) {
	strLen = len([]rune(str))

	if strLen < length {
		cutLen = strLen
	} else {
		cutLen = length
	}

	return strLen, cutLen
}

// substr
/**
 *  @Author: rym 2022-11-16 17:44:22
 *  @Description: 截取字符串
 *  @param str
 *  @param start
 *  @param length
 *  @return string
 */
func substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	runeStr := []rune(str)
	lenStr := len(runeStr)

	if start < 0 {
		start = lenStr + start
	}
	if start > lenStr {
		start = lenStr
	}
	end := start + length
	if end > lenStr {
		end = lenStr
	}
	if length < 0 {
		end = lenStr + length
	}
	if start > end {
		start, end = end, start
	}
	return string(runeStr[start:end])
}

// listSearchString
/**
 *  @Author: rym 2022-11-15 17:05:12
 *  @Description: string类型列表搜索
 *  @param elem
 *  @param list
 *  @param enableSort
 *  @return out
 */
func listSearchString(elem string, list []string) (out int) {
	sort.Strings(list)

	out = sort.SearchStrings(list, elem)

	if out < len(list) && list[out] == elem {
		return out
	}

	return -1
}
