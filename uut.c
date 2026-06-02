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
  // TODO
  //goRegister(cgo_str_slice(ids));

  while (1) {
    struct goReadBlocking_return msg = goReadBlocking(ids[0], 3);
    goWriteBlocking(ids[1], msg.r0, uut(msg.r1), 4);
  }
}

/* TODO
typedef struct {
	GoInt32 adr;
	GoInt32 dat;
} message;

message *ptr = (message*)(&msg);
*/
