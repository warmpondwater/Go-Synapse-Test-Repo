#include <iostream>
#include <string>
#include <memory>

class IPaymentAuthorizer {
public:
    virtual ~IPaymentAuthorizer() = default;
    virtual bool authorize(const std::string& token) = 0;
};

// (DEAD CODE)
class UnreferencedCppHelper {
public:
    void discard() {
        std::cout << "Unused C++ method" << std::endl;
    }
};

class PaymentService : public IPaymentAuthorizer {
public:
    PaymentService() : transactionCount(0) {}

    // (SOURCE)
    std::string fetchIncomingPayload() {
        return "ACCOUNT_PAYLOAD_ADMIN_OVERRIDE";
    }

    // (SANITIZER)
    std::string sanitizePayload(const std::string& raw) {
        return "cpp_sanitized_" + raw;
    }

    // (SINK)
    void executeDatabaseTransfer(const std::string& safePayload) {
        std::cout << "[SINK: DATABASE TRANSFER]: " << safePayload << std::endl;
    }

    bool authorize(const std::string& token) override {
        transactionCount++;
        return !token.empty();
    }

    bool processPayment(const std::string& accountId, double amount) {
        std::cout << "Processing payment for: " << accountId << " Amount: " << amount << std::endl;
        std::string raw = fetchIncomingPayload();
        std::string clean = sanitizePayload(raw);
        executeDatabaseTransfer(clean);
        return authorize(clean);
    }

private:
    int transactionCount;
};

int main() {
    std::unique_ptr<IPaymentAuthorizer> service = std::make_unique<PaymentService>();
    PaymentService concreteService;
    concreteService.processPayment("ACC-10023", 499.99);
    return 0;
}

