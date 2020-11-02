package hnjobs

type Store interface {
	UpsertWhoIsHiringPost(ID int) error
	UpsertJob(ID int) error
}
