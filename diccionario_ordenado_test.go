package diccionario_test

import (
	TDADiccionario "diccionario"
	"strings"

	//"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func mayorEntreInts(clave1, clave2 int) int {
	return clave1 - clave2
}

func mayorEntreStrings(clave1, clave2 string) int {
	return strings.Compare(clave1, clave2)
}

func TestDiccionarioVacio(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, string](mayorEntreInts)
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Obtener(2) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(2) })
	require.False(t, dicc.Pertenece(2))
	require.EqualValues(t, 0, dicc.Cantidad())

}

func TestDiccionarioUnElemento(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, int](mayorEntreInts)
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dicc.Guardar(0, 56)
	require.NotPanics(t, func() { dicc.Obtener(0) })
	require.EqualValues(t, 56, dicc.Obtener(0))
	require.False(t, dicc.Pertenece(6))
	require.EqualValues(t, 1, dicc.Cantidad())
}

func TestUnElementoPrimitivas(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, int](mayorEntreInts)
	t.Log("Comprueba que Diccionario con un elemento funciona correctamente")
	dicc.Guardar(0, 56)
	//actualizamos dato y se mantiene estable
	dicc.Guardar(0, 70)
	require.NotPanics(t, func() { dicc.Obtener(0) })
	require.EqualValues(t, 70, dicc.Obtener(0))
	require.True(t, dicc.Pertenece(0))
	require.EqualValues(t, 1, dicc.Cantidad())
	require.EqualValues(t, 70, dicc.Borrar(0))
	//después de borrar actua como vacio
	require.EqualValues(t, 0, dicc.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(0) })
	//podemos volver a guardar después de borrar único dato
	dicc.Guardar(5, 65)
	require.NotPanics(t, func() { dicc.Obtener(5) })
	require.EqualValues(t, 65, dicc.Obtener(5))
	require.True(t, dicc.Pertenece(5))
}

func TestBorrarVariosElementos(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, string](mayorEntreInts)
	dicc.Guardar(10, "Argentina")
	dicc.Guardar(2, "Brasil")
	dicc.Guardar(5, "Peru")
	dicc.Guardar(3, "Chile")
	dicc.Guardar(1, "EEUU")
	dicc.Guardar(7, "Italia")
	dicc.Guardar(12, "Australia")
	dicc.Guardar(14, "China")
	dicc.Guardar(13, "Mexico")

	require.EqualValues(t, 9, dicc.Cantidad())
	require.True(t, dicc.Pertenece(10))
	require.True(t, dicc.Pertenece(5))
	require.EqualValues(t, "Brasil", dicc.Obtener(2))
	require.EqualValues(t, "Australia", dicc.Obtener(12))

	//Borro una hoja
	require.EqualValues(t, "EEUU", dicc.Borrar(1))
	require.EqualValues(t, 8, dicc.Cantidad())
	//Borro un elemento con un solo hijo
	require.EqualValues(t, "Brasil", dicc.Borrar(2))
	require.EqualValues(t, 7, dicc.Cantidad())
	require.True(t, dicc.Pertenece(3))
	//Borro un elemento con dos hijos
	require.EqualValues(t, "Australia", dicc.Borrar(12))
	require.EqualValues(t, 6, dicc.Cantidad())
	require.True(t, dicc.Pertenece(13))
	require.True(t, dicc.Pertenece(14))
	require.EqualValues(t, "Italia", dicc.Obtener(7))
	//Borro la raíz
	require.EqualValues(t, "Argentina", dicc.Borrar(10))
	require.EqualValues(t, 5, dicc.Cantidad())
	require.True(t, dicc.Pertenece(7))
	require.True(t, dicc.Pertenece(5))
	require.EqualValues(t, "Peru", dicc.Obtener(5))
}

// TESTS HEREDADOS DE LA CATEDRA PARA DICCIONARIOS DE HASH -------->>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](mayorEntreStrings)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](mayorEntreStrings)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](mayorEntreStrings)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](mayorEntreStrings)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](mayorEntreStrings)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](mayorEntreStrings)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](mayorEntreStrings)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](mayorEntreStrings)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	require.EqualValues(t, primero, iter.Siguiente())
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	require.EqualValues(t, segundo, iter.Siguiente())
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	require.EqualValues(t, tercero, iter.Siguiente())

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](mayorEntreStrings)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero := iter3.Siguiente()
	segundo := iter3.Siguiente()
	tercero := iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](mayorEntreStrings)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	require.EqualValues(t, clave1, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
}
