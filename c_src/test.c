#include <stdio.h>
#include <stdlib.h>
#include <setjmp.h>

jmp_buf env;

void risky_operation() {
    if (rand() % 2 == 0) {
        printf("Error: Risky operation failed.\n");
        longjmp(env, 1);
    }
}

int process_data(int* data) {
    if (data == NULL) {
        goto cleanup;
    }

    if (setjmp(env) == 0) {
        risky_operation();
    } else {
        printf("Recovered from non-local goto.\n");
        goto cleanup;
    }

    return 0;

cleanup:
    printf("Cleaning up resources...\n");
    return -1;
}

int main() {
    int data = 42;
    process_data(&data);
    return 0;
}
