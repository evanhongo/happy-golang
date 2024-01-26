package schema

type ISchema interface {
	Parse(data any) error
}
