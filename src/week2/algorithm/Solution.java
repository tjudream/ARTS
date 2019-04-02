package week2.algorithm;

/**
 * Created by mengxiansen on 2019/4/1.
 */
public class Solution {
    public int maxIncreaseKeepingSkyline(int[][] grid) {
        int[] maxrow = new int[grid.length];
        int[] maxcol = new int[grid[0].length];
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                maxrow[i] = Math.max(maxrow[i], grid[i][j]);
                maxcol[j] = Math.max(maxcol[j], grid[i][j]);
            }
        }
        int maxInc = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                maxInc += Math.min(maxrow[i], maxcol[j]) - grid[i][j];
            }
        }
        return maxInc;
    }
//public int maxIncreaseKeepingSkyline(int[][] grid) {
//        int[] maxrow = new int[grid.length];
//        int[] maxcol = new int[grid[0].length];
//        for (int i = 0; i < grid.length; i++) {
//            for (int j = 0; j < grid[i].length; j++) {
//                maxrow[i] = Math.max(maxrow[i], grid[i][j]);
//                maxcol[j] = Math.max(maxcol[j], grid[i][j]);
//                System.out.println("maxrow[" + i + "] = " + maxrow[i]);
//                System.out.println("maxcol[" + j + "] = " + maxcol[j]);
//            }
//        }
//        int maxInc = 0;
//        for (int i = 0; i < grid.length; i++) {
//            for (int j = 0; j < grid[i].length; j++) {
//                maxInc += Math.min(maxrow[i], maxcol[j]) - grid[i][j];
//                System.out.println("maxrow[" + i + "] = " + maxrow[i]);
//                System.out.println("maxcol[" + j + "] = " + maxcol[j]);
//                System.out.println("(" + i + "," + j + ") : " + maxInc);
//            }
//        }
//        return maxInc;
//    }
    public static void main(String[] args) {
        Solution s = new Solution();
        int[][] grid = {{3,0,8,4},{2,4,5,7},{9,2,6,3},{0,3,1,0}};
        System.out.println(s.maxIncreaseKeepingSkyline(grid));
    }
}
