package vinago_test

import (
	"testing"
	"VinaGo"
)


func TestNZ(t *testing.T) {
	vinastr := vinago.NewVina()
	obj1 := vinago.NewObject("Object")
	obj1.AddKey("obj1", vinago.NewInt(123))
	obj1.AddKey("obj2", vinago.NewString("char"))
	obj1.AddKey("牛至", vinago.NewString("鸿"))
	obj1.AddKey("obj3", vinago.NewFloat(0.123))

	obj2 := vinago.NewObject("AnotherObj")
	obj2.AddKey("abc", vinago.NewString("abc"))

	vinastr.AddObject(*obj1).AddObject(*obj2)
	if vinastr.String()!="Object{obj1(123),obj2(\"char\"),牛至(\"鸿\"),obj3(0.123)}\nAnotherObj{abc(\"abc\")}"{
		t.Fatalf("%s",vinastr.String())
	}

}
