#include <iostream>

/*
 * Escribe un programa que muestre por consola (con un print) los
 * números de 1 a 100 (ambos incluidos y con un salto de línea entre
 * cada impresión), sustituyendo los siguientes:
 * - Múltiplos de 3 por la palabra "fizz".
 * - Múltiplos de 5 por la palabra "buzz".
 * - Múltiplos de 3 y de 5 a la vez por la palabra "fizzbuzz".
*/

int main() {
  for(int i = 1; i <= 100; i++){
    if(i % 3 == 0 && i % 5 == 0){
      std::cout << i << " Fizz Buzz\n";
    } else if(i % 3 == 0){
      std::cout << i << " Fizz\n";
    } else if(i % 5 == 0){
      std::cout << i << " Buzz\n";
    } else {
      std::cout << i << "\n";
    }
  }

  return 0;
}
