package subject

import "github.com/yoshinorihisakawa/design-pattern/observer/observer"

type Subject interface {
	RegisterObserver(observer observer.Observer)
	RemoveObserver(observer observer.Observer)
	NotifyObserver()
	DetectError(str string)
}