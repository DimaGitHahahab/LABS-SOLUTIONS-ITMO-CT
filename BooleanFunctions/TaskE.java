import java.util.Scanner;

public class TaskE {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        String[] table = new String[(int) Math.pow(2, n)];
        int[] f = new int[(int) Math.pow(2, n)];
        int[] outF = new int[(int) Math.pow(2, n)];
        for (int i = 0; i < Math.pow(2, n); i++) {
            table[i] = sc.next();
            f[i] = sc.nextInt();
        }
        for (int i = 0; i < Math.pow(2, n); i++) {
            outF[i] = f[0];
            for (int j = 0; j < Math.pow(2, n) - i - 1; j++) {
                f[j] = (f[j] + f[j + 1]) % 2;

            }
        }
        for (int i = 0; i < Math.pow(2, n); i++) {
            System.out.print(table[i] + " " + outF[i]);
            System.out.println();
        }

    }
}
