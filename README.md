# Guía: Árboles Binarios de Búsqueda

## Ejercicios

1. Escribir una función que encuentre el segundo elemento mas grande en un ABB. Ejemplo:

    ```output
    Input: Root of below ABB

        10
       /
      5

    Output:  5
    ```

    ```output
    Input: Root of below ABB

        10
       /  \
      5    20
            \ 
             30 

    Output:  20
    ```

2. Escribir una funcion que dado un ABB, encuentre el predecesor inorder de una clave dada. Si la clave no se encuentra en el ABB, devuelve el nodo más grande anterior (si lo hay) presente en la ABB.

    ```output
             15
            /  \
           /    \
         10      20
        /  \    /  \
       /    |  |    \
      8    12  16   25

    El predecesor de 8 no existe. 
    El predecesor de 10 es 8 
    El predecesor de 12 es 10 
    El predecesor de 20 es 16  
    ```

3. Escribir una funcion para comprobar si un árbol binario es ABB o no. Devuelve un booleano.

4. Implementar un `TreeSet[T constraints.Ordered]` a partir de un ABB.

5. Implementar un método de ABB que devuelva la cantidad de nodos en el árbol.

6. Implementar un iterador _Inorden_ para un ABB.

7. Implementar un iterador _Preorden_ para un ABB.

8. Implementar un iterador _Postorden_ para un ABB.

9. Implementar un iterador por niveles para un ABB.

    Tal que primero devuelve la raíz, luego los hijos de la raíz que se encuentran en el nivel 1, luego los hijos de los nodos del nivel 1 que se encuentran en el nivel 2 y así sucesivamente.

10. (a) Dada la siguiente lista de números realizar lo que se pide en cada ítem devolviendo el dibujo del árbol resultante:

    > 4 19 -7 49 100 0 22 12

    (b) Construir el ABB que resulta de insertar los números en el orden en que aparecen:

    - Tomar el árbol anterior e insertar elemento 10.
    - Tomar el árbol anterior y eliminar el elemento 49.
    - Tomar el árbol anterior e insertar el elemento 1.

11. (a) Dada la siguiente lista de números realizar lo que se pide en cada ítem devolviendo el dibujo del árbol resultante:

    > 23 7 41 26 32 52 11 5 56

    (b) Construir el ABB que resulta de insertar los números en el orden en que aparecen:

    - Tomar el árbol anterior y eliminar el elemento 41.
    - Tomar el árbol anterior e insertar el elemento 9.
    - Tomar el árbol anterior e insertar el elemento 28.
