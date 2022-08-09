package fake

type RepositoryMock[Entity any] struct {
	ReturnsCreate  bool
	ReturnsGetBy   *Entity
	ReturnsGetById *Entity
	ReturnsList    *[]Entity
}

func (r *RepositoryMock[Entity]) Create(obj *Entity) bool {
	return r.ReturnsCreate
}

func (r *RepositoryMock[Entity]) GetBy(where Entity) *Entity {
	return r.ReturnsGetBy
}

func (r *RepositoryMock[Entity]) GetById(id string) *Entity {
	return r.ReturnsGetById
}

func (r *RepositoryMock[Entity]) List(where Entity) *[]Entity {
	return r.ReturnsList
}
