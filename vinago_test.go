package vinago_test

import (
	"github.com/VilinkoStudio/VinaGo"
	"fmt"
	"testing"
)


func TestNZ(t *testing.T) {
	vina_builder:=vinago.CreateVinaBuilder()
	obj_builder:=vinago.NewObjectBuilder("Object")
	obj_builder.AddIntKey("obj1", 123)
	obj_builder.AddStringKey("obj2", "char")
	obj_builder.AddStringKey("牛至", "鸿")
	obj_builder.AddFloatKey("obj3", 0.123)


	obj_builder_2:=vinago.NewObjectBuilder("AnotherObj")
	obj_builder_2.AddStringKey("abc", "abc")

	vina_builder.AddObject(obj_builder.Build()).AddObject(obj_builder_2.Build())
	vina := vina_builder.Build()
	if vina.String() != "Object{obj1(123),obj2(\"char\"),牛至(\"鸿\"),obj3(0.123)}\nAnotherObj{abc(\"abc\")}" {
		t.Fatalf("%s", vina.String())
	}

}
func TestNZWithAddKey(t *testing.T) {
	vina_builder:=vinago.CreateVinaBuilder()
	obj_builder:=vinago.NewObjectBuilder("Object")
	obj_builder.AddKey("obj1", 123)
	obj_builder.AddKey("obj2", "char")
	obj_builder.AddKey("牛至", "鸿")
	obj_builder.AddKey("obj3", 0.123)

	obj_builder_2:=vinago.NewObjectBuilder("AnotherObj")
	obj_builder_2.AddKey("abc", "abc")

	vina_builder.AddObject(obj_builder.Build()).AddObject(obj_builder_2.Build())
	vina := vina_builder.Build()
	if vina.String() != "Object{obj1(123),obj2(\"char\"),牛至(\"鸿\"),obj3(0.123)}\nAnotherObj{abc(\"abc\")}" {
		t.Fatalf("%s", vina.String())
	}
}
func TestSerializeStruct(t *testing.T) {
	vina_builder := vinago.CreateVinaBuilder()
	obj_builder := vinago.NewObjectBuilder("Object")
	anon_struct := struct {
		Obj1 int
		Obj2 string
		Obj3 float64
	}{
		Obj1: 123,
		Obj2: "char",
		Obj3: 0.123,
	}
	obj_builder.SerializeStruct(anon_struct)

	obj_builder_2 := vinago.NewObjectBuilder("AnotherObj")
	anon_struct_2 := struct {
		Abc string
	}{
		Abc: "abc",
	}
	obj_builder_2.SerializeStruct(anon_struct_2)

	vina_builder.AddObject(obj_builder.Build()).AddObject(obj_builder_2.Build())
	vina := vina_builder.Build()
	expected := "Object{Obj1(123),Obj2(\"char\"),Obj3(0.123)}\nAnotherObj{Abc(\"abc\")}"
	if vina.String() != expected {
		t.Fatalf("%s", vina.String())
	}
}
func TestSerializeMap(t *testing.T) {
	vina_builder := vinago.CreateVinaBuilder()
	obj_builder := vinago.NewObjectBuilder("Object")
	anon_map := map[string]any{
		"Obj1": 123,
		"Obj2": "char",
		"Obj3": 0.123,
	}
	obj_builder.SerializeMap(anon_map)

	obj_builder_2 := vinago.NewObjectBuilder("AnotherObj")
	anon_map_2 := map[string]any{
		"Abc": "abc",
	}
	obj_builder_2.SerializeMap(anon_map_2)

	vina_builder.AddObject(obj_builder.Build()).AddObject(obj_builder_2.Build())
	vina := vina_builder.Build()
	fmt.Println(vina.String())				//map不保证顺序 所以只做输出了
}