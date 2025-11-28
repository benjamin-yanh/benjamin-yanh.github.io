package interview_100;

public class L_487 {
    public int findMaxConsecutiveOnes(int[] nums) {
        int res = 0;

        int mark = 0;
        int count = 0;
        int len = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                len++;
            } else if (nums[i] == 0 && count < 1) {
                count++;
                len++;
                mark = i;
            } else if (nums[i] == 0 && count == 1) {
                len = i - 1 - mark;
                mark = i;
            }
            res = Math.max(res, len);
        }

        return res;
    }

    public static void main(String[] args) {
        L_487 l = new L_487();
        System.out.println(l.findMaxConsecutiveOnes(new int[]{1, 0, 1, 1, 0}));
        System.out.println(l.findMaxConsecutiveOnes(new int[]{1, 0, 1, 1, 0, 1}));
        System.out.println(l.findMaxConsecutiveOnes(new int[]{1, 0, 1, 1, 0, 0, 1}));
        System.out.println(l.findMaxConsecutiveOnes(new int[]{1, 0, 1, 1, 0, 0, 0, 0, 1}));
        System.out.println(l.findMaxConsecutiveOnes(new int[]{0, 0, 0}));
        System.out.println(l.findMaxConsecutiveOnes(new int[]{0, 0, 1}));
        System.out.println(l.findMaxConsecutiveOnes(new int[]{0}));
        System.out.println(l.findMaxConsecutiveOnes(new int[]{1}));
    }
}
