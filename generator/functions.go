// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package generator

import (
	"net/http"
	"sort"
	"strings"
	"text/template"

	"github.com/yunify/snips/capsules"
	"github.com/yunify/snips/utils"
)

var funcMap = template.FuncMap{
	"snakeCase": utils.SnakeCase,
	"camelCase": utils.CamelCase,

	"lower":          lower,
	"lowerFirst":     utils.LowerFirstCharacter,
	"lowerFirstWord": utils.LowerFirstWord,
	"upperFirst":     utils.UpperFirstCharacter,
	"normalized":     normalized,
	"dashConnected":  dashConnected,

	"commaConnected":          commaConnected,
	"commaConnectedWithQuote": commaConnectedWithQuote,

	"replace":     replace,
	"passThrough": passThrough,

	"firstPropertyIDInCustomizedType": firstPropertyIDInCustomizedType,

	"statusText": statusText,

	"contains": contains,
	"indexOf": indexOf,
	"substr": substr,

	"plus": plus,
	"plusOne": plusOne,
	"minus": minus,
	"minusOne": minusOne,

	"hasURIParams": hasURIParams,
	"uriParams": uriParams,
	"formatURI": formatURI,

	"isFirst": isFirst,
}

func lower(original string) string {
	return strings.ToLower(original)
}

func normalized(original string) string {
	return utils.CamelCaseToCamelCase(utils.SnakeCaseToSnakeCase(original))
}

func dashConnected(original string) string {
	return utils.SnakeCaseToDashConnected(utils.SnakeCase(original))
}

func commaConnected(stringArray []string) string {
	return strings.Join(stringArray, ", ")
}

func commaConnectedWithQuote(stringArray []string) string {
	quoteStringArray := []string{}
	for _, value := range stringArray {
		quoteStringArray = append(quoteStringArray, `"`+value+`"`)
	}
	return strings.Join(quoteStringArray, ", ")
}

func replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func passThrough(data ...interface{}) []interface{} {
	return data
}

func firstPropertyIDInCustomizedType(customizedType *capsules.Property) string {
	keys := []string{}
	for key := range customizedType.Properties {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	if len(keys) > 0 {
		return keys[0]
	}

	return ""
}

// statusText translates the integer status code into string text in camelcase.
// For example:
//     200 -> "OK"
//     201 -> "Created"
//     418 -> "Teapot"
func statusText(statusCode int) (statusText string) {
	statusText = http.StatusText(statusCode)

	// Replace special HTTP status description.
	statusText = strings.Replace(statusText, "I'm a teapot", "Teapot", -1)

	// Remove dash and space.
	statusText = strings.Replace(statusText, "-", "", -1)
	statusText = strings.Replace(statusText, " ", "", -1)

	return
}



func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func indexOf(s, sep string) int {
	return strings.Index(s, sep)
}

func substr(s string, start, end int) string {
	rs := []rune(s)
    rl := len(rs)
        
    if start > end {
        start, end = end, start
    }
    
    if start < 0 {
        start = 0
    }

    if start > rl {
        start = rl
    }

    if end < 0 {
        end = 0
    }

    if end > rl {
        end = rl
    }

    return string(rs[start:end])
}

func plus(i, size int) int {
	return i + size
}

func plusOne(i int) int {
	return plus(i, 1)
}

func minus(i, size int) int {
	return i - size
}

func minusOne(i int) int {
	return minus(i, 1)
}

func hasURIParams(uri string) bool {
	return contains(uri, "{") && contains(uri, "}")
}

func uriParams(uri string) []string {
	params := []string{}

	for hasURIParams(uri) {
		startIndex := indexOf(uri, "{") + 1
		endIndex := indexOf(uri, "}")
		param := utils.LowerFirstWord(utils.CamelCase(substr(uri, startIndex, endIndex)))
		params = append(params, param)

		uri = replace(uri, "{", "", 1)
		uri = replace(uri, "}", "", 1)
    }

    return params
}

func formatURI(uri string) string {
	for hasURIParams(uri) {
		startIndex := indexOf(uri, "{") + 1
		endIndex := indexOf(uri, "}")
		param := substr(uri, startIndex, endIndex)
		uri = replace(uri, "{" + param + "}", "\\(input." + utils.LowerFirstWord(utils.CamelCase(param)) + ")", 1)
    } 

    return uri
}

func isFirst(stringArray []string, content string) bool {
	return stringArray[0] == content
}
