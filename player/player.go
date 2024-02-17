package player

type IPlayer interface {
	DisplayInfo()
	DeleteHealth(damage int)
}
