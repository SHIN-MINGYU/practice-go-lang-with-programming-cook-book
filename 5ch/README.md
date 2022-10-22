# Network Programming

- Go standard library provide various feature for network action

  - In here, there is package for adminstrating TCP/IP, UDP and DNS, Mail, RPC using HTTP
  - Third party libraries can fill scare part of feature
  - For instance, gorilla/websockets provide feature for implying websocket in normal HTTP handler

- This example cant show high-level interfcaing like REST or RRPC, In contrast, can show How to work network action
- This is useful for DevOps application working with raw-eamil or DNS searching

## What will we do?

- Create TCP/IP Echo server and client
- Create UDP server and client
- Do search domain name
- Work with web-socket
- Call remote-function using for net/rpc
- Parsing e-mail using for net/mail
