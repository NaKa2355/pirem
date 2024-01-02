package boundary

type Boundary interface {
	ButtonGetter
	ButtonPusher

	RemoteCreator
	RemoteGetter
	RemoteLister
	RemoteUpdater
	RemoteDeleter
	IRLearner

	DeviceGetter
	DevicesLister
	IRSender
	IRReceiver
	IRGetter
}
