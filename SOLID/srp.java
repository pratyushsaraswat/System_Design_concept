
import java.util.*;

class ShoppingCart {
    private List<Double> prices = new ArrayList<>();

    public void addProduct(double price) {
        prices.add(price);
    }

    public double calculateTotal() {
        double total = 0;

        for (double price : prices) {
            total += price;
        }

        return total;
    }

    public void saveToDatabase() {
        System.out.println("Saving order to database");
    }

    public void sendEmail() {
        System.out.println("Sending confirmation email");
    }
}
