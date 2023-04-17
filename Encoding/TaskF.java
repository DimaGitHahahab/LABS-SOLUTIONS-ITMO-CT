import java.util.ArrayList;
import java.util.Scanner;

public class TaskF {
    public static void main(String[] args) {
        StringBuilder s = new StringBuilder();
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] a = new int[n];
        for (int i = 0; i < n; i++) {
            a[i] = sc.nextInt();
        }
        String alp = "abcdefghijklmnopqrstuvwxyz";
        int minIndex = 0;
        ArrayList<String> dict = new ArrayList<>();
        for (int i = 0; i < alp.length(); i++) {
            dict.add(String.valueOf(alp.charAt(i)));
            minIndex++;
        }
        s.append(dict.get(a[0]));
        int i = 0;
        while (i < n - 1) {
            i++;
            if (dict.size() > a[i]) {
                s.append(dict.get(a[i]));
                StringBuilder temp = new StringBuilder();
                temp.append(dict.get(a[i - 1]));
                temp.append(dict.get(a[i]).charAt(0));
                dict.add(temp.toString());
                minIndex++;
                temp.setLength(0);
            } else {
                StringBuilder temp = new StringBuilder();
                temp.append(dict.get(a[i - 1]));
                temp.append(dict.get(a[i - 1]).charAt(0));
                dict.add(temp.toString());
                s.append(dict.get(minIndex));
                minIndex++;
                temp.setLength(0);
            }
        }
        System.out.println(s);
    }
}
