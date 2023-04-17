import java.util.*;

public class TaskF {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int k = in.nextInt();
        int[][] a = new int[k][n];
        for (int i = 0; i < k; i++) {
            for (int j = 0; j < n; j++) {
                a[i][j] = in.nextInt();
            }
        }
        int[] x = new int[n];
        Arrays.fill(x, -1);
        for (int cnt = 0; cnt < n; ++cnt) {
            for (int i = 0; i < k; ++i) {
                int pos = -1;
                boolean f = false;
                for (int j = 0; j < n; j++) {
                    if (a[i][j] != -1) {
                        if (x[j] == -1) {
                            pos = pos == -1 ? j : n;
                        } else {
                            f = f || a[i][j] == x[j];
                        }
                    }
                }
                if (!f) {
                    if (pos == -1) {
                        System.out.print("YES");
                        return;
                    } else if (pos != n) {
                        x[pos] = a[i][pos];
                    }
                }
            }
        }
        System.out.print("NO");
    }
}
