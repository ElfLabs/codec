package codec

type codecError string

func (e codecError) Error() string {
	return string(e)
}
