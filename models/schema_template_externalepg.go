package models

type TemplateExternalepg struct {
	Ops   string                 `json:",omitempty"`
	Path  string                 `json:",omitempty"`
	Value map[string]interface{} `json:",omitempty"`
}

func NewTemplateExternalepg(ops, path, name, displayName string, vrfRef map[string]interface{}, l3outRef map[string]interface{}) *TemplateExternalepg {
	var externalepgMap map[string]interface{}
	externalepgMap = map[string]interface{}{
		"name":        name,
		"displayName": displayName,
		"vrfRef":      vrfRef,
	}

	if l3outRef != nil {
		externalepgMap["l3outRef"] = l3outRef
	}

	return &TemplateExternalepg{
		Ops:   ops,
		Path:  path,
		Value: externalepgMap,
	}

}

func (externalepgAttributes *TemplateExternalepg) ToMap() (map[string]interface{}, error) {
	externalepgAttributesMap := make(map[string]interface{})
	A(externalepgAttributesMap, "op", externalepgAttributes.Ops)
	A(externalepgAttributesMap, "path", externalepgAttributes.Path)
	if externalepgAttributes.Value != nil {
		A(externalepgAttributesMap, "value", externalepgAttributes.Value)
	}

	return externalepgAttributesMap, nil
}
