package app

type Watcher interface {
	String() (ch <- chan string)
	Strings() (ch <- chan []string)
	Int() (ch <- chan int)
	Ints() (ch <- chan []int)
}

type Group interface {
	Int(key string) (value int, err error)
	Ints(key string) (value []int, err error)
	String(key string) (value string, err error)
	Strings(key string) (value []string, err error)
	Watcher(key string) (watcher Watcher, err error)
}

type Config interface {
	GetGroup(key string) (group Group, err error)
}
