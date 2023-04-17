import java.util.*;

public class TaskB {
    static boolean F0(String s) {
        return s.charAt(0) == '0';
    }

    static boolean F1(String s) {
        return s.charAt(s.length() - 1) == '1';
    }

    static boolean Fs(String s) {
        if (s.length() == 1) {
            return false;
        }
        for (int i = 0; (i << 1) < s.length(); i++) {
            if (s.charAt(i) == s.charAt(s.length() - i - 1)) {
                return false;
            }
        }
        return true;
    }

    static boolean Fm(String s) {
        for (int i = 0; i < s.length(); i++) {
            for (int j = 0; j < s.length(); j++) {
                if ((i & j) == j && s.charAt(i) < s.charAt(j)) {
                    return false;
                }
            }
        }
        return true;
    }

    static boolean Fl(String s) {
        for (int i = 0; i < s.length() * 2; i++) {
            boolean ok = true;
            for (int j = 0; j < s.length(); j++) {
                int x = (j * 2 + 1) & i, f = '0';
                while (x != 0) {
                    f ^= x & 1;
                    x >>= 1;
                }
                if (f != s.charAt(j)) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return true;
            }
        }
        return false;
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        boolean allF0 = true, allF1 = true, allFm = true, allFs = true, allFl = true;
        for (int i = 0; i < n; i++) {
            in.next();
            String f = in.next();
            allF0 = allF0 && F0(f);
            allF1 = allF1 && F1(f);
            allFm = allFm && Fm(f);
            allFs = allFs && Fs(f);
            allFl = allFl && Fl(f);
        }
        System.out.print(allF0 || allF1 || allFm || allFs || allFl ? "NO" : "YES");
    }
}
