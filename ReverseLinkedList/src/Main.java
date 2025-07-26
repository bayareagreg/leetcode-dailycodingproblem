//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    static class Node {
        int data;
        Node next;

        Node(int data, Node next) {
            this.data = data;
            this.next = next;
        }
    }

    static Node reverseLinkedList(Node node) {
        Node next = node.next;
        if (next == null) {
            return node;
        }
        Node next2 = next.next;
        node.next = null;
        do {
            next.next = node;
            if (next2 == null) {
                return next;
            }
            node = next;
            next = next2;
            next2 = next2.next;
        } while (true);
    }

    static void printLinkedList(Node node) {
        while (node != null) {
            System.out.println(node.data);
            node = node.next;
        }
    }

    public static void main(String[] args) {
        Node n4 = new Node(4, null);
        Node n3 = new Node(3, n4);
        Node n2 = new Node(2, n3);
        Node n1 = new Node(1, n2);

        printLinkedList(n1);
        n1 = reverseLinkedList(n1);
        System.out.println("Reversed:");
        printLinkedList(n1);
    }
}