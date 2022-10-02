package interfaces

//go:generate mockgen --destination=./mocks/gun_state.go interfaces GunState
type GunState interface {
	ReloadGun(sizeChamber, numBullets int)
	SpinChamber()
	ShuffleChamber()
	Shoot() bool
	NumBullets() int
	NumBulletsLeft() int
	ChamberSize() int
	Seed() int64
	SetSeed(seed int64)
	CurrentChamber() int
	SetCurrentChamber(currChamber int)
	Chamber() []int
	SetChamber(chamber []int)
}
