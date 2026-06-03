# Copyright 2018-2026
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


FROM golang:bookworm

RUN apt update && apt install -y unzip dos2unix \
 && wget https://github.com/protocolbuffers/protobuf/releases/download/v35.0/protoc-35.0-linux-x86_64.zip -O protoc.zip \
 && unzip protoc.zip -d /usr/local
