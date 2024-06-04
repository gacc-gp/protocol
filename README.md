
# Remote protocol

Remote protocol can be wired with TCP or UDP.

Message := `|FixedHeader:32|Message|`

FixedHeader := `|Ver:4|Flag:4|RESERVED:8|message-len:16|`

Flag := AuthReq |
        AuthRsp | 
        Frame |
        Ping |
        Pong |
        ConnectionClose |
        BadFrame

Message := `{Protobuf-message}`

# Regenerate

```shell
sh gen
```
