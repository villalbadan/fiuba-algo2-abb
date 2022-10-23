package diccionario

const (
	VALOR_CMP = 0
)

// ################################### ESTRUCTURAS ###############################################################

type funcCmp[K comparable] func(K, K) int

type ab[K comparable, V any] struct {
	raiz     *nodoAb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type nodoAb[K comparable, V any] struct {
	izq   *nodoAb[K, V]
	der   *nodoAb[K, V]
	clave K
	dato  V
}

type iteradorDict[K comparable, V any] struct {
	diccionario
	actual
	desde
	hasta
}

// ##############################################################################################################

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	dict := new(ab[K, V])
	dict.cmp = funcion_cmp[K, K]
	return dict
}

/* func Comparacion[K comparable](K, K)
¿Hay que definirla?
*/

func (dict *ab[K, V]) buscar(clave K, nodoActual **nodoAb[K, V]) **nodoAb[K, V] {
	if (*nodoActual) == nil {
		return nodoActual
	}

	comparacion := dict.cmp(clave, (*nodoActual).clave)
	//clave a evaluar es menor a la clave actual
	if comparacion < VALOR_CMP {
		return dict.buscar(clave, &(*nodoActual).izq)
	}

	//clave a evaluar es mayor a la clave actual
	if comparacion > VALOR_CMP {
		return dict.buscar(clave, &(*nodoActual).der)
	}

	//clave a evaluar es la clave del nodo
	return nodoActual

}

// ################################### Aux. Borrar #########################################################

//no estoy segura de esta función, se hizo lo que se pudo
func (dict *ab[K, V]) reemplazante(nodoActual **nodoAb[K, V]) *nodoAb[K, V] {

	if *nodoActual == nil {
		return nil
	}

	if noTieneHijos(nodoActual) {
		return *nodoActual
	}

	nodo := dict.reemplazante(&(*nodoActual).izq)
	if nodo == nil {
		nodo = dict.reemplazante(&(*nodoActual).der)
	}
	return nodo
}

func noTieneHijos[K comparable, V any](nodo **nodoAb[K, V]) bool {
	return (*nodo).izq == nil && (*nodo).der == nil
}

func (dict *ab[K, V]) transplantar(nodo **nodoAb[K, V]) {

	//nodo con un solo hijo
	if (*nodo).izq == nil && (*nodo).der != nil {
		*nodo = (*nodo).der
		return
	}

	if (*nodo).izq != nil && (*nodo).der == nil {
		*nodo = (*nodo).izq
		return
	}

	//nodo con dos hijos
	//Busco reemplazante menor
	nuevoNodo := dict.reemplazante(&(*nodo).izq)
	nuevaClave := nuevoNodo.clave
	nuevoDato := dict.Borrar(nuevaClave)

	//piso datos
	(*nodo).clave = nuevaClave
	(*nodo).dato = nuevoDato
}

// ################################### PRIMITIVAS DICCIONARIO #################################################

func (dict *ab[K, V]) Guardar(clave K, dato V) {
	nodo := dict.buscar(clave, &dict.raiz)
	if *nodo != nil {
		(*nodo).dato = dato
		return
	}

	*nodo = &nodoAb[K, V]{clave: clave, dato: dato}
	dict.cantidad++
}

func (dict *ab[K, V]) Pertenece(clave K) bool {
	return *(dict.buscar(clave, &dict.raiz)) != nil
}

func (dict *ab[K, V]) Obtener(clave K) V {
	nodo := dict.buscar(clave, &dict.raiz)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).dato
}

func (dict *ab[K, V]) Cantidad() int {
	return dict.cantidad
}

func (dict *ab[K, V]) Borrar(clave K) V {
	nodo := dict.buscar(clave, &dict.raiz)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	//dato del nodo a borrar
	borrado := (*nodo).dato
	dict.cantidad--

	if noTieneHijos(nodo) {
		//nodo sin hijos
		*nodo = nil
	} else {
		//nodo con hijos
		dict.transplantar(nodo)
	}

	return borrado

}

func (dict *ab[K, V]) Iterar(visitar func(K, V) bool) {

}

// ################################### PRIMITIVAS ITERADOR EXTERNO ################################################

func (dict *ab[K, V]) Iterador() IterDiccionario[K, V] {
	return
}

func (iter *iteradorDict[K, V]) HaySiguiente() bool {
	return
}

func (iter *iteradorDict[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return
}

func (iter *iteradorDict[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return
}

// ################################### PRIMITIVAS DICCIONARIO ORDENADO #############################################

func IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

}
