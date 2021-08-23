package model

type Jenkins struct {
	URL     	string `json:"url"`
	Result 		string `json:"result"`
	ID   		string `json:"id"`
	Description	string `json:"description"`
	BuiltOn		string `json:"built_on"`

	Actions []struct {
		Causes []struct {
			UserName string `json:"userName"`
		} `json:"causes,omitempty"`
	} `json:"actions"`
}