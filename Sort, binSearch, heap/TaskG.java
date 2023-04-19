import java.util.Scanner;

public class TaskG {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt() - 1;
        int x = sc.nextInt();
        int y = sc.nextInt();
        int minV = x;
        if (y < x) {
            minV = y;
        }
        long time = minV;
        int left = -1;
        int right = Integer.MAX_VALUE;
        while (right > left + 1) {
            int mid = (right + left) / 2;
            if (mid / x + mid / y < n) {
                left = mid;
            } else {
                right = mid;
            }
        }
        time += right;
        System.out.println(time);
    }
}