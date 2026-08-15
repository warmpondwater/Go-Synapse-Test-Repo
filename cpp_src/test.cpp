#include <iostream>
#include <string>

class PaymentService {
public:
    bool processPayment(const std::string& accountId, double amount) {
        std::cout << "Processing payment for: " << accountId << " Amount: " << amount << std::endl;
        return validateAccount(accountId);
    }

private:
    bool validateAccount(const std::string& accountId) {
        return !accountId.empty();
    }
};

int main() {
    PaymentService service;
    service.processPayment("ACC-10023", 499.99);
    return 0;
}
