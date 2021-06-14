package rules

type element struct {
	content Content
	code    string
}

func createElement(
	content Content,
	code string,
) Element {
	out := element{
		content: content,
		code:    code,
	}

	return &out
}

// Content returns the content
func (obj *element) Content() Content {
	return obj.content
}

// Code returns the code
func (obj *element) Code() string {
	return obj.code
}
