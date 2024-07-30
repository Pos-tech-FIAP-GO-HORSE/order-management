package migrations

type IMigration interface {
	Up() error
	Down() error
}
