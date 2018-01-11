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
	"regexp"
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
	"formatSwiftPath": formatSwiftPath,
	"isFirst": isFirst,
	"joinResponses": joinResponses,
	"isUploadOperation": isUploadOperation,

	"isLetterFirstOfWord": isLetterFirstOfWord,
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

func hasPathParams(path string) bool {
	return contains(path, "{") && contains(path, "}")
}

func pathParams(path string) []string {
	params := []string{}

	for hasPathParams(path) {
		startIndex := indexOf(path, "{") + 1
		endIndex := indexOf(path, "}")
		params = append(params, substr(path, startIndex, endIndex))

		path = replace(path, "{", "", 1)
		path = replace(path, "}", "", 1)
    }

    return params
}

func formatSwiftPath(operation capsules.Operation) string {
	var path = operation.Request.Path
	var params = pathParams(path)
	
	for _, property := range operation.Request.Properties.Properties {
		for _, param := range params {
			if contains(param, property.ID) {
				path = replace(path, "{" + param + "}", "\\(input." + utils.LowerFirstWord(utils.CamelCase(property.ID)) + "!)", 1)
				break
			}
		}
	}
	
	return path
}

func isFirst(stringArray []string, content string) bool {
	return stringArray[0] == content
}

func joinResponses(operation capsules.Operation) *capsules.Response {
	var response *capsules.Response

	for _, value := range operation.Responses {
		if response == nil {
			response = value
		} else {	
			appenProperties(response.Headers, value.Headers)
			appenProperties(response.Elements, value.Elements)
			appenProperties(response.Body, value.Body)
		}
	}

	return response
}

func appenProperties(proeprty1 *capsules.Property, property2 *capsules.Property) {
	for key, value := range property2.Properties {
		proeprty1.Properties[key] = value
	}
}

func isUploadOperation(operation capsules.Operation) bool {
	if operation.Request.Body.Type == "binary" {
		return true
	}

	for _, property := range operation.Request.FormData.Properties {
		if property.Type == "file" {
			return true
		}
	}

	return false
}

func isLetterFirstOfWord(word string) bool {
	return regexp.MustCompile("^[^A-Za-z]").MatchString(word)
}
