package repoistorydb

import (
	"fmt"
	"reflect"
	"strings"
)

func KeyValueString(m []fieldMeta) string {
	keysValues := make([]string, 0, len(m))
	for i := range m {
		p := &m[i]
		keysValues = append(keysValues, fmt.Sprintf("%s %s %s", p.name, p.sqltype, p.constraint))
	}
	return strings.Join(keysValues, ",")
}
func generateSqlCreateTables(r *Repoistory) []string {

	CreateCommands := make([]string, len(r.tabels))
	index := 0
	for k, v := range r.tabels {
		CreateCommands[index] = fmt.Sprintf("CREATE TABLE %s (%s)", k, KeyValueString(v.fields))
		index++
	}
	return CreateCommands
}
func createOneToOneRealtions(tableName string, id string, ref string) string {
	return fmt.Sprintf("ALTER TABLE %s ADD FOREIGN KEY (%s) REFERENCES (%s) (%s)\tDEFERRABLE INITIALLY DEFERRED", tableName, id, ref, id)
}

func createOneToManyRealtions(tableName string, id string, ref string) string {
	return fmt.Sprintf("ALTER TABLE %s ADD FOREIGN KEY (%s) REFERENCES (%s) (%s)", tableName, id, ref, id)
}

func generateSqlCreateRelations(r *Repoistory) []string {

	CreateAlters := make([]string, 0)

	for k, v := range r.tabels {
		for _, rel := range v.relations {
			x := (v.uniqueIds)
			switch rel.relationkind {
			case ONETOMANY:
				CreateAlters = append(CreateAlters, createOneToOneRealtions(k, strings.Join(*(x), ","), rel.relatedto))
			case MANYTOMANY:
				CreateAlters = append(CreateAlters, createOneToManyRealtions(rel.relatedto, strings.Join(*(x), ","), k))
			}
		}
	}
	return CreateAlters
}

func checkIfStructEmpty(ty reflect.Type, object reflect.Value) bool {
	v := reflect.New(ty).Elem().Interface()
	return v == object
}
func (t *table) InsertCommand(object interface{}) string {
	values := reflect.ValueOf(object)
	for _, v := range t.relations {
		if v.relationkind == ONETOONE {
			checkIfStructEmpty(reflect.TypeOf(v), values.FieldByName(v.relatedto))
		} else if v.relationkind == ONETOMANY {
			ptr := values.FieldByName(v.relatedto)
			for i := 0; i < ptr.Len(); i++ {
				checkIfStructEmpty(reflect.TypeOf(ptr.Index(i)), ptr.Index(i))
			}
		}
	}
	insertCommand := `INSERT INTO %s(%s) VALUES (%s);`
	fields := make([]string, len(t.fields))
	vals := make([]string, len(t.fields))
	for i, v := range t.fields {
		val := values.FieldByName(v.name)
		if val.Kind() == reflect.String {
			vals[i] = fmt.Sprintf("'%s'", val)
		} else {
			vals[i] = fmt.Sprintf("%v", val)
		}
		fields[i] = v.name

	}
	return fmt.Sprintf(insertCommand, t.tableName, strings.Join(fields, ","), strings.Join(vals, ","))
}
