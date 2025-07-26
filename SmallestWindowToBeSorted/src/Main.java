import java.util.Arrays;

public class Main {
    // O(n*log(n)) due to sort()
    public static int[] smallestWindowToBeSorted(int[] input) {
        int[] sorted = new int[input.length];
        System.arraycopy(input, 0, sorted, 0, input.length);
        Arrays.sort(sorted);
        int lower = -1;
        int higher = -1;
        for (int i = 0; i < input.length; i++) {
            if (lower == -1 && input[i] != sorted[i]) {
                lower = i;
            } else if (lower >= 0 && input[i] != sorted[i]) {
                higher = i;
            }
        }
        return new int[] { lower, higher };
    }

    public static int[] smallestWindowToBeSorted2(int[] array) {
        int left = -1;
        int right = -1;
        int max_seen = Integer.MIN_VALUE;
        int min_seen = Integer.MAX_VALUE;
        for (int i = 0; i < array.length; i++) {
            max_seen = Math.max(max_seen, array[i]);
            if (array[i] < max_seen) {
                right = i;
            }
        }
        for (int i = array.length - 1; i >= 0; i--) {
            min_seen = Math.min(min_seen, array[i]);
            if (array[i] > min_seen) {
                left = i;
            }
        }
        return new int[] { left, right };
    }

    public static void main(String[] args) {
        int[] input1 = { 3, 7, 5, 6, 9 };
        System.out.println(Arrays.toString(smallestWindowToBeSorted2(input1)));

        int[] input2 = { 3, 7, 5, 6, 9, 10, 7, 7 };
        System.out.println(Arrays.toString(smallestWindowToBeSorted2(input2)));

        int[] input3 = { 3, 7, 9, 10, 8, 8, 11, 8, 8 };
        System.out.println(Arrays.toString(smallestWindowToBeSorted2(input3)));
    }
}