package interview_100;

public class L_427 {
    public String stringShift(String s, int[][] shift) {

        int len = s.length();
        for (int[] ints : shift) {
            int i = ints[0];
            int j = ints[1] % len;
            if (i == 0) {
                String str = s.substring(0, j);
                s = s.substring(j) + str;
            } else {
                String str = s.substring(len - j);
                s = str + s.substring(0, len - j);
            }
        }

        return s;
    }

    public static void main(String[] args) {
        L_427 l427 = new L_427();
        System.out.println(l427.stringShift("abc", new int[][]{{0, 1}, {1, 2}}));;
    }
}
