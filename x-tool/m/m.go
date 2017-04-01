package m

type M struct {
}

var MLst map[string]M

func (m *M) reject(str string, model M) {
	MLst[str] = model
}

func (m *M) get(str string) (model M) {
	return MLst[str]
}

func init() {
	MLst = make(map[string]M)
}
