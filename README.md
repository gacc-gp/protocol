
# Auth protocol

Auth protocol must be wired with TCP.

Package := `|ver:4|message-len:28|pb-message|`

pb-message := `auth/v1/auth.proto`


# Tunnel protocol

Tunnel protocol can be wired with TCP or UDP, depends on `auth protocol`.

Package := `|fixed-header:32|payload: {header.payload-len}|`

Fixed-header := `|ver:4|flag:4|reserved:8|payload-len:16|`

flag := Frame |
        PingReq |
        PingResp |
        ConnectionClose |
        BadFrame

