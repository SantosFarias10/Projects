/*
 * Escribe un programa que imprima los 50 primeros números de la sucesión
 * de Fibonacci empezando en 0.
 * - La serie Fibonacci se compone por una sucesión de números en
 *   la que el siguiente siempre es la suma de los dos anteriores.
 *   0, 1, 1, 2, 3, 5, 8, 13...
 */

#include <iostream>

int main(){
  unsigned int n1 = 1;
  unsigned int n2 = 1;
  unsigned int next;

  for(int i = 0; i < 50; ++i){
    std::cout << n1 << "\n";

    next = n1+n2;

    n1 = n2;
    n2 = next;
  }

  return 0;
}
