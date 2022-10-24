package diccionario_test

import (
	TDADiccionario "diccionario"
	//"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func mayor(clave1, clave2 int) int {
	return clave1 - clave2
}

func TestDiccionarioVacio(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, string](mayor)

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Obtener(2) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicc.Borrar(2) })
	require.False(t, dicc.Pertenece(2))
	require.EqualValues(t, 0, dicc.Cantidad())

}

func TestDiccionarioUnElemento(t *testing.T) {
	dicc := TDADiccionario.CrearABB[int, int](mayor)

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
	dicc := TDADiccionario.CrearABB[int, string](mayor)

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
	require.EqualValues(t, "China", dicc.Borrar(14))
	require.EqualValues(t, 7, dicc.Cantidad())
	require.True(t, dicc.Pertenece(13))
	//Borro un elemento con dos hijos
	require.EqualValues(t, "Peru", dicc.Borrar(5))
	require.EqualValues(t, 6, dicc.Cantidad())
	require.True(t, dicc.Pertenece(3))
	require.True(t, dicc.Pertenece(7))
	require.EqualValues(t, "Italia", dicc.Obtener(7))
	//Borro la raíz
	require.EqualValues(t, "Argentina", dicc.Borrar(10))
	require.EqualValues(t, 5, dicc.Cantidad())
	require.True(t, dicc.Pertenece(7))
	require.True(t, dicc.Pertenece(12))
	require.EqualValues(t, "Italia", dicc.Obtener(7))

}
