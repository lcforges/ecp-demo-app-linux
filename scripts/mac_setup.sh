#!/bin/sh
# Copyright 2023 Google LLC.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eux

ISSUER_NAME="TestIssuer"
PASSWORD="1234"
WORK_DIR=$(mktemp -d)
KEYCHAIN="BuildTest.keychain"
KEYCHAIN_TEST_BINARY=$(echo "$PWD/$(find . -iname keychain.test)")

pushd "${WORK_DIR}"

openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -sha256 -days 5 -nodes -subj "/C=US/ST=WA/L=Kirkland/O=Temp/OU=CI/CN=${ISSUER_NAME}/emailAddress=dev@example.com" 
openssl pkcs12 -inkey key.pem -in cert.pem -export -out cred.p12 -passin pass:${PASSWORD} -passout pass:${PASSWORD}

security create-keychain -p ${PASSWORD} ${KEYCHAIN}

# Disable password prompt timeout
security set-keychain-setting ${KEYCHAIN}

# Put custom keychain on keychain path
security list-keychains -d user -s ${KEYCHAIN}

security default-keychain -s "${KEYCHAIN}"

security import cred.p12 -P ${PASSWORD} -k ${KEYCHAIN} -A
security unlock-keychain -p ${PASSWORD} ${KEYCHAIN}

popd


cat << EOF > macos_config.json
{
  "cert_configs": {
    "macos_keychain": {
      "cert_issuer": "$ISSUER_NAME"
    }
  },
  "libs": {
    "ecp": "$PWD/bin/ecp",
    "ecp_client": "$PWD/bin/libecp.dylib",
    "tls_offload": "$PWD/libtls_offload.dylib"
  }
}
EOF