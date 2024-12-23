package scalar

type OptionsProp struct {
	CDN       string `json:"cdn,omitempty"`
	SpecURL   string `json:"specUrl,omitempty"` // allow external URL ou local path file
	CustomCss string `json:"customCss,omitempty"`
	CustomOptions
}

type CustomOptions struct {
	PageTitle string `json:"pageTitle,omitempty"`
}
