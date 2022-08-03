package mongo

type DbConfig struct {
	DatabaseName   string
	CollectionName string
}

func NewDbConfig(dbName string, collectionName string) DbConfig {
	return DbConfig{DatabaseName: dbName, CollectionName: collectionName}
}
