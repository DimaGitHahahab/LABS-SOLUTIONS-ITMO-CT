import java.util.ArrayList;
import java.util.Scanner;

public class TaskE {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.nextLine();
        int[] stack = new int[s.length()];
        int stackIndex = -1;
        StringBuilder temp = new StringBuilder();
        ArrayList<String> buffer = new ArrayList<>();
        for (int i = 0; i < s.length(); i++) {
            char symbol = s.charAt(i);
            if (symbol == ' ') {
                buffer.add(temp.toString());
                temp.setLength(0);
            } else {
                temp.append(symbol);
            }
        }
        buffer.add(temp.toString());
        for (String symbol : buffer) {
            if (!(symbol.equals("*") || symbol.equals("+") || symbol.equals("-"))) {
                int number = Integer.parseInt(symbol);
                stackIndex++;
                stack[stackIndex] = number;
            } else {
                int firstNumber = stack[stackIndex];
                stackIndex--;
                int secondNumber = stack[stackIndex];
                int result = 0;
                switch (symbol) {
                    case "*" -> result = firstNumber * secondNumber;
                    case "+" -> result = firstNumber + secondNumber;
                    case "-" -> result = secondNumber - firstNumber;
                    default -> System.out.println("Error");
                }
                stack[stackIndex] = result;
            }
        }
        System.out.println(stack[0]);
    }
}
