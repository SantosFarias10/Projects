import java.util.Scanner;
import java.util.Random;

public class Adivina {
  public static void main(String[] args) {
    
    Random random = new Random();
    Scanner scanner = new Scanner(System.in);   // Scanner toma parametro System.in, que es un InputStream que representa la entrada estandar del sistema, basicamente lee cualquier flujo de entrada.

    int numeroAleatorio = random.nextInt(100) + 1;
    int intentos = 1;
    int numeroUser;
    int diferencia;
    
    System.out.println("Bienvenido al juego de adivina el numero");
    System.out.println("Intenta adivinar el numero entre 1 y 100");
    System.out.println("Tienes 10 intentos");
  
    while(intentos <= 10) {
        System.out.println("------------------------------");
      System.out.println("Intento " + intentos + " de 10");
      System.out.println("Introduce un numero: ");

      numeroUser = scanner.nextInt();

      diferencia = Math.abs(numeroUser - numeroAleatorio);

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
          } else {
            System.out.println("Felicitaciones!!! Has adivinado el numero");
          }
        } else { 
          if (5 <= diferencia && diferencia <= 10) {
            System.out.println("🔥Caliente! Te estás acercando.");
          } else if (10 < diferencia && diferencia <= 20) {
            System.out.println("Tibio");
          } else {
            System.out.println("🧊 Frío...");
          }
        }

      }

      intentos++;
    }

    if (intentos > 10) {
      System.out.println("💀 Se te acabaron los intentos. El número era: " + numeroAleatorio);
    }

    scanner.close();
  }
}
