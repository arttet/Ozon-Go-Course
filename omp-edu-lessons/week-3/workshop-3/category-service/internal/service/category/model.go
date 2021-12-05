package category

type Category struct {
	ID   uint64
	Name string
}

type Categories []*Category

func (c Categories) FilterByID(id uint64) *Category {
	for _, cat := range c {
		if cat.ID == id {
			return cat
		}
	}
	return nil
}
