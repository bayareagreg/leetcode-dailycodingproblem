import java.util.Arrays;

public class Main {
    // brute force O(n^2)
    public static int[] numberOfSmallerElementsToTheRight(int[] array) {
        int[] result = new int[array.length];
        for (int i = 0; i < array.length; i++) {
            for (int j = i + 1; j < array.length; j++) {
                if (array[i] > array[j]) {
                    result[i]++;
                }
            }
        }
        return result;
    }

    public static int insertSorted(int[] arr, int length, int element) {
        int index = findInsertionIndex(arr, length, element);
        arr[index] = element;
        System.arraycopy(arr, index, arr, index + 1, length - index);
        return index;
    }

    private static int findInsertionIndex(int[] arr, int length, int element) {
        int low = 0, high = length - 1;

        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (arr[mid] == element) {
                return mid;
            } else if (arr[mid] < element) {
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }
        return low;
    }

    // O(n*log(n))
    public static int[] numberOfSmallerElementsToTheRight2(int[] array) {
        int[] result = new int[array.length];
        int[] sorted = new int[array.length];
        for (int i = array.length - 1; i >= 0; i--) {
            result[i] = insertSorted(sorted, array.length - i - 1, array[i]);
        }
        return result;
    }

    public static void main(String[] args) {
        int[] input1 = new int[] {3, 4, 9, 6, 1};
        System.out.println(Arrays.toString(numberOfSmallerElementsToTheRight2(input1)));
    }
}