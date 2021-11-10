package cardversions

type CardVersion struct {
	Name string `json:"version_name"`
	ID   string `json:"get"`
}

type cardVersionsData struct {
	CardVersions []CardVersion `json:"data"`
}
