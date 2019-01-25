// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package memory

/*
A keystore is a file containing an encrypted wallet private key.
Keystores in go-ethereum can only contain one wallet key pair per file.
To generate keystores first you must invoke NewKeyStore giving it
the directory path to save the keystores. After that, you may
generate a new wallet by calling the method NewAccount passing it
a password for encryption. Every time you call NewAccount it will
generate a new keystore file on disk.
*/
