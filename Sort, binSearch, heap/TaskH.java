import java.util.Scanner;

public class TaskH {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        double c = sc.nextFloat();
        double left = 0;
        double right = 20000000;
        while (right > left + 0.000001) {
            double x = (left + right) / 2;
            double eq = Math.pow(x, 2) + Math.sqrt(x);
            if (eq >= c) {
                right = x;
            } else {
                left = x;
            }
        }
        System.out.println(right);
    }
}