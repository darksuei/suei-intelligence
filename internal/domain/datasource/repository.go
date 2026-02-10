package datasource

type DatasourceRepository interface {
	Find(projectId uint) (*[]Datasource, error)
	Create(payload *Datasource) (*Datasource, error)
	Update(payload *Datasource) error
	Delete(datasourceId uint, projectId uint) error
}