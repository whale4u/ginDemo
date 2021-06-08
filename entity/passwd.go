package entity

type Passwd struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Note     string `json:"note"`
}
