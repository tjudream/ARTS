package week5.algorithm;

import java.util.concurrent.ConcurrentHashMap;
import java.util.Map;
import java.security.SecureRandom;

//535. Encode and Decode TinyURL
public class Codec {
    private static final String ALPHABET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_";
    private static final SecureRandom RANDOM = new SecureRandom();
    private Map<String,String> map = new ConcurrentHashMap();
    private String prefixUrl = "http://tinyurl.com/";

    public static String generate(int count) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < count; ++i) {
            sb.append(ALPHABET.charAt(RANDOM.nextInt(ALPHABET.length())));
        }
        return sb.toString();
    }

    // Encodes a URL to a shortened URL.
    public String encode(String longUrl) {
        String shortUrl = prefixUrl + generate(6);
        while (map.containsKey(shortUrl)) {
            shortUrl = prefixUrl + generate(6);
        }
        map.put(shortUrl, longUrl);
        return shortUrl;
    }

    // Decodes a shortened URL to its original URL.
    public String decode(String shortUrl) {
        return map.get(shortUrl);
    }

    public static void main(String[] args) {
        String[] urls = new String[]{
                "https://leetcode.com/problems/design-tinyurl"
        };
        Codec codec = new Codec();
        for (String url : urls) {
            System.out.println(codec.decode(codec.encode(url)));
        }
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.decode(codec.encode(url));
