package domain

type RemoteTag = string

func NewRemoteTag(tag string) RemoteTag {
	return RemoteTag(tag)
}
