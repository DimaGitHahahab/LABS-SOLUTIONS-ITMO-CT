import java.util.Scanner;

public class TaskA {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        long n = sc.nextLong();
        int count = 0;
        for (long i = 1; i * i < n + 1; i++) {
            if (n % i == 0) {
                count += 2;
                if (i * i != n) {
                    count += 2;
                }
            }
            if (count > 4) {
                break;
            }
        }
        System.out.println(count == 4 ? "True" : "False");
    }
}