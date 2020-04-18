package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// LeerProductos abre el archivo especificado y carga los productos
// en una lista de listas de strings. Si no se puede abrir el archivo,
// el archivo no es json o no existe retorna un error.
func LeerProductos(archivo string) ([][]string, error) {
	f, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	productos := [][]string{}
	if err = json.NewDecoder(f).Decode(&productos); err != nil {
		return nil, err
	}

	return productos, nil
}

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

func main() {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		log.Printf("No se puedo leer archivo de datos: %s", err)
	}

	// Parseo de productos y cálculo de carritos
	tiendas := make(map[string]int)
	precios := make(map[string][]Carrito)

	for _, producto := range productos {
		tienda := producto[0]
		id := producto[1]
		precio, _ := strconv.Atoi(producto[2])

		// recopila productos de cada tienda
		_, ok := tiendas[tienda]
		if ok {
			tiendas[tienda] = tiendas[tienda] + precio
		} else {
			tiendas[tienda] = precio
		}

		// recopila precios de cada producto
		_, ok = precios[id]
		if ok {
			precios[id] = append(precios[id], Carrito{Tienda: tienda, Precio: precio})
		} else {
			precios[id] = []Carrito{Carrito{Tienda: tienda, Precio: precio}}
		}

	}
	var carritos = []Carrito{}
	for tienda, precio := range tiendas {
		carritos = append(carritos, Carrito{Tienda: tienda, Precio: precio})
	}
	fmt.Println("Carritos completado:", carritos, "\n")

	preciosMedios := make(map[string]float64)
	precioMinimo := make(map[string]Carrito)

	for id, listaDePrecios := range precios {
		var total int = 0
		for _, item := range listaDePrecios {
			total += item.Precio

			_, ok := precioMinimo[id]
			if ok {
				if item.Precio < precioMinimo[id].Precio {
					precioMinimo[id] = item
				}
			} else {
				precioMinimo[id] = item
			}
		}
		preciosMedios[id] = float64(total) / float64(len(listaDePrecios))

	}

	fmt.Println("Precios medios:", preciosMedios, "\n")

	fmt.Println("Precios mínimos:", precioMinimo, "\n")

}
