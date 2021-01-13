package core

type Validatable interface {
	Validate() error
}

type Editable interface {
	Update(obj interface{}) (Object, error)
}

type Keyable interface {
	GenerateAndSetKey() (Object, error)
	GetKey() string
	SetKey(key string) (Object, error)
}

type Object interface {
	Validatable
	Editable
	Keyable
	UnmarshalFromMap(objMap map[string]interface{}) (Object, error)
	New() Object
	IsValidUID() bool
}
