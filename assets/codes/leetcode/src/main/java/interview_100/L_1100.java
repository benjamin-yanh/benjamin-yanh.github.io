package interview_100;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

public class L_1100 {
    public int numKLenSubstrNoRepeats(String s, int k) {
        int res = 0;
        char[] list = s.toCharArray();
        LinkedList<Character> strings = new LinkedList<>();
        for (char c : list) {
            while (strings.contains(c)) {
                strings.removeFirst();
            }
            while (strings.size() > k) {
                strings.removeFirst();
            }
            strings.add(c);
            if (strings.size() == k) {
                res++;
            }
        }
        return res;
    }

    public static void main(String[] args) {

    }
}
