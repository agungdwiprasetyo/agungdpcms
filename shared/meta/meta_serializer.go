package meta

type MetaSchema struct {
	Meta *Meta
}

func (m *MetaSchema) Page() int32 {
	return int32(m.Meta.Page)
}

func (m *MetaSchema) Limit() int32 {
	return int32(m.Meta.Limit)
}

func (m *MetaSchema) TotalRecords() int32 {
	return int32(m.Meta.TotalRecords)
}

func (m *MetaSchema) TotalPages() int32 {
	return int32(m.Meta.TotalPages)
}
