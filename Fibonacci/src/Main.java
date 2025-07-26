//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    interface Func {
        int fib();
    }

    public static Func fibonacci() {
        return new Func() {
            int i = 1;
            int j = 0;
            @Override
            public int fib() {
                int x = j;
                j = i;
                i = x + i;
                return x;
            }
        };
    }

    public static void main(String[] args) {

        Func fib = fibonacci();
        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        // to see how IntelliJ IDEA suggests fixing it.
        System.out.println(fib.fib());
        System.out.println(fib.fib());
        System.out.println(fib.fib());
        System.out.println(fib.fib());
        System.out.println(fib.fib());
        System.out.println(fib.fib());
        System.out.println(fib.fib());
        System.out.println(fib.fib());
    }
}