import java.util.Arrays;
import java.util.Scanner;

public class TaskC {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.next();
        int n = s.length();
        StringBuilder[] t = new StringBuilder[n];
        for (int i = 0; i < n; i++) {
            t[i] = new StringBuilder();
        }
        for (int k = 0; k < n; k++) {
            for (int i = 0; i < n; i++) {
                t[i].insert(0, s.charAt(i)+ "");
            }
            Arrays.sort(t);
        }

        System.out.println(t[0]);
    }
}
