import java.utils.*;
import java.util.Collections;

public class Deck {
    private List<Card> cards;

    public Deck() {
        cards = new ArrayList<>();
        String[] suits = {"Corazones", "Picas", "Treboles", "Diamantes"};
        String[] ranks = {"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", 
        "Q", "K", "A"};
        int[] values = {2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11};

        for (String suit : suits) {
            for (int i = 0; i < ranks.length; i++) {
                cards.add(new Card(suit, ranks[i], values[i]));
            }
        }

        // Mezclamos
        shuffle();
    }

    public void shuffle() {
        Collections.shuffle(cards);     // metodo que mezcla aleatoriamente los elementos de una lista
    }

    public int size() {
        return cards.size();
    }

    // Retorna la primera carta del mazo y la elimina
    public Card distribute() {
        if (cards.isEmpty()) {
            System.out.println("No hay mas cartas");
            return null;    //¿¿??
        }

        return cards.remove(0);     // retornamos la primera carta de la lista y luego la elimina.
    }
}