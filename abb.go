package diccionario

import (
	TDAPila "diccionario/pila"
)

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
	diccionario   *ab[K, V]
	actual        *nodoAb[K, V]
	rangoMax      *K
	pilaElementos TDAPila.Pila[*nodoAb[K, V]]
}

// ##############################################################################################################

func CrearABB[K comparable, V any](funcion_cmp funcCmp[K]) DiccionarioOrdenado[K, V] {
	dict := new(ab[K, V])
	dict.cmp = funcion_cmp
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

func noTieneHijos[K comparable, V any](nodo **nodoAb[K, V]) bool {
	return (*nodo).izq == nil && (*nodo).der == nil
}

func (dict *ab[K, V]) reemplazante(nodoActual **nodoAb[K, V]) *nodoAb[K, V] {
	if noTieneHijos(nodoActual) || (*nodoActual).izq == nil {
		return *nodoActual
	}
	return dict.reemplazante(&(*nodoActual).izq)
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
	//Busco reemplazante menor de la derecha
	nuevoNodo := dict.reemplazante(&(*nodo).der)
	nuevaClave := nuevoNodo.clave
	nuevoDato := dict.Borrar(nuevaClave)
	dict.cantidad++ //Para contrarrestar el Borrar de la linea de arriba

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

//Iterador interno ------------------------------------------------->>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (dict ab[K, V]) Iterar(visitar func(K, V) bool) {
	dict.raiz.iterar(visitar)
}

func (nodo *nodoAb[K, V]) iterar(visitar func(K, V) bool) {
	if nodo == nil {
		return
	}
	if nodo.izq != nil {
		nodo.izq.iterar(visitar)
	}
	if !visitar(nodo.clave, nodo.dato) {
		return
	}
	if nodo.der != nil {
		nodo.der.iterar(visitar)
	}

}

// ################################### PRIMITIVAS ITERADOR EXTERNO ################################################
func (dict *ab[K, V]) crearIter(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := iteradorDict[K, V]{diccionario: dict, rangoMax: hasta}
	iter.pilaElementos = TDAPila.CrearPilaDinamica[*nodoAb[K, V]]()
	if desde == nil {
		dict.raiz.buscarHijosIzquierdayApilar(iter.pilaElementos)
	} else {
		nodoInicial := dict.raiz.buscarMinimo(iter.pilaElementos, dict.cmp, desde)
		nodoInicial.buscarHijosIzquierdayApilar(iter.pilaElementos)
	}
	return &iter
}

func (nodo *nodoAb[K, V]) buscarHijosIzquierdayApilar(pila TDAPila.Pila[*nodoAb[K, V]]) *nodoAb[K, V] {
	if nodo == nil {
		return nil
	}
	pila.Apilar(nodo)
	return nodo.izq.buscarHijosIzquierdayApilar(pila)
}

func (dict *ab[K, V]) Iterador() IterDiccionario[K, V] {
	return dict.crearIter(nil, nil)
}

func (iter *iteradorDict[K, V]) HaySiguiente() bool {
	return !iter.pilaElementos.EstaVacia()
}

func (iter *iteradorDict[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave, dato := iter.pilaElementos.VerTope().clave, iter.pilaElementos.VerTope().dato
	return clave, dato
}

func (iter *iteradorDict[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoActual := iter.pilaElementos.Desapilar()
	if nodoActual.der != nil {
		nodoActual.der.buscarHijosIzquierdayApilar(iter.pilaElementos)
	}
	return nodoActual.clave
}

// ################################### PRIMITIVAS DICCIONARIO ORDENADO #############################################

func (dict ab[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return dict.crearIter(desde, hasta)

}

func (dict ab[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	dict.raiz.iterarRango(desde, hasta, visitar, dict.cmp)
}

func (nodo *nodoAb[K, V]) iterarRango(desde *K, hasta *K, visitar func(K, V) bool, cmp funcCmp[K]) {
	if nodo == nil {
		return
	}
	if nodo.izq != nil && cmp(nodo.clave, *desde) > VALOR_CMP {
		nodo.izq.iterar(visitar)
	}
	if !visitar(nodo.clave, nodo.dato) {
		return
	}
	if nodo.der != nil && cmp(nodo.clave, *hasta) < VALOR_CMP {
		nodo.der.iterar(visitar)
	}

}

func (nodo *nodoAb[K, V]) buscarMinimo(pila TDAPila.Pila[*nodoAb[K, V]],
	cmp funcCmp[K], desde *K) *nodoAb[K, V] {
	if nodo == nil {
		return nil
	}

	comparacion := cmp(*desde, nodo.clave)

	//desde es menor a la clave actual
	if comparacion < VALOR_CMP {
		return nodo.izq.buscarMinimo(pila, cmp, desde)
	}

	//desde es mayor o igual a la clave actual
	return nodo

}
