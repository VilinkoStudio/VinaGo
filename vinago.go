package vinago

import (
	"fmt"
	"reflect"
	"strings"
)
type VinaInt=int64
type VinaString=string
type VinaFloat=float64

// 字符串类型
type strObject struct {
	Name  string
	Value VinaString
}

func (so strObject) String() string {
	return fmt.Sprintf(`%s("%s")`, so.Name, so.Value)
}

// 整数类型
type intObject struct {
	Name  string
	Value VinaInt
}

func (io intObject) String() string {
	return fmt.Sprintf("%s(%d)", io.Name, io.Value)
}

// 浮点数类型
type floatObject struct {
	Name  string
	Value VinaFloat
}

func (fo floatObject) String() string {
	return fmt.Sprintf("%s(%g)", fo.Name, fo.Value)
}

type VinaValue = fmt.Stringer

// 嵌套类型
type Object struct {
	Name   string
	Values []VinaValue
}

func (oo *Object) String() string {
	var values []string
	for _, value := range oo.Values {
		values = append(values, value.String())
	}
	return fmt.Sprintf("%s{%s}", oo.Name, strings.Join(values, ","))
}

// 添加通用的内容
func (oo *Object) addElement(obj VinaValue) *Object {
	oo.Values = append(oo.Values, obj)
	return oo
}

// ObjectBuilder 构造Object树
type ObjectBuilder struct {
	Object
}

func NewObjectBuilder(name string) *ObjectBuilder {
	return &ObjectBuilder{
		Object: Object{
			Name:   name,
			Values: []VinaValue{},
		},
	}
}
func (obj *ObjectBuilder) AddIntKey(name string, value VinaInt) *ObjectBuilder {
	obj.addElement(intObject{Name: name, Value: value})
	return obj
}

func (obj *ObjectBuilder) AddStringKey(name string, value VinaString)*ObjectBuilder {
	obj.addElement(strObject{Name: name, Value: value})
	return obj
}

func (obj *ObjectBuilder) AddFloatKey(name string, value VinaFloat) *ObjectBuilder {
	obj.addElement(floatObject{Name: name, Value: value})
	return obj
}
func (obj *ObjectBuilder) AddKey(name string, value any) *ObjectBuilder {
	switch v := value.(type) {
	case int:
		obj.AddIntKey(name, VinaInt(v))
	case int8:
		obj.AddIntKey(name, VinaInt(v))
	case int16:
		obj.AddIntKey(name, VinaInt(v))
	case int32:
		obj.AddIntKey(name, VinaInt(v))
	case int64:
		obj.AddIntKey(name, VinaInt(v))
	case uint:
		obj.AddIntKey(name, VinaInt(v))
	case uint8:
		obj.AddIntKey(name, VinaInt(v))
	case uint16:
		obj.AddIntKey(name, VinaInt(v))
	case uint32:
		obj.AddIntKey(name, VinaInt(v))
	case uint64:
		obj.AddIntKey(name, VinaInt(v))
	case string:
		obj.AddStringKey(name, VinaString(v))
	case float64:
		obj.AddFloatKey(name, VinaFloat(v))
	case float32:
		obj.AddFloatKey(name, VinaFloat(v))
	default:
		panic(fmt.Sprintf("Unsupported type: %T", v))
	}
	return obj
}
func (obj *ObjectBuilder) SerializeStruct(input any) *ObjectBuilder {
	Type := reflect.TypeOf(input)
	if Type.Kind() != reflect.Struct {
		panic("Input must be a struct")
	}
	Value := reflect.ValueOf(input)
	for i := 0; i < Type.NumField(); i++ {
		field := Type.Field(i)
		if field.PkgPath != "" { // unexported field, skip
			continue
		}
		fieldValue := Value.Field(i)
		obj.AddKey(field.Name, fieldValue.Interface())
	}
	return obj
}
func (obj *ObjectBuilder) SerializeMap(input map[string]any) *ObjectBuilder {
	for key, value := range input {
		obj.AddKey(key, value)
	}
	return obj
}
func (obj *ObjectBuilder) Build() Object {
	return obj.Object
}

// vina树
type Vina struct {
	Objects []Object
}

func (v *Vina) String() string {
	var objs []string
	for _, obj := range v.Objects {
		objs = append(objs, obj.String())
	}
	return strings.Join(objs, "\n")
}

// VinaBuilder 构造整个vina树
type VinaBuilder struct {
	Vina
}

func CreateVinaBuilder() *VinaBuilder {
	return &VinaBuilder{
		Vina: Vina{
			Objects: []Object{},
		},
	}
}

func (vi *VinaBuilder) AddObject(obj Object) *VinaBuilder {
	vi.Vina.Objects = append(vi.Vina.Objects, obj)
	return vi
}

func (vi *VinaBuilder) Build() Vina {
	return vi.Vina
}
