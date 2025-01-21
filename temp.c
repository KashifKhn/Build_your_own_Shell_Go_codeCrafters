#include <stdio.h>

unsigned long calculate_signature(const char *program_name) {
  unsigned long signature = 0;
  while (*program_name) {
    signature = signature * 31 + *program_name;
    program_name++;
  }
  return signature;
}

int main(int argc, char *argv[]) {
  printf("Program was passed %d args (including program name).\n", argc);

  for (int i = 0; i < argc; i++) {
    printf("Arg #%d%s: %s\n", i, (i == 0) ? " (program name)" : "", argv[i]);
  }

  unsigned long signature = calculate_signature(argv[0]);
  printf("Program Signature: %lu\n", signature);

  return 0;
}
