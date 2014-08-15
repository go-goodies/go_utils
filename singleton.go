package go_utils

type Obj map[string]interface{}

type Singleton struct {
	Data Obj
}

var instance *Singleton = nil

// NewSingleton is a singleton object which restricts object creation to only one instance.
// You can use this when you want an application wide in-memory data store, e.g. to store
// application configuration data.
func NewSingleton() *Singleton {
	if instance == nil {
		instance = &Singleton{};
		instance.Data = make(Obj)
	}
	return instance;
}
