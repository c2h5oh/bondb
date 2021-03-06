package bondb

var DefaultSession *Session

func mustDefaultSession() *Session {
	if DefaultSession == nil {
		panic("DefaultSession has not been set.")
	}
	return DefaultSession
}

func Query(dst interface{}) *query {
	return mustDefaultSession().Query(dst)
}

func Q(dst interface{}) *query {
	return Query(dst)
}

func Create(item interface{}) (interface{}, error) {
	return mustDefaultSession().Create(item)
}

func Save(item interface{}) error {
	return mustDefaultSession().Save(item)
}

func Delete(item interface{}) error {
	return mustDefaultSession().Delete(item)
}
