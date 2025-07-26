import java.util.Arrays;
import java.util.HashMap;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void printZigZagForm(String s, int k) {
        int []map = new int[s.length()];
        int line = 0, delta = 1;
        for (int i = 0; i < s.length(); i++) {
            map[i] = line;
            line += delta;
            if (line == 0 || line == k - 1) {
                delta = -delta;
            }
        }
        for (int i = 0; i < k; i++) {
            StringBuilder buf = new StringBuilder();
            for (int j = 0; j < s.length(); j++) {
                if (map[j] == i) {
                    buf.append(s.charAt(j));
                } else {
                    buf.append(' ');
                }
            }
            System.out.println(buf);
        }

    }

    public static void main(String[] args) {

        printZigZagForm("thisisazigzag", 3);
        printZigZagForm("thisisanotherazigzag", 5);
    }
}