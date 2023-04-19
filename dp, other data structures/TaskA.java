import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class TaskA {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer tokenizer = new StringTokenizer(reader.readLine());
        int n = Integer.parseInt(tokenizer.nextToken());
        int[] stack = new int[n];
        int[] minStack = new int[n];
        int index = -1;
        int minIndex = -1;
        for (int i = 0; i < n; i++) {
            tokenizer = new StringTokenizer(reader.readLine());
            int command = Integer.parseInt(tokenizer.nextToken());
            if (command == 1) {
                int value = Integer.parseInt(tokenizer.nextToken());
                index++;
                stack[index] = value;
                if (minIndex == -1) {
                    minIndex++;
                    minStack[minIndex] = value;
                } else {
                    if (value <= minStack[minIndex]) {
                        minIndex++;
                        minStack[minIndex] = value;
                    }
                }
            } else if (command == 2) {
                if (stack[index] == minStack[minIndex]) {
                    minIndex--;
                }
                index--;
            } else if (command == 3) {
                System.out.println(minStack[minIndex]);
            } else {
                System.out.println("Error");
            }
        }
    }
}