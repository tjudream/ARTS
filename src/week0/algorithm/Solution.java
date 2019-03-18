package week0.algorithm;

import java.util.Arrays;

/**
 * Created by mengxiansen on 2019/3/16.
 */
class Solution {
    // 算法1：暴力法
//    public int lengthOfLongestSubstring(String s) {
//        if (s == null || "".equals(s)) {
//            return 0;
//        }
//        int n = s.length();
//        int len = 1;
//        for (int i = 0; i < n; i++) {
//            for (int j = i + 1; j < n; j++) {
//                boolean isContains = false;
//                String subStr = s.substring(i, j);
//                for (int k = 0; k < subStr.length(); k++) {
//                    if (s.charAt(j) == subStr.charAt(k)) {
//                        isContains = true;
//                        break;
//                    }
//                }
//                if (isContains) {
//                    break;
//                }
//                len = Math.max(len, subStr.length() + 1);
//            }
//        }
//        return len;
//    }
//算法2：窗口切片法
//    public int lengthOfLongestSubstring(String s) {
//        if (s == null || "".equals(s)) {
//            return 0;
//        }
//        int n = s.toCharArray().length;
//        Map<Character,Integer> map = new ConcurrentHashMap<>();
//        int start = 0;
//        int len = 1;
//        for (int i = 0; i < n; i++) {
//            if (map.keySet().contains(s.charAt(i))) {
//                start = Math.max(start, map.get(s.charAt(i)));
//            }
//            map.put(s.charAt(i), i + 1);
//            len = Math.max(len, i - start + 1);
//        }
//        return len;
//    }

    // 算法3：对算法2的优化
    public int lengthOfLongestSubstring(String s) {
        if (s == null || "".equals(s)) {
            return 0;
        }
        int n = s.length();
        int len = 0;
        int[] index = new int[128];
        for (int i = 0, j = 0; j < n; j++) {
            i = Math.max(i, index[s.charAt(j)]);
            index[s.charAt(j)] = j + 1;
            len = Math.max(len, j - i + 1);
        }
        return len;
    }
    public static void main(String[] args) {
        Solution s = new Solution();
        Arrays.asList(
                s.lengthOfLongestSubstring("abcabcbb"),
                s.lengthOfLongestSubstring("bbbbb"),
                s.lengthOfLongestSubstring("pwwkew"),
                s.lengthOfLongestSubstring("abcdbdfhik"),
                s.lengthOfLongestSubstring("dvdf"),
                s.lengthOfLongestSubstring(" ")
        ).stream().forEach(System.out::println);

    }
}
