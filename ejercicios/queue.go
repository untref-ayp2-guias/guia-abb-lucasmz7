// Package queue implementa una cola genérica sobre un arreglo dinámico.
// Expone la estructura Queue y sus métodos para manipular una cola.
package ejercicios

import "errors"

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue[T any] struct {
	data []T
}

// NewQueue crea una nueva cola vacía. O(1)
//
// Uso:
//
//	q := queue.New[int]() // Crea una nueva cola de enteros.
func NewQueue[T any]() *Queue[T] {
	var queue = Queue[T]{
		data: []T{},
	}
	return &queue
}

// Enqueue agrega un elemento a la cola. O(1)
//
// Uso:
//
//	q.Enqueue(10) // Agrega el entero 10 a la cola.
//
// Parámetros:
//   - `v`: el elemento a agregar a la cola.
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue saca un elemento de la cola. O(1)
//
// Uso:
//
//	v, err := q.Dequeue() // Saca un elemento de la cola.
//
// Retorna:
//   - el elemento sacado de la cola.
//   - un error si la cola está vacía.
func (q *Queue[T]) Dequeue() (T, error) {
	var x T

	if len(q.data) == 0 {
		return x, errors.New("cola vacía")
	}

	x = q.data[0]
	q.data = q.data[1:]

	return x, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
// Esta operación no modifica la cola.
//
// Uso:
//
//	v, err := q.Front() // Obtiene el elemento del frente de la cola.
//
// Retorna:
//   - el elemento del frente de la cola.
//   - un error si la cola está vacía.
func (q *Queue[T]) Front() (T, error) {
	var x T

	if len(q.data) == 0 {
		return x, errors.New("cola vacía")
	}

	x = q.data[0]

	return x, nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
//
// Uso:
//
//	empty := q.IsEmpty() // Verifica si la cola está vacía.
//
// Retorna:
//   - `true` si la cola está vacía; `false` en caso contrario.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Length() int {
	return len(q.data)
}
