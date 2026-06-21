class NetworkClient {
    async fetchResource(url) {
        try {
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return await response.json();
        } catch (error) {
            console.error("Failed to fetch resource:", error.message);
            throw error;
        } finally {
            console.log("Fetch operation concluded.");
        }
    }
}

async function main() {
    const client = new NetworkClient();
    try {
        const data = await client.fetchResource('https://api.example.com/data');
        console.log(data);
    } catch (e) {
        console.error("Main execution caught error:", e);
    }
}

main();
