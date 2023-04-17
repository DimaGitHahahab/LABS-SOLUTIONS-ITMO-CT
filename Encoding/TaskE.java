import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.Scanner;

public class TaskE {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.next();
        String alp = "abcdefghijklmnopqrstuvwxyz";
        ArrayList<Integer> output = new ArrayList<>();
        LinkedHashMap<String, Integer> dict = new LinkedHashMap<>();
        for (int i = 0; i < 26; i++) {
            dict.putIfAbsent(String.valueOf(alp.charAt(i)), i);
        }
        int minIndex = 26;
        StringBuilder buffer = new StringBuilder();
        //buffer.append(s.charAt(0));
        for (int i = 0; i < s.length(); i++) {
            String symbol = String.valueOf(s.charAt(i));
            if (dict.get(buffer + symbol) != null) {
                buffer.append(symbol);
            } else {
                output.add(dict.get(buffer.toString()));
                buffer.append(symbol);
                dict.put(buffer.toString(), minIndex);
                minIndex++;
                buffer.setLength(0);
                buffer.append(symbol);
            }

        }
        output.add(dict.get(buffer.toString()));
        for (int i : output) {
            System.out.print(i + " ");

        }
    }
}
