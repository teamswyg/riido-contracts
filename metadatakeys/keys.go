package metadatakeys

type Key string

func (key Key) String() string {
	return string(key)
}
