package repoistorydb

import (
	"reflect"
	"strings"
)

type relationkind int

const (
	NONE relationkind = iota
	ONETOONE
	ONETOMANY
	MANYTOMANY
)

type relation struct {
	relationkind
	relatedto string
}

type tablestruct interface{}
type fieldMeta struct {
	name       string
	constraint string
	sqltype    string
	value      interface{}
}

type table struct {
	tableName string
	uniqueIds *[]string
	fields    []fieldMeta
	relations []relation
}

func (t *table) New() {
	t.fields = make([]fieldMeta, 0)
	t.relations = make([]relation, 0)
}

func (t *table) addField(protoField reflect.StructField) {
	uniqueIds := make([]string, 0)

	if constraint, ok := protoField.Tag.Lookup("constraint"); ok {
		switch c := strings.ToUpper(constraint); {
		case c == "":
			t.fields = append(t.fields, fieldMeta{
				sqltype:    SqlTypeNames[protoField.Type.Kind()],
				constraint: c,
				name:       protoField.Name,
			})

		case strings.Contains(c, "PRIMARY"):
			t.fields = append(t.fields, fieldMeta{
				sqltype:    SqlTypeNames[protoField.Type.Kind()],
				constraint: c,
				name:       protoField.Name,
			})
			uniqueIds = append(uniqueIds, protoField.Name)
		}
	} else {
		t.fields = append(t.fields, fieldMeta{
			sqltype:    SqlTypeNames[protoField.Type.Kind()],
			constraint: "",
			name:       protoField.Name,
		})

	}
	if len(uniqueIds) == 0 {
		uniqueIds = append(uniqueIds, "id")
	}
	t.uniqueIds = &uniqueIds
}

func createTable(tableprotype tablestruct) table {
	var tablePrototype reflect.Type = reflect.TypeOf(tableprotype)

	var newtable table
	newtable.New()

	newtable.tableName = tablePrototype.Name()
	for i := 0; i < tablePrototype.NumField(); i++ {
		protoField := tablePrototype.Field(i)
		switch kind := protoField.Type.Kind(); kind {
		case reflect.Struct:
			newtable.relations = append(newtable.relations, relation{ONETOONE, protoField.Type.Name()})
		case reflect.Slice:
			element := protoField.Type.Elem()
			if element.Kind() == reflect.Struct {
				newtable.relations = append(newtable.relations, relation{MANYTOMANY, element.Name()})
			}
		default:

			newtable.addField(protoField)
		}

	}

	return newtable //fmt.Sprintf("CREATE TABLE %s (\n%s\n)", createdTableName, KeyValueString(fmap))

}
