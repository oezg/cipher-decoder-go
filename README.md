# Cipher Decoder

## About
Cryptography is a crucial concept when some things are meant to be kept secret. One of the main issues of cryptography is how to send an encryption key to another person over an unsecured communication channel. In this project, I apply the Diffie-Hellman key exchange protocol to keep the conversation private.

## Learning Outcomes
In this project, I studied 
- input and output, 
- the Diffie-Hellman protocol, 
- a basic cipher to keep conversations private.

## Description
Messages will be encrypted using the Caesar cipher.
Each letter of the alphabet is cyclically shifted down the alphabet by a certain number of positions. 
I will encrypt only letters; other characters will remain as they are.

## Objectives
- Encrypt the message `Will you marry me?` 
- Send it to Alice. 
  - If her answer is `Yeah, okay!`, encrypt `Great!` and send it back. 
  - If Alice's answer is `Let's be friends.`, encrypt and output `What a pity!` 
  - Otherwise, don't print anything.