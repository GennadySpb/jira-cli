package jira

const (
	customFieldFormatOption  = "option"
	customFieldFormatArray   = "array"
	customFieldFormatNumber  = "number"
	customFieldFormatProject = "project"
	customFieldFormatUser    = "name"
)

type customField map[string]interface{}

type customFieldTypeNumber float64

type customFieldTypeNumberSet struct {
	Set customFieldTypeNumber `json:"set"`
}

type customFieldTypeStringSet struct {
	Set string `json:"set"`
}

type customFieldTypeOption struct {
	Value string `json:"value"`
}

type customFieldTypeName struct {
	Name string `json:"name"`
}

type customFieldTypeOptionSet struct {
	Set customFieldTypeOption `json:"set"`
}

type customFieldTypeOptionAddRemove struct {
	Add    *customFieldTypeOption `json:"add,omitempty"`
	Remove *customFieldTypeOption `json:"remove,omitempty"`
}

type customFieldTypeNameAddRemove struct {
	Add    *customFieldTypeName `json:"add,omitempty"`
	Remove *customFieldTypeName `json:"remove,omitempty"`
}

type customFieldTypeProject struct {
	Value string `json:"key"`
}

type customFieldTypeProjectSet struct {
	Set customFieldTypeProject `json:"set"`
}
