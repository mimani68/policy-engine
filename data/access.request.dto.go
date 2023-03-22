package data

type AccessRequestDto struct {
	Realm        string `json:"realm"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
}
