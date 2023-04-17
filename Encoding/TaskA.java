import java.util.*;
import java.util.concurrent.ThreadLocalRandom;

public class TaskA {
    private static void quick_sort(long[] arr, int left, int right) {
        if (left >= right) return;
        int index = (int) ThreadLocalRandom.current().nextLong(left, right + 1);
        long q = arr[index];
        int i = left;
        int j = right;
        while (i <= j) {
            while (arr[i] > q) i++;
            while (q > arr[j]) j--;
            if (i <= j) {
                long tempVar = arr[i];
                arr[i] = arr[j];
                arr[j] = tempVar;
                i++;
                j--;
            }
        }
        quick_sort(arr, left, j);
        quick_sort(arr, i, right);
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        long n = sc.nextLong();
        long[] a = new long[(int) n];
        for (int i = 0; i < n; i++) {
            a[i] = sc.nextLong();
        }
        quick_sort(a, 0, a.length - 1);
        long ans = 0;
        for (int i = (int) (n - 1); i > 0; i--) {
            a[i] = a[i] + a[i - 1];
            ans += a[i];
            quick_sort(a, 0, a.length - 1);
        }
        System.out.println(ans);
    }
}