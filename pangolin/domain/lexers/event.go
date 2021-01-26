package lexers

type event struct {
	token string
	fn    EventFn
}

func createEvent(token string, fn EventFn) Event {
	out := event{
		token: token,
		fn:    fn,
	}

	return &out
}

// Token returns the token
func (obj *event) Token() string {
	return obj.token
}

// Fn returns the event fn
func (obj *event) Fn() EventFn {
	return obj.fn
}
