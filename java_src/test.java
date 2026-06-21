import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class PolyglotTest {
    public static String readFile(String path) throws IOException {
        try (BufferedReader reader = new BufferedReader(new FileReader(path))) {
            StringBuilder sb = new StringBuilder();
            String line;
            while ((line = reader.readLine()) != null) {
                sb.append(line).append("\n");
            }
            return sb.toString();
        }
    }

    public static void main(String[] args) {
        try {
            String content = readFile("config.txt");
            System.out.println(content);
        } catch (IOException e) {
            System.err.println("Failed to read file: " + e.getMessage());
        }
    }
}
