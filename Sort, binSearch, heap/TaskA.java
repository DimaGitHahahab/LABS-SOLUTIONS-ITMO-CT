import java.util.concurrent.ThreadLocalRandom;
import java.util.Arrays;
import java.util.Scanner;

public class TaskA {

    public static void quick_sort(int[] arr, int left, int right) {
        if (left >= right) return;
        int index = ThreadLocalRandom.current().nextInt(left, right + 1);
        int q = arr[index];
        int i = left;
        int j = right;
        while (i <= j) {
            while (arr[i] < q) i++;
            while (q < arr[j]) j--;
            if (i <= j) {
                int tempVar = arr[i];
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
        int n = sc.nextInt();
        int[] a = new int[n];
        for (int i = 0; i < n; i++) {
            a[i] = sc.nextInt();
        }
        quick_sort(a, 0, a.length - 1);
        for (int j : a) {
            System.out.print(j + " ");
        }
    }
}