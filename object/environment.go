package object

// NewEnclosedEnvironment extending the environment, create a new instance of object.Environment with
// a pointer to the environment it should extend
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment create new env
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment env
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get stored value for given name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set set value for given name
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
