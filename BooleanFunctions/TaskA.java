import java.util.Scanner;

public class TaskA {

    public static int refl(int[][] a, int n) {
        for (int i = 0; i < n; i++) {
            if (a[i][i] == 0) {
                return 0;
            }
        }
        return 1;
    }

    public static int arefl(int[][] a, int n) {
        for (int i = 0; i < n; i++) {
            if (a[i][i] == 1) {
                return 0;
            }
        }
        return 1;
    }

    public static int sim(int[][] a, int n) {
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (!((a[i][j] == 0) || (a[j][i] == 1))) {
                    return 0;
                }
            }
        }
        return 1;
    }

    public static int asim(int[][] a, int n) {
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (!(!((a[i][j] == 1) && (a[j][i] == 1)) || (i == j))) {
                    return 0;
                }
            }

        }
        return 1;
    }

    public static int tr(int[][] a, int n) {
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                for (int k = 0; k < n; k++) {
                    if (!(!(a[i][j] == 1 && a[j][k] == 1) || (a[i][k] == 1))) {
                        return 0;
                    }
                }
            }
        }
        return 1;
    }


    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[][] matrix1 = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                matrix1[i][j] = sc.nextInt();
            }
        }
        int[][] matrix2 = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                matrix2[i][j] = sc.nextInt();
            }
        }


        System.out.println(refl(matrix1, n) + " " + arefl(matrix1, n) + " " + sim(matrix1, n) + " " + asim(matrix1, n) + " " + tr(matrix1, n));
        System.out.println(refl(matrix2, n) + " " + arefl(matrix2, n) + " " + sim(matrix2, n) + " " + asim(matrix2, n) + " " + tr(matrix2, n));
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                int composition = 0;
                for (int k = 0; k < n; k++) {
                    if (matrix1[i][k] == 1 && matrix2[k][j] == 1) {
                        composition = 1;
                        break;
                    }
                }
                System.out.print(composition + " ");

            }
            System.out.println();

        }
    }
}
