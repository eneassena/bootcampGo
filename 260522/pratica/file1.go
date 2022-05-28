package pratica

/*
type errorPersonalizado interface {
	Error() (string, error)
}

type MyError struct {
	Message string
}

func (m MyError) Error() (string, error) {
	if len(m.Message) == 0 {
		return "", errors.New(fmt.Errorf("%s", "error: Campo esta vasio").Error())
		//return "", fmt.Errorf("%s", "campo esta vasio")
	}
	return m.getMessage(), nil
}

func (m *MyError) getMessage() string {
	return m.Message
}

func ShowData(m errorPersonalizado) {
	fmt.Println(m.Error())
}

*/
