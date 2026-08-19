import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

interface IDataProcessor {
    String processPayload(String input);
}

// (DEAD CODE)
class AbandonedJavaHelper {
    public void unusedRoutine() {
        System.out.println("Dead java code");
    }
}

public class PolyglotTest implements IDataProcessor {
    private int processedCount = 0;

    // (SOURCE)
    public String getUntrustedPayload() {
        return "DROP TABLE accounts; --";
    }

    // (SANITIZER)
    public String sanitizePayload(String input) {
        if (input == null) return "";
        return input.replace("'", "''").replace(";", "");
    }

    // (SINK)
    public void persistToDatabase(String safePayload) {
        System.out.println("[JAVA SINK DATABASE]: " + safePayload);
    }
    public String processPayload(String input) {
        processedCount++;
        String clean = sanitizePayload(input);
        persistToDatabase(clean);
        return clean;
    }

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
        IDataProcessor processor = new PolyglotTest();
        PolyglotTest concrete = new PolyglotTest();
        String raw = concrete.getUntrustedPayload();
        processor.processPayload(raw);
    }
}

