package vinago

import (
	"fmt"
	"strconv"
	"strings"
)

// Keys里面的值
type Value struct {
	Type  string
	Value interface{}
}

// Keys
type Keys struct {
	Name  string
	Value Value
}

// 最大的Obj
type Object struct {
	Name string
	Keys []Keys
}

// vina树
type Vina struct {
	Objects []Object
}

func NewInt(val int) Value {
	return Value{Type: "int", Value: val}
}

func NewString(val string) Value {
	return Value{Type: "string", Value: val}
}

func NewFloat(val float64) Value {
	return Value{Type: "float", Value: val}
}

// Object 构造函数
func NewObject(name string) *Object {
	return &Object{
		Name: name,
		Keys: make([]Keys, 0),
	}
}

func (obj *Object) AddKey(name string, value Value) *Object {
	obj.Keys = append(obj.Keys, Keys{Name: name, Value: value})
	return obj
}

// Vina 构造函数
func NewVina() *Vina {
	return &Vina{
		Objects: make([]Object, 0),
	}
}

func (vf *Vina) AddObject(obj Object) *Vina {
	vf.Objects = append(vf.Objects, obj)
	return vf
}

func (val Value) String() string {
	switch val.Type {
	case "int":
		return strconv.Itoa(val.Value.(int))
	case "string":
		return fmt.Sprintf(`"%s"`, val.Value.(string))
	case "float":
		return strconv.FormatFloat(val.Value.(float64), 'f', -1, 64)
	default:
		return ""
	}
}

func (key Keys) String() string {
	return fmt.Sprintf("%s(%s)", key.Name, key.Value.String())
}

func (obj Object) String() string {
	var Keys []string
	for _, Key := range obj.Keys {
		Keys = append(Keys, Key.String())
	}
	return fmt.Sprintf("%s{%s}", obj.Name, strings.Join(Keys, ","))
}

// 新建vina
func (vi Vina) String() string {
	var objects []string
	for _, obj := range vi.Objects {
		objects = append(objects, obj.String())
	}
	return strings.Join(objects, "\n")
}

// ObjectBuilder 构造Object树
type ObjectBuilder struct {
	object *Object
}

func NewObjectBuilder(name string) *ObjectBuilder {
	return &ObjectBuilder{
		object: NewObject(name),
	}
}

func (obj *ObjectBuilder) AddIntKey(name string, value int) *ObjectBuilder {
	obj.object.AddKey(name, NewInt(value))
	return obj
}

func (obj *ObjectBuilder) AddStringKey(name string, value string) *ObjectBuilder {
	obj.object.AddKey(name, NewString(value))
	return obj
}

func (obj *ObjectBuilder) AddFloatKey(name string, value float64) *ObjectBuilder {
	obj.object.AddKey(name, NewFloat(value))
	return obj
}

func (obj *ObjectBuilder) Build() Object {
	return *obj.object
}

// VinaBuilder 构造整个vina树
type VinaBuilder struct {
	vina *Vina
}

func CreateVinaBuilder() *VinaBuilder {
	return &VinaBuilder{
		vina: NewVina(),
	}
}

func (vi *VinaBuilder) AddObject(obj Object) *VinaBuilder {
	vi.vina.AddObject(obj)
	return vi
}

func (vi *VinaBuilder) Build() Vina {
	return *vi.vina
}
