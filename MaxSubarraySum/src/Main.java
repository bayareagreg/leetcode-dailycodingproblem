import java.util.Arrays;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static int sum(int[] array, int lower, int higher) {
        int sum = 0;
        for (int i = lower; i <= higher; i++) {
            sum += array[i];
        }
        return sum;
    }

    public static int maxSubarraySum(int[] array) {
        int current_max = 0;
        for (int i = 0; i < array.length - 1; i++) {
            for (int j = i; j < array.length; j++) {
                current_max = Math.max(current_max, sum(array, i, j));
            }
        }
        return current_max;
    }

    // brute force O(n^2)
    public static int maxSubarraySum2(int[] array) {
        int res = array[0];
        for (int i = 0; i < array.length; i++) {
            int curSum = 0;
            for (int j = i; j < array.length; j++) {
                curSum += array[j];
                res = Math.max(res, curSum);
            }
        }
        return res;
    }

    // kadane O(n)
    public static int maxSubarraySum3(int[] array) {
        int max_sum = array[0];
        int cur_sum = 0;
        for (int i = 0; i < array.length; i++) {
            //cur_sum = Math.max(cur_sum, 0);
            //cur_sum += array[i];
            cur_sum = Math.max(cur_sum + array[i], array[i]);
            max_sum = Math.max(max_sum, cur_sum);
        }
        return max_sum;
    }


    public static void main(String[] args) {
        /*
        int[] input1 = { 34, -50, 42, 14, -5, 86 };
        System.out.println(maxSubarraySum2(input1));

        int[] input2 = { -1, -2, -4 };
        System.out.println(maxSubarraySum2(input2));

        int[] input3 = {2, 3, -8, 7, -1, 2, 3};
        System.out.println(maxSubarraySum2(input3));
         */

        int[] input4 = {-50, -2, -4};
        System.out.println(maxSubarraySum3(input4));

        /*
        int[] input5 = {5, 4, 1, 7, 8};
        System.out.println(maxSubarraySum2(input5));
         */

        int[] input6 = {-2, -1, -3, 4, -1, 2, 1, -5, 4};
        System.out.println(maxSubarraySum3(input6));

    }
}