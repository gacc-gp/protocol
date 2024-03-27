# Tunnel protocol

## What

Tunnel protocol defines the frame when `gacc-client` communicates with `gacc-remote`.

```
                                    --------------
                                    | gacc-remote|<-----------------
                                    --------------                 |
                                                             tunnel frame
                                                                   |
     ---------------                                        ---------------
     | game client | -----pkt------> <tun/tap> <----pkt-----| gacc-cient  |
     ---------------                                        ---------------
```

## Tunnel frame

Tunnel frame := {fixed-header:fixed32} + {message:bytes}

> version: 0...3 
> reserved: 4...15
> message-len: 16...31
fixed-header := ({version:4bits} << 28) | {message-len:16} 

message := "depends on `version`"

## V1 message

V1 message is used when `fixed-header.version` is 1. V1 message bytes are encoded with proto3, refer to `proto/tunnel/v1/`.

```
message Message {
    Header header = 1;
    google.protobuf.Any body = 2;
}
```
