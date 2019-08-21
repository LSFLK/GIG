package requests

const (
	PropTypeContent    string = "&prop=extracts&explaintext"
	PropTypeLinks      string = "&prop=links"
	PropTypeCategories string = "&prop=categories"
)

func PropTypes() []string {
	typeList := []string{
		PropTypeLinks,
		PropTypeCategories,
		PropTypeContent,
	}
	return typeList
}
