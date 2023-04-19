import java.util.Arrays;
import java.util.Scanner;

public class TaskC {

    public static long count = 0;

    public static void merge(int[] a, int left, int mid, int right) {

        int it1 = 0;
        int it2 = 0;
        int[] result = new int[right - left];
        while (left + it1 < mid && mid + it2 < right) {
            if (a[left + it1] < a[mid + it2]) {
                result[it1 + it2] = a[left + it1];
                it1 += 1;
            } else {
                result[it1 + it2] = a[mid + it2];
                it2 += 1;
                count += (mid - left) - it1;
            }
        }
        while (left + it1 < mid) {
            result[it1 + it2] = a[left + it1];
            it1 += 1;
        }
        while (mid + it2 < right) {
            result[it1 + it2] = a[mid + it2];
            it2 += 1;
        }
        if (it1 + it2 >= 0) System.arraycopy(result, 0, a, left, it1 + it2);
    }

    public static void mergeSort(int[] a, int left, int right) {
        if (left + 1 >= right) {
            return;
        }
        int mid = (left + right) / 2;
        mergeSort(a, left, mid);
        mergeSort(a, mid, right);
        merge(a, left, mid, right);
    }


    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] a = new int[n];
        for (int i = 0; i < n; i++) {
            a[i] = sc.nextInt();
        }
        mergeSort(a, 0, a.length);
        System.out.println(count);
    }
}