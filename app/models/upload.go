package models

type Upload struct {
	Source string `json:"source" bson:"source"`
	Title     string `json:"title" bson:"title"`
}

func (u Upload) SetSource(value string) Upload {
	u.Source = value
	return u
}

func (u Upload) GetSource() string {
	return u.Source
}

func (u Upload) SetTitle(value string) Upload {
	u.Title = value
	return u
}

func (u Upload) GetTitle() string {
	return u.Title
}
