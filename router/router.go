package router

import (
	"github.com/denizakturk/servfront/bridge"
	"regexp"
	"strings"
)
func (r *RouteAddress) PrepareAddress() {
	var regexpString = ""
	if nil == r.Params {
		r.PrepareParams()
	}
	if nil != r.Params {
		regexpString = r.Pattern
		for _, val := range *r.Params {
			oldStr := strings.Join([]string{"{", val.ParamName, "}"}, "")
			var rString = ""
			if val.ParamType == "number" {
				rString = "[0-9]+"
			} else if val.ParamType == "string" {
				rString = "[a-z|A-Z|0-9-_]+"
			}
			newStr := strings.Join([]string{"(?m)(?P<", val.ParamName, ">", rString, ")"}, "")
			regexpString = strings.Replace(regexpString, oldStr, newStr, 1)
		}
	}

	r.RegexpPattern = regexp.MustCompile(regexpString)
}

func (r *RouteAddress) PrepareParams() {
	mathParams := regexp.MustCompile("({[a-z|A-Z-_]+})")
	matches := mathParams.FindAll([]byte(r.Pattern), -1)
	matchParams := []RouteAddressParameter{}
	for _, val := range matches {
		str := string(val)
		matchParams = append(matchParams, RouteAddressParameter{ParamType: "string", ParamName: strings.Trim(str, "{}")})
	}
	r.Params = &matchParams
}

func (r *RouteAddress) CatchAddressParametersValue(url string) {
	dest := []byte{}
	if nil == r.Params {
		r.PrepareParams()
	}
	if nil != r.Params {
		var paramTemplate = []string{}
		for _, val := range *r.Params {
			paramTemplate = append(paramTemplate, strings.Join([]string{"$", val.ParamName}, ""))
		}

		for _, submatches := range r.RegexpPattern.FindAllSubmatchIndex([]byte(url), -1) {
			dest = r.RegexpPattern.Expand(dest, []byte(strings.Join(paramTemplate, ";")), []byte(url), submatches)
		}
		parameters := strings.Split(string(dest), ";")
		for key, _ := range *r.Params {
			(*r.Params)[key].Value = &parameters[key]
		}
	}

}

func (r *RouteAddress) ParamsToMap() (params map[string]string) {
	params = make(map[string]string)
	for _, val := range *r.Params {
		params[val.ParamName] = *val.Value
	}

	return params
}

type Route struct {
	Name       string
	Pattern    *regexp.Regexp
	Address    *RouteAddress
	Controller bridge.Struct
	Endpoint   func() *bridge.ServiceResponse
}

type RouteAddressParameter struct {
	ParamName string  `json:"param_name"`
	ParamType string  `json:"param_type"`
	Value     *string `json:"value"`
}

type RouteAddress struct {
	Pattern       string `json:"pattern"`
	RegexpPattern *regexp.Regexp
	Params        *[]RouteAddressParameter
}
