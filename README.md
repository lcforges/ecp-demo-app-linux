# ECP Encryption/Decryption Demo App

This app demonstrates how to encrypt and decrypt messages using the Enterprise Certificate Proxy (ECP) client library.

Usage
To use the app, you need to specify the following flags:

mode: The mode of operation, either -e for encryption or -d for decryption.
msg: The message to encrypt or decrypt.
For example, to encrypt the message "Hello, world!", you would run the following command:

go run main.go -mode -e -msg "Hello, world!"
The app will print the encrypted message to the console.

To decrypt the message, you would run the following command:

go run main.go -mode -d -msg <encrypted message>
The app will print the decrypted message to the console.

Requirements
Go 1.17 or higher
The ECP client library
Installation
To install the ECP client library, run the following command:

go get github.com/googleapis/enterprise-certificate-proxy/client
License
This app is licensed under the Apache License, Version 2.0.