package week5.algorithm;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.zip.GZIPInputStream;
import java.util.zip.GZIPOutputStream;

public class CodecByAlgo {

    private String prefixUrl = "http://tinyurl.com/";

    public static String byteArrayToHex(byte[] a) {
        StringBuilder sb = new StringBuilder(a.length * 2);
        for(byte b: a)
            sb.append(String.format("%02x", b));
        return sb.toString();
    }
    public static byte[] hexToByteArray(String hex) {
        byte[] val = new byte[hex.length() / 2];
        for (int i = 0; i < val.length; i++) {
            int index = i * 2;
            int j = Integer.parseInt(hex.substring(index, index + 2), 16);
            val[i] = (byte) j;
        }
        return val;
    }

    // Encodes a URL to a shortened URL.
    public String encode(String longUrl) {
        ByteArrayOutputStream out = null;
        GZIPOutputStream gzip = null;
        String compress = "";
        try {
            out = new ByteArrayOutputStream();
            gzip = new GZIPOutputStream(out);
            gzip.write(longUrl.getBytes());
            gzip.close();
            String hex = byteArrayToHex(out.toByteArray());
            //compress = out.toString("ISO-8859-1");
            compress = prefixUrl + hex;
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (null != out) {
                try {
                    out.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
        return compress;
    }

    // Decodes a shortened URL to its original URL.
    public String decode(String shortUrl) {
        String hex = shortUrl.replace(prefixUrl, "");
        byte[] val = hexToByteArray(hex);
        ByteArrayOutputStream out = null;
        ByteArrayInputStream in = null;
        GZIPInputStream gzip = null;
        String uncompress = "";
        try {
            out = new ByteArrayOutputStream();
            in = new ByteArrayInputStream(val);
            gzip = new GZIPInputStream(in);
            byte[] buffer = new byte[1024];
            int offset = -1;
            while ((offset = gzip.read(buffer)) != -1) {
                out.write(buffer, 0, offset);
            }
            uncompress = out.toString();
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (null != gzip) {
                try {
                    gzip.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            if (null != in) {
                try {
                    in.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            if (null != out) {
                try {
                    out.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
        return uncompress;
    }

    public static void main(String[] args) {
        String[] urls = new String[]{
                "https://leetcode.com/problems/design-tinyurl"
        };
        CodecByAlgo codec = new CodecByAlgo();
        for (String url : urls) {
            System.out.println(codec.encode(url));
            System.out.println(codec.decode(codec.encode(url)));
        }
    }
}
