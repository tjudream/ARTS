package week1.algorithm;

import java.util.HashSet;
import java.util.Set;
import java.util.stream.IntStream;

/**
 * Created by mengxiansen on 2019/3/25.
 */
public class Solution {
    public int numJewelsInStonesSet(String J, String S) {
        if (J == null || S == null || "".equals(J) || "".equals(S)) {
            return 0;
        }
        Set<Character> jSet = new HashSet<>();
        for (Character c : J.toCharArray()) {
            jSet.add(c);
        }
        int result = 0;
        for (Character c : S.toCharArray()) {
            if (jSet.contains(c)) {
                result++;
            }
        }
        return result;
    }

    public int numJewelsInStonesArrSet(String J, String S) {
        if (J == null || S == null || "".equals(J) || "".equals(S)) {
            return 0;
        }
        int[] jArr = new int[128];
        for (Character c : J.toCharArray()) {
            jArr[c] = 1;
        }
        int result = 0;
        for (Character c : S.toCharArray()) {
            if (jArr[c] == 1) {
                result++;
            }
        }
        return result;
    }
    public int numJewelsInStones(String J, String S) {
        return (int)S.chars().parallel()
                .mapToObj(i -> (char)i)
                .filter(c -> J.contains(c+""))
                .count();
    }

    public int numJewelsInStonesRegular(String J, String S) {
        return S.replaceAll("[^" + J + "]", "").length();
    }
    public static void main(String[] args) {
        Solution s = new Solution();
        IntStream.of(
            s.numJewelsInStones("aA","aAAbbbb"),
            s.numJewelsInStones("z", "ZZ")
        ).forEach(System.out::println);
    }
}
