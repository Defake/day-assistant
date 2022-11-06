package json

type JsonSerializable interface {
    ToJsonString() (string, error)
}

