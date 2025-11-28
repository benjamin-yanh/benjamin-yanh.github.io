package interview_100;

import java.util.*;

public class L_1055 {
    public int shortestWay(String source, String target) {
        for (int i = 0; i < target.length(); i++) {
            if (!source.contains(String.valueOf(target.charAt(i)))) {
                return -1;
            }
        }
        int count = 1;
        String valid = source;
        for (int i = 0; i < target.length(); i++) {
            char c = target.charAt(i);
            int index = valid.indexOf(c);
            if (index == -1) {
                count++;
                i--;
                valid = source;
            } else {
                valid = valid.substring(index);
            }
        }
        return count;
    }

    public void pwd(int index, String p, String value, Set<String> pwd) {
        for (int i = index; i < p.length(); i++) {
            String b = value + p.charAt(i);
            pwd.add(b);
            pwd(i + 1, p, b, pwd);
        }
    }

    public static void main(String[] args) {
        L_1055 l = new L_1055();
        System.out.println(l.shortestWay("abc", "abc"));
        System.out.println(l.shortestWay("abc", "abcbc"));
        System.out.println(l.shortestWay("abc", "acdbc"));
        System.out.println(l.shortestWay("xyz", "xzyxz"));
    }
}
