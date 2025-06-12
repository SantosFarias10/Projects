import java.util.Scanner;
import java.util.Random;

public class Adivina {
  public static void main(String[] args) {
    
    Random random = new Random();
    Scanner scanner = new Scanner(System.in);   // Scanner toma parametro System.in, que es un InputStream que representa la entrada estandar del sistema, basicamente lee cualquier flujo de entrada.

    int numeroAleatorio = random.nextInt(100) + 1;
    int intentos = 1;
    int numeroUser;
    int diferenciaActual;
    int diferenciaAnterior = Integer.MAX_VALUE;
    
    System.out.println("Bienvenido al juego de adivina el numero");
    System.out.println("Intenta adivinar el numero entre 1 y 100");
    System.out.println("Tienes 10 intentos");
  
    while(intentos <= 10) {
      System.out.println("------------------------------");
      System.out.println("Intento " + intentos + " de 10");
      System.out.println("Introduce un numero: ");

      numeroUser = scanner.nextInt();

      diferenciaActual = Math.abs(numeroUser - numeroAleatorio);

      if(numeroUser == numeroAleatorio) {
        System.out.println("Felicitaciones!!! Has adivinado el numero");
        break;
      } else {
        if (intentos == 1) {
          System.out.println("Primera pista:");
          if (numeroUser > numeroAleatorio) {
            System.out.println("Tu numero es mas grande que el numero escondido");
          } else if (numeroUser < numeroAleatorio) {
            System.out.println("Tu numero es mas chico que el numero escondido");
          }
        } else { 
          if (diferenciaActual < diferenciaAnterior) {
            System.out.println("🔥Caliente! Te estás acercando.");
          } else if (diferenciaActual > diferenciaAnterior) {
            System.out.println("🧊 Frío... te estás alejando.");
          } else {
            System.out.println("Igual de lejos que antes");
          }
        }

        diferenciaAnterior = Math.min(diferenciaActual, diferenciaAnterior);
      }

      intentos++;
    }

    if (intentos > 10) {
      System.out.println("💀 Se te acabaron los intentos. El número era: " + numeroAleatorio);
    }

    scanner.close();
  }
}
