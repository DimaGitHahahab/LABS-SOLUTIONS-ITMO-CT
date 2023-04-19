import java.util.Scanner;

public class TaskB {
    public static void countingSort(int[] arr) {
        int[] D = new int[101];
        for (int j : arr) {
            D[j]++;
        }
        int j = 0;
        for (int i = 0; i < 101; i++) {
            while (D[i] > 0) {
                arr[j] = i;
                D[i]--;
                j++;
            }
        }
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] a = new int[n];
        for (int i = 0; i < n; i++) {
            a[i] = sc.nextInt();
        }
        countingSort(a);
        for (int j : a) {
            System.out.print(j + " ");
        }
    }
}