import java.util.Arrays;
import java.util.Scanner;

public class TaskB {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int leftBorder = sc.nextInt();
        int rightBorder = sc.nextInt();
        boolean[] arr = new boolean[rightBorder + 1];
        Arrays.fill(arr, true);
        arr[0] = false;
        arr[1] = false;
        for (int i = 2; i < arr.length; i++) {
            if (arr[i]) {
                for (int j = 2; j < rightBorder / i + 1; j++) {
                    arr[i * j] = false;
                }
            }
        }
        StringBuilder builder = new StringBuilder();
        for (int i = leftBorder; i < arr.length; i++) {
            if (arr[i]) {
                builder.append(i).append(" ");
            }
        }
        System.out.println(builder);
    }
}