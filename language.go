package carbon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	// defaultDir default directory
	defaultDir = "lang"

	// defaultLocale default locale
	defaultLocale = "en"

	// invalidLocaleError returns an invalid locale error.
	invalidLocaleError = func(locale string) error {
		return fmt.Errorf("invalid locale file %q, please make sure the json file exists and is valid", locale)
	}
)

// Language define a Language struct.
// 定义 Language 结构体
type Language struct {
	dir       string
	locale    string
	resources map[string]string
	Error     error
}

// NewLanguage return a new Language instance.
// 初始化 Language 结构体
func NewLanguage() *Language {
	return &Language{
		dir:       defaultDir,
		locale:    defaultLocale,
		resources: make(map[string]string),
	}
}

// SetLocale sets language locale.
// 设置区域
func (lang *Language) SetLocale(locale string) {
	if len(lang.resources) != 0 {
		return
	}
	lang.locale = locale
	fileName := lang.dir + string(os.PathSeparator) + locale + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		lang.Error = invalidLocaleError(fileName)
	}
	if json.Unmarshal(bytes, &lang.resources) != nil {
		lang.Error = invalidLocaleError(fileName)
	}
}

// SetResources sets language resources.
// 设置资源
func (l *Language) SetResources(resources map[string]string) {
	if len(l.resources) == 0 {
		l.resources = resources
		return
	}
	for k, v := range resources {
		if _, ok := l.resources[k]; ok {
			l.resources[k] = v
		}
	}
}

// translate translates by unit string.
// 翻译转换
func (lang *Language) translate(unit string, vaule int64) string {
	if len(lang.resources) == 0 {
		lang.SetLocale(defaultLocale)
	}
	slice := strings.Split(lang.resources[unit], "|")
	number := getAbsValue(vaule)
	if len(slice) == 1 {
		return strings.Replace(slice[0], "%d", strconv.FormatInt(vaule, 10), 1)
	}
	if int64(len(slice)) <= number {
		return strings.Replace(slice[len(slice)-1], "%d", strconv.FormatInt(vaule, 10), 1)
	}
	if !strings.Contains(slice[number-1], "%d") && vaule < 0 {
		return "-" + slice[number-1]
	}
	return strings.Replace(slice[number-1], "%d", strconv.FormatInt(vaule, 10), 1)
}
