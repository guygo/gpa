package repoistorydb

type Repoistory struct {
	tabels map[string]*table
}

func (r *Repoistory) Init() {
	r.tabels = make(map[string]*table)
}

func (r *Repoistory) CreateTable(tc TableContext) {
	t := createTable(tc.newTableObject())
	r.tabels[t.tableName] = &t

}
func (r *Repoistory) CreateMultipleTables(tabels []TableContext) {
	for _, v := range tabels {
		r.CreateTable(v)
	}
}

type TableContext interface {
	newTableObject() interface{}
}
