import java.util.Arrays;
public class Main {
    // O(n^2)
    public static int[] productOfAllOtherElements(int[] input) {
        int[] output = new int[input.length];
        for (int i = 0; i < input.length; i++) {
            int product = 0;
            for (int j = 0; j < input.length; j++) {
                if (i != j) {
                    product = product == 0 ? input[j] : product * input[j];
                }
            }
            output[i] = product;
        }
        return output;
    }

    // O(n)
    public static int[] productOfAllOtherElements2(int[] input) {
        int[] output = new int[input.length];
        int total = 0;
        for (int i = 0; i < input.length; i++) {
            total = total == 0 ? input[i] : total * input[i];
        }

        for (int i = 0; i < input.length; i++) {
            output[i] = total / input[i];
        }

        return output;
    }

    // can't use division, O(n)
    public static int[] productOfAllOtherElements3(int[] input) {
        int[] pre = new int[input.length];
        int[] post = new int[input.length];
        int[] output = new int[input.length];

        for (int i = 0; i < input.length; i++) {
            if (i == 0) {
                pre[i] = input[i];
            } else {
                pre[i] = pre[i-1] * input[i];
            }
        }

        for (int i = input.length - 1; i >= 0; i--) {
            if (i == input.length - 1) {
                post[i] = input[i];
            } else {
                post[i] = post[i+1] * input[i];
            }
        }

        for (int i = 0; i < input.length; i++) {
            if (i == 0) {
                output[i] = post[i+1];
            } else if (i == input.length - 1) {
                output[i] = pre[i-1];
            } else {
                output[i] = pre[i - 1] * post[i + 1];
            }
        }

        return output;
    }

    public static void main(String[] args) {
        int[] input1 = { 1, 2, 3, 4, 5 };
        int[] output1 = productOfAllOtherElements3(input1);
        System.out.println(Arrays.toString(output1));

        int[] input2 = { 3, 2, 1 };
        int[] output2 = productOfAllOtherElements3(input2);
        System.out.println(Arrays.toString(output2));

        int[] input3 = { 3 };
        // undefined?
        //int[] output3 = productOfAllOtherElements3(input3);
        //System.out.println(Arrays.toString(output3));
    }
}