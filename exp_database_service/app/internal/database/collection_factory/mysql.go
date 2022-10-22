package collectionfactory

import (
	conn "app/internal/database/connection"
	"app/internal/database/entities"
	"context"
)

type MySql struct {
	ctx         context.Context
	Collections map[entityName]entities.Collection
	Conn        *conn.MySQLConnection
}

func NewMySQLHandler(ctx context.Context, MysqlUri string, DbName string) (CollectionFactory, error) {
	mysql := &MySql{
		ctx:         ctx,
		Collections: make(map[entityName]entities.Collection),
		Conn:        conn.NewMySQLConnection(MysqlUri, DbName, true)}
	err := mysql.Conn.Connect()
	if err != nil {
		return nil, err
	}
	mysql.Collections[EnterpriseInfo] = entities.NewProductCollection(mysql.Conn.DB)
	var c CollectionFactory = mysql
	return c, nil
}

func (m *MySql) Initialize() error {
	for _, collection := range m.Collections {
		if err := collection.InitCollection(m.ctx); err != nil {
			return err
		}
	}
	return nil
}

func (m *MySql) GetCollection(name string) entities.Collection {
	return m.Collections[entityName(name)]
}
