import java.util.Arrays;
import java.util.Scanner;

public class TaskD {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] heap = new int[1];
        int key;
        int size = 0;
        for (int q = 0; q < n; q++) {
            key = sc.nextInt();
            if (key == 0) {
                if (size == heap.length) {
                    heap = Arrays.copyOf(heap, heap.length * 2);
                }
                size++;
                int number = size - 1;
                heap[number] = sc.nextInt();
                while (number > 0 && heap[number] > heap[(number - 1) / 2]) {
                    int tempVar = heap[number];
                    heap[number] = heap[(number - 1) / 2];
                    heap[(number - 1) / 2] = tempVar;
                    number = (number - 1) / 2;
                }
            } else {
                int maximum = heap[0];
                int number = size - 1;
                int tempVar = heap[0];
                heap[0] = heap[number];
                heap[number] = tempVar;
                size--;
                int i = 0;
                while (2 * i + 1 < size) {
                    int j = 2 * i + 1;
                    if (2 * i + 2 < size && (heap[j] < heap[2 * i + 2])) {
                        j = 2 * i + 2;
                    }
                    if (heap[i] < heap[j]) {
                        tempVar = heap[i];
                        heap[i] = heap[j];
                        heap[j] = tempVar;
                        i = j;
                    } else {
                        break;
                    }
                }
                System.out.println(maximum);
            }
        }
    }
}