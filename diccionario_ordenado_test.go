package diccionario_test

import (
	TDADiccionario "diccionario"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

var (
	TAMS_VOLUMEN    = []int{12500, 25000, 50000, 100000, 200000, 400000}
	ARREGLO_STRINGS = []string{"G", "K", "M", "B", "C", "W", "O", "A", "V", "F"}
	ARREGLO_INTS    = []int{4, 5, 6, 1, 2, 9, 7, 0, 8, 3}
)

//FUNC CMP -----------------------------------------------------------------------------------------------------------

func mayorEntreInts(clave1, clave2 int) int {
	return clave1 - clave2
}

func mayorEntreStrings(clave1, clave2 string) int {
	return strings.Compare(clave1, clave2)
}

//--------------------------------------------------------------------------------------------------------------------

func armarDictDeClavesInt[V any](arregloClaves []int, arregloDatos []V) TDADiccionario.DiccionarioOrdenado[int, V] {
	dicc := TDADiccionario.CrearABB[int, V](mayorEntreInts)
	for i := 0; i < len(arregloClaves); i++ {
		dicc.Guardar(arregloClaves[i], arregloDatos[i])
	}
	return dicc
}

func armarDictDeClavesString[V any](arregloClaves []string, arregloDatos []V) TDADiccionario.DiccionarioOrdenado[string, V] {
	dicc := TDADiccionario.CrearABB[string, V](mayorEntreStrings)
	for i := 0; i < len(arregloClaves); i++ {
		dicc.Guardar(arregloClaves[i], arregloDatos[i])
	}
	return dicc
}

func TestDiccionarioVacio(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, string](mayorEntreInts)
	t.Log("Diccionario recién creado actua como vacio")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Obtener(2) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(2) })
	require.False(t, dicc.Pertenece(2))
	require.EqualValues(t, 0, dicc.Cantidad())

}

func TestDiccionarioUnElemento(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, int](mayorEntreInts)
	t.Log("Guardar y borrar un solo elemento en dicc recién creado")
	dicc.Guardar(0, 56)
	require.NotPanics(t, func() { dicc.Obtener(0) })
	require.EqualValues(t, 56, dicc.Obtener(0))
	require.True(t, dicc.Pertenece(0))
	require.EqualValues(t, 1, dicc.Cantidad())
	require.EqualValues(t, 56, dicc.Borrar(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(0) })
}

func TestBorrarVariosElementos(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, string](mayorEntreInts)
	t.Log("Borrar elementos con y sin hijos")
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
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(1) })
	require.EqualValues(t, 8, dicc.Cantidad())
	//Borro un elemento con un solo hijo
	require.EqualValues(t, "China", dicc.Borrar(14))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(14) })
	require.EqualValues(t, 7, dicc.Cantidad())
	require.True(t, dicc.Pertenece(13))
	//Borro un elemento con dos hijos
	require.EqualValues(t, "Peru", dicc.Borrar(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(5) })
	require.EqualValues(t, 6, dicc.Cantidad())
	require.True(t, dicc.Pertenece(3))
	require.True(t, dicc.Pertenece(7))
	require.EqualValues(t, "Italia", dicc.Obtener(7))
	//Borro la raíz
	require.EqualValues(t, "Argentina", dicc.Borrar(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(10) })
	require.EqualValues(t, 5, dicc.Cantidad())
	require.True(t, dicc.Pertenece(7))
	require.True(t, dicc.Pertenece(12))
	require.EqualValues(t, "Italia", dicc.Obtener(7))

}

func TestUnElementoPrimitivas(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, int](mayorEntreInts)
	t.Log("Primitivas funcionan en un dicc con un único elemento")
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

func TestReemplazoDatos(t *testing.T) {

	dicc := TDADiccionario.CrearABB[string, string](strings.Compare)
	t.Log("Reemplazar datos en claves ya guardadas")
	dicc.Guardar("Belgrano", "Azul")
	dicc.Guardar("Sarmiento", "Verde")
	dicc.Guardar("San Martin", "Rojo")
	require.EqualValues(t, "Verde", dicc.Obtener("Sarmiento"))
	require.EqualValues(t, 3, dicc.Cantidad())
	dicc.Guardar("Sarmiento", "Marron")
	require.EqualValues(t, "Marron", dicc.Obtener("Sarmiento"))
	require.EqualValues(t, 3, dicc.Cantidad())
	dicc.Guardar("Rosas", "Amarillo")
	require.EqualValues(t, 4, dicc.Cantidad())
	dicc.Guardar("Belgrano", "Blanco")
	require.EqualValues(t, 4, dicc.Cantidad())
	require.EqualValues(t, "Blanco", dicc.Obtener("Belgrano"))
	require.EqualValues(t, "Amarillo", dicc.Obtener("Rosas"))

}

