
# Auth protocol

Auth protocol must be wired with TCP.

Message := `|FixHeader:32|pb-message|`

FixHeader := `|ver:4|message-len:28|`

pb-message := `auth/v1/auth.proto`

# Tunnel protocol

Tunnel protocol can be wired with TCP or UDP, depends on `auth protocol`.

Message := `|FixedHeader:32|payload|`

FixedHeader := `|ver:4|flag:4|RESERVED:8|payload-len:16|`

flag := Frame |
        PingReq |
        PingResp |
        ConnectionClose |
        BadFrame

