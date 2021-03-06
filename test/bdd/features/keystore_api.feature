#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

@all
@keystore
Feature: Keystore operations

  Scenario: User creates a keystore
    Given Key Server is running on "localhost" port "8076"
    When  User sends an HTTP POST to "https://localhost:8076/kms/keystores" to create a keystore
    Then  User gets a response with HTTP 201 Created and Location with a valid URL for the newly created keystore
