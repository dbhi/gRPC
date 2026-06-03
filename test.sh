#!/usr/bin/env bash

# Copyright 2026
# Unai Martinez-Corral <unai.martinezcorral@ehu.eus>
# University of the Basque Country (UPV/EHU) <ehu.eus>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

set -e

cd "$(dirname "$0")"

./dist/server & SERVER_PID=$!

sleep 1

./dist/gen & GEN_PID=$!

sleep 1

LD_LIBRARY_PATH=$(pwd)/dist ./dist/"${UUT:-uut-go}" & UUT_PID=$!

while kill -0 $GEN_PID >/dev/null 2>&1; do sleep 1; done
kill $UUT_PID $SERVER_PID
