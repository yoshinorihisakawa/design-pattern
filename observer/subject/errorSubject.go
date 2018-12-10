package subject

import "github.com/yoshinorihisakawa/design-pattern/observer/observer"

type ErrorSubject struct {
	observers []observer.Observer
	str string
}

func (e *ErrorSubject) RegisterObserver(observer observer.Observer) {
	e.observers = append(e.observers, observer)
}

func (e *ErrorSubject) RemoveObserver(observer observer.Observer) {
	result := ErrorSubject{}
	for _, o := range e.observers {
		if o != observer {
			result.observers = append(result.observers, o)
		}
	}
	e.observers = result.observers
}

func (e *ErrorSubject) NotifyObserver() {
	for _, o := range e.observers {
		o.Update(e.str)
	}
}

func (e *ErrorSubject) DetectError(str string) {
	e.str = str
	e.NotifyObserver()
}


