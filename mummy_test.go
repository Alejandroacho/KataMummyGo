package main
import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDevuelveUnArray(test *testing.T){
	text:="Holi"
	expected:=[]string{"H","o","l","i"}
	result:=stringToArray(text)
	assert.Equal(test,expected,result)
}
func TestCuentaArray(test *testing.T){
	splited:=[]string{"H","o","l","i"}
	expected:= 4
	result:=countArray(splited)
	assert.Equal(test,expected,result)
}
func TestObtienePorcentajeArray(test *testing.T){
	splited:=[]string{"H","o","l","i"}
	count:= 4
	expected:=50
	result:=obtenerPorcentajeDeVocales(splited, count)
	assert.Equal(test,expected,result)
}
func TestObtieneSubstituye(test *testing.T){
	splited:=[]string{"H","o","l","i"}
	count:= 4
	percentage:=50
	result:=sustituirPorMummy(splited, count, percentage)
	expected:=[]string{"H","mummy","l","mummy"}
	assert.Equal(test,expected,result)
}
func TestBorraDobles(test *testing.T){
	splited:=[]string{"H","mummy","mummy"}
	count:= 3
	result:=buscarDobleMummy(splited, count)
	expected:=[]string{"H","mummy",""}
	assert.Equal(test,expected,result)
}
func TestBorraDobles2(test *testing.T){
	splited:=[]string{"H","mummy","mummy","l","mummy"}
	count:= 3
	result:=buscarDobleMummy(splited, count)
	expected:=[]string{"H","mummy","","l","mummy"}
	assert.Equal(test,expected,result)
}
func TestConvierteArrayEnString(test *testing.T){
	splited:=[]string{"H","o","l","i"}
	expected:="Holi"
	result:=arrayToString(splited)
	assert.Equal(test,expected,result)
}

func TestTodo(test *testing.T){
	text:="holi"
	result:=todo(text)
	expected:="hmummylmummy"
	assert.Equal(test,expected,result)
}

func TestTodo2(test *testing.T){
	text:="hooli"
	result:=todo(text)
	expected:="hmummylmummy"
	assert.Equal(test,expected,result)
}
