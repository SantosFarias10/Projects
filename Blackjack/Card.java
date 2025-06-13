public class Card {
    private String suit;    // Palo: Corazon, Diamante, Trebol, Pica.
    private String rank;    // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A.
    private int value;      // J=10, Q=10, K=10, A=11.

    public Card(String suit, String rank, int value){
        this.suit = suit;
        this.rank = rank;
        this.value = value;
    }

    // Metodo para verificar si son del mismo numero
    public boolean equalsRank(card c1, card c2) {
        return c1.getRank().equals(c2.getRank());
    }

    // Getters y setters
    public String getSuit() {
        return suit;
    }

    public String getRank() {
        return rank;
    }

    public int getValue() {
        return value;
    }

    public void setSuit(String suit) {
        this.suit = suit;
    }

    public void setRank(String rank) {
        this.rank = rank;
    }

    public void setValue(int value) {
        this.value = value;
    }
    
    // metodo para imprimir la cards
    public String toString() {
        return rank + " de " + suit;
    }
}