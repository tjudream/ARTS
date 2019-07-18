package week17;

public class Solution {
    public String defangIPaddr(String address) {
        return address.replace(".","[.]");
    }

    public String defangIPaddr2(String address) {
        StringBuilder sb = new StringBuilder();

        for(char c: address.toCharArray()){
            if(c == '.')
                sb.append("[.]");
            else
                sb.append(c);
        }

        return sb.toString();
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.defangIPaddr("1.1.1.1"));
    }
}
