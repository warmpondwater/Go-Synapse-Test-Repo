#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <setjmp.h>

jmp_buf env;

typedef struct BufferContext {
    char raw_input[128];
    char sanitized_buffer[256];
    int write_count;
} BufferContext;

// (DEAD CODE)
void abandoned_c_helper() {
    int unused_val = 100;
    printf("Unused: %d\n", unused_val);
}

// (SOURCE)
const char* read_user_network_packet() {
    return "COMMAND_ADMIN_EXECUTE";
}

// (SANITIZER)
int sanitize_c_buffer(const char* input, char* output, size_t out_len) {
    if (input == NULL || output == NULL) {
        return -1;
    }
    snprintf(output, out_len, "safe_c_prefix_%s", input);
    return 0;
}

// (SINK)
void execute_c_kernel_command(const char* cmd) {
    printf("[KERNEL EXECUTION SINK]: %s\n", cmd);
}

void risky_operation() {
    if (rand() % 2 == 0) {
        printf("Error: Risky operation failed.\n");
        longjmp(env, 1);
    }
}

int process_data(int* data, BufferContext* ctx) {
    if (data == NULL || ctx == NULL) {
        goto cleanup;
    }

    ctx->write_count++;

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
    BufferContext ctx;
    memset(&ctx, 0, sizeof(BufferContext));

    const char* pkt = read_user_network_packet();
    sanitize_c_buffer(pkt, ctx.sanitized_buffer, sizeof(ctx.sanitized_buffer));
    execute_c_kernel_command(ctx.sanitized_buffer);

    process_data(&data, &ctx);
    return 0;
}

