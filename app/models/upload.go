package models

type Upload struct {
	sourceURL string
	title     string
}

func (u Upload) SetSource(value string) Upload {
	u.sourceURL = value
	return u
}

func (u Upload) GetSource() string {
	return u.sourceURL
}

func (u Upload) SetTitle(value string) Upload {
	u.title = value
	return u
}

func (u Upload) GetTitle() string {
	return u.title
}
