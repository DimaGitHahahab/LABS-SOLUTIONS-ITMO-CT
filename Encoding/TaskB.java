import java.util.Arrays;
import java.util.Scanner;

public class TaskB {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.next();
        String[] cycles = new String[s.length()];
        cycles[0] = s;
        for (int i = 1; i < s.length(); i++) {
            StringBuilder tempBuilder = new StringBuilder();
            tempBuilder.append(s, s.length() - i, s.length());
            tempBuilder.append(s, 0, s.length() - i);
            cycles[i] = tempBuilder.toString();
        }
        Arrays.sort(cycles);
        StringBuilder ans = new StringBuilder();
        for (int i = 0; i < s.length(); i++) {
            ans.append(cycles[i].charAt(s.length() - 1));
        }
        System.out.println(ans);
    }
}
