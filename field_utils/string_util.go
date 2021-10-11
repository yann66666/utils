// Author: yann
// Date: 2020/5/25 6:26 下午
// Desc: field_utils

package field_utils

import "strings"

//是否为空字符串
func IsEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}
