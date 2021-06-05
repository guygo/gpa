package repoistorydb

type Sqltype int

const (
	Invalid Sqltype = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

var SqlTypeNames = []string{
	Invalid:       "invalid",
	Bool:          "bit",
	Int:           "int",
	Int8:          "tinyint",
	Int16:         "smallint",
	Int32:         "int",
	Int64:         "bigint",
	Uint:          "tinyint",
	Uint8:         "invalid",
	Uint16:        "invalid",
	Uint32:        "invalid",
	Uint64:        "invalid",
	Uintptr:       "invalid",
	Float32:       "real",
	Float64:       "float",
	Complex64:     "invalid",
	Complex128:    "invalid",
	Array:         "array",
	Chan:          "invalid",
	Func:          "invalid",
	Interface:     "invalid",
	Map:           "invalid",
	Ptr:           "invalid",
	Slice:         "array",
	String:        "varchar",
	Struct:        "struct",
	UnsafePointer: "invalid",
}