func TestDiccionarioTipoLista(t *testing.T) {

	dicc := TDADiccionario.CrearABB[int, int](mayorEntreInts)
	t.Log("Guardado de claves ordenadas")
	dicc.Guardar(0, 34)
	dicc.Guardar(1, 3)
	dicc.Guardar(2, 5)
	dicc.Guardar(3, 8)
	dicc.Guardar(5, -2)
	dicc.Guardar(7, 96)
	dicc.Guardar(10, 24)

	require.True(t, dicc.Pertenece(7))
	require.True(t, dicc.Pertenece(10))

}

func TestIteraEnOrden(t *testing.T) {
	t.Log("Iterador interno itera en orden")
	arregloClaves := []string{"G", "K", "M", "B", "C", "W", "O", "A", "V", "F"}
	dicc := armarDictDeClavesString(arregloClaves, ARREGLO_INTS)
	sort.Strings(arregloClaves)
	i := 0
	dicc.Iterar(func(clave string, dato int) bool {
		require.EqualValues(t, i, dato)
		i++
		return true
	})

}

func TestIterarConRangoFueraDeDict(t *testing.T) {
	t.Log("Iterador interno con rango fuera del dict itera correctamente")
	dicc := armarDictDeClavesInt(ARREGLO_INTS, ARREGLO_STRINGS)

	inicio := 1
	fin := 15
	inicioPtr := &inicio
	finPtr := &fin

	dicc.IterarRango(inicioPtr, finPtr, func(clave int, dato string) bool {
		return true
	})

}

func TestIterarConSinDesde(t *testing.T) {
	t.Log("Iterador interno con rango fuera del dict itera correctamente")
	dicc := armarDictDeClavesInt(ARREGLO_INTS, ARREGLO_STRINGS)

	fin := 4
	finPtr := &fin
	claves := []int{0, 1, 2, 3, 4}
	i := 0
	iptr := &i
	dicc.IterarRango(nil, finPtr, func(clave int, dato string) bool {
		require.EqualValues(t, claves[i], clave)
		*iptr = *iptr + 1
		return true
	})

}

func TestIterarConSinHasta(t *testing.T) {
	t.Log("Iterador interno con rango fuera del dict itera correctamente")
	dicc := armarDictDeClavesInt(ARREGLO_INTS, ARREGLO_STRINGS)

	inicio := 6
	inicioPtr := &inicio
	claves := []int{6, 7, 8, 9}
	i := 0
	iptr := &i
	dicc.IterarRango(inicioPtr, nil, func(clave int, dato string) bool {
		require.EqualValues(t, claves[i], clave)
		*iptr = *iptr + 1
		return true
	})

}

func TestIterarRangoPara(t *testing.T) {
	t.Log("Iterador interno con rango fuera del dict itera correctamente")
	dicc := armarDictDeClavesInt(ARREGLO_INTS, ARREGLO_STRINGS)

	inicio := 6
	inicioPtr := &inicio
	var ultimaClave int
	ultimaPtr := &ultimaClave
	dicc.IterarRango(inicioPtr, nil, func(clave int, dato string) bool {
		*ultimaPtr = clave
		return clave != 8
	})

	require.EqualValues(t, 8, ultimaClave)

}

func TestIterarParaEnFalse(t *testing.T) {
	t.Log("Iterador interno con rango fuera del dict itera correctamente")
	dicc := armarDictDeClavesInt(ARREGLO_INTS, ARREGLO_STRINGS)
	var ultimaClave int
	ultimaPtr := &ultimaClave

	dicc.Iterar(func(clave int, dato string) bool {
		*ultimaPtr = clave
		return clave != 3
	})
	require.EqualValues(t, 3, ultimaClave)

}

func TestIterarConRangoClavesEnDict(t *testing.T) {
	t.Log("Iterador interno con rango, claves de inicio y fin existentes en el dict")
	dicc := armarDictDeClavesInt(ARREGLO_INTS, ARREGLO_STRINGS)

	inicio := 6
	inicioPtr := &inicio
	fin := 9
	finPtr := &fin
	claves := []int{6, 7, 8, 9}

	i := 0
	iptr := &i
	dicc.IterarRango(inicioPtr, finPtr, func(clave int, dato string) bool {
		require.EqualValues(t, claves[i], clave)
		*iptr = *iptr + 1
		return true
	})

}

func TestIterarConRangoClavesOrdenadas(t *testing.T) {
	t.Log("Iterador interno con rango, claves ingresan ordenadas")
	arregloClaves := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	dicc := armarDictDeClavesInt(arregloClaves, ARREGLO_STRINGS)

	inicio := 2
	inicioPtr := &inicio
	fin := 5
	finPtr := &fin
	claves := []int{2, 3, 4, 5}

	i := 0
	iptr := &i
	dicc.IterarRango(inicioPtr, finPtr, func(clave int, dato string) bool {
		require.EqualValues(t, claves[i], clave)
		*iptr = *iptr + 1
		return true
	})

}

