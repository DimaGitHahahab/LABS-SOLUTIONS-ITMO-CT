import java.util.*;

public class TaskH {
    static String nand(String a, String b) {
        return "(" + a + "|" + b + ")";
    }

    static String or(String a, String b) {
        return nand(nand(a, a), nand(b, b));
    }

    static String solve(int n) {
        String a = "A" + (n - 1);
        String b = "B" + (n - 1);
        if (n == 1) {
            return nand(nand(a, b), nand(a, b));
        }
        return nand(nand(solve(n - 1), or(a, b)), nand(a, b));
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        System.out.print(solve(n));
    }
}
