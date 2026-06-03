/*
Copyright 2018-2026
Unai Martinez-Corral <unai.martinezcorral@ehu.eus>
University of the Basque Country (UPV/EHU) <ehu.eus>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

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