func TestIterarSumaCorrectamente(t *testing.T) {
	t.Log("Iterador interno realiza la suma de todos los datos del dicc")
	dicc := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := []string{"G", "K", "M", "B"}
	datos := []int{4, 5, 6, 1}

	for i := 0; i < len(claves); i++ {
		dicc.Guardar(claves[i], datos[i])
	}
	suma := 0
	sumaPtr := &suma

	dicc.Iterar(func(clave string, dato int) bool {
		*sumaPtr = *sumaPtr + dato
		return true
	})

	require.EqualValues(t, 16, suma)

}

func TestIterarConRangoSuma(t *testing.T) {
	t.Log("Iterador interno con rango realiza la suma de todos los datos del rango")
	claves := []string{"G", "K", "M", "B"}
	datos := []int{4, 5, 6, 1}
	dicc := armarDictDeClavesString(claves, datos)

	suma := 0
	sumaPtr := &suma
	inicio := "G"
	fin := "N"

	dicc.IterarRango(&inicio, &fin, func(clave string, dato int) bool {
		*sumaPtr = *sumaPtr + dato
		return true
	})

	require.EqualValues(t, 15, suma)

}

// TEST EXTERNO --------------------------------------------------------------------------

func TestIteradorVacio(t *testing.T) {
	t.Log("Se puede crear iterador de diccionario vacio")
	dicc := TDADiccionario.CrearABB[int, int](mayorEntreInts)
	iter := dicc.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorEnOrden(t *testing.T) {
	t.Log("Iterador externo itera en orden")

	arregloClaves := []int{15, 12, 8, 2, 4, 34, 26, 51, 16, 22, 45, 30}
	arregloDatos := []string{"Vaca", "Gato", "Perro", "Pollo", "Raton", "Paloma",
		"Avestruz", "Pez", "Tigre", "Leon", "Jirafa", "Oso"}
	dicc := armarDictDeClavesInt(arregloClaves, arregloDatos)

	iter := dicc.Iterador()
	sort.Ints(arregloClaves)

	for i := 0; i < len(arregloClaves); i++ {
		require.EqualValues(t, arregloClaves[i], iter.Siguiente())
	}

}

func TestIteradorConRango(t *testing.T) {
	t.Log("Iterador con rango itera en orden, con desde y hasta claves no existentes en el dict")
	arregloClaves := []int{8, 1, 24, 12, 32, 5, 149, 224}
	arregloDatos := []string{"A", "Ante", "Bajo", "Contra", "Desde", "Entre", "Hacia", "Hasta"}
	dicc := armarDictDeClavesInt(arregloClaves, arregloDatos)
	desde := 10
	hasta := 100

	iterRango := dicc.IteradorRango(&desde, &hasta)
	sort.Ints(arregloClaves)
	arregloClaves = arregloClaves[3:6] // Valores que entran dentro del rango
	primeraClave, _ := iterRango.VerActual()
	require.EqualValues(t, 12, primeraClave) // para verificar que se eligio bien el primer elemento

	for i := 0; iterRango.HaySiguiente(); i++ {
		require.EqualValues(t, arregloClaves[i], iterRango.Siguiente())
	}

}

func TestIteradorConRangoClavesExistentes(t *testing.T) {
	t.Log("Iterador con rango itera en orden, con desde y hasta claves existentes en el dict")
	arregloClaves := []int{8, 1, 24, 12, 32, 5, 149, 224}
	arregloDatos := []string{"A", "Ante", "Bajo", "Contra", "Desde", "Entre", "Hacia", "Hasta"}
	dicc := armarDictDeClavesInt(arregloClaves, arregloDatos)

	desde := 5
	hasta := 32

	iterRango := dicc.IteradorRango(&desde, &hasta)
	claves := []int{5, 8, 12, 24, 32}
	primeraClave, _ := iterRango.VerActual()
	require.EqualValues(t, 5, primeraClave) // para verificar que se eligio bien el primer elemento

	for i := 0; iterRango.HaySiguiente(); i++ {
		require.EqualValues(t, claves[i], iterRango.Siguiente())
	}

}

// PRUEBAS DE VOLUMEN -----------------------------------------------------------------------------------------------
//Basadas en las pruebas de benchmark de hash de la cátedra con algunas modificaciones para que se ingresen elementos
//desordenados e itere ordenadamente

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func listaNumerosRandoms(n int) []int {

	claves := make([]int, n)
	for i := 0; i < n; i++ {
		claves[i] = i
	}

	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		swap(&claves[i], &claves[j])
	}

	return claves
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[int, string](mayorEntreInts)
	claves := listaNumerosRandoms(n)
	valores := make([]string, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		valores[i] = fmt.Sprintf("%08d", claves[i])
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[int, string](mayorEntreInts)

	claves := listaNumerosRandoms(n)
	valores := make([]string, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		valores[i] = fmt.Sprintf("%08d", claves[i])
		dic.Guardar(claves[i], valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	sort.Ints(claves)
	ok := true
	var i int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}

		c1, _ := iter.VerActual()
		if c1 != claves[i] {
			ok = false
			break
		}
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

}

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
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
