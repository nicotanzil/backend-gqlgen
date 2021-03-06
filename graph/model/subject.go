package model

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}
