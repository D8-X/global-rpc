# global-rpc
EVM RPC and WSS manager

# RPC Config File Example

```
[
    {
        "chainId": 999,
        "https": [
            "https://rpc.hyperliquid.xyz/evm",
            "https://rpc.hypurrscan.io",
            "https://hyperliquid.drpc.org",
            "https://rpc.hyperlend.finance",
        ],
        "wss": [
            "wss://hyperliquid.drpc.org",
            "wss://api.hyperliquid.xyz/ws"
        ]
    },
    {
        "chainId": 989,
        "https": [
            "https://spectrum-01.simplystaking.xyz/hyperliquid-tn-rpc/evm",
            "https://rpc.hyperliquid-testnet.xyz/evm"
        ],
        "wss": []
    }
]
```