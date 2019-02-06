#include <stdio.h>
#include <string.h>
#include <libgrpc-go.h>

GoString cgo_str(char *s) {
  GoString id = {s, strlen(s)};
  return id;
}

/*
GoSlice cgo_str_slice(GoString **s) {
  int n = sizeof(s)/sizeof(s[0]);
  GoSlice l = {s, n, n};
  return l;
}
*/

GoSlice cgo_str_slice(GoString **s) {
  int n = sizeof(s)/sizeof(s[0]);
  printf("%d\n", n);
  GoSlice l = {s, n, n};
  return l;
}

int uut(int i) {
  return 3*i;
}

int main() {
  printf("Using go gRPC client lib from C:\n");

  goConnect(cgo_str(":8888"));

  GoString ids[] = {cgo_str("uut/axis/s"), cgo_str("uut/axis/m")};

  int n = sizeof(ids)/sizeof(ids[0]);
  GoSlice l = {ids, n, n};
  goRegister(l);

  while (1) {
    goWriteBlocking(ids[1], uut(goReadBlocking(ids[0], 3)), 4);
  }
}
