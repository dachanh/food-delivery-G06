package memcache

type Caching interface {
	Write(k string, value interface{})
	Read(k string) interface{}
	WriteTTL(k string, value interface{}, exp int)
}
