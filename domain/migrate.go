package domain

type MigrateType string

const (
	Schema MigrateType = "schema"
	Data   MigrateType = "data"
)

func IsNotMigrateType(migrate MigrateType) bool {
	switch migrate {
	case Schema, Data:
		return false
	default:
		return true
	}
}

type Migrate struct {
	ForceVersion int `json:"forceVersion"`
}
