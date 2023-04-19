import java.util.ArrayList;
import java.util.Scanner;

public class TaskF {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int k = sc.nextInt();
        int[] arr = new int[n];
        int[] req = new int[k];
        for (int i = 0; i < n; i++) {
            arr[i] = sc.nextInt();
        }
        for (int i = 0; i < k; i++) {
            req[i] = sc.nextInt();
        }

        for (int i = 0; i < k; i++) {
            int left = -1;
            int right = n;
            while (right > left + 1) {
                int mid = (left + right) / 2;
                if (arr[mid] >= req[i]) {
                    right = mid;
                } else {
                    left = mid;
                }
            }
            ArrayList<Integer> candidates = new ArrayList<>();
            if (left > -1) {
                candidates.add(arr[left]);
            }
            if (right < n) {
                candidates.add(arr[right]);
            }
            int nearest = candidates.get(0);
            int dif = Math.abs(req[i] - nearest);
            for (int f : candidates) {
                int abs = Math.abs(f - req[i]);
                if (abs < dif) {
                    dif = abs;
                    nearest = f;
                }
            }
            System.out.println(nearest);
        }
    }
}
 