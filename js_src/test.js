// (DEAD CODE)
function abandonedJsHelper() {
    let unused = 42;
    console.log("Dead JS code", unused);
}

class NetworkClient {
    requestCount = 0;
    lastUrl = "";


    // (SOURCE)
    getUserSuppliedUrl() {
        return "https://api.example.com/data?user_input=admin_payload";
    }

    // (SANITIZER)
    sanitizeUrl(rawUrl) {
        if (!rawUrl) return "https://api.example.com/safe";
        return encodeURI(rawUrl);
    }

    // (SINK)
    async executeNetworkRequest(safeUrl) {
        console.log("[JS NETWORK SINK]: Fetching safe URL:", safeUrl);
        return { status: 200, data: "ok" };
    }

    async fetchResource(url) {
        this.requestCount++;
        this.lastUrl = url;
        const cleanUrl = this.sanitizeUrl(url);
        return await this.executeNetworkRequest(cleanUrl);
    }
}

async function main() {
    const client = new NetworkClient();
    try {
        const untrusted = client.getUserSuppliedUrl();
        const data = await client.fetchResource(untrusted);
        console.log("Fetch result:", data);
    } catch (e) {
        console.error("Main execution caught error:", e);
    }
}

main();

