package tp1

import (
	"log"
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

// Productos es una lista de productos donde para cada producto
// se sabe el nombre del super mercado, el id y su precio.
// Esta estructura se puede cargar usando la funcion LeerProductos
// que carga informacion guardada en `productos.json`.
type Productos [][]string

// Carrito contiene el nombre de la tienda y el precio final luego
// de sumar todos los productos.
type Carrito struct {
	Tienda string
	Precio int
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {

	tiendas := make(map[string]int)

	for _, producto := range ids {
		log.Println(producto)
		/*tienda := producto
		precio := producto
		value, ok := tiendas[tienda]
		if ok {
			tiendas[tienda] += precio
		} else {
			tiendas[tienda] = precio
		}
		*/
	}
	var carritos = []Carrito{}
	for tienda, precio := range tiendas {
		carritos = append(carritos, Carrito{Tienda: tienda, Precio: precio})
	}
	log.Println(carritos)
	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	return 0
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	return nil, false
}
