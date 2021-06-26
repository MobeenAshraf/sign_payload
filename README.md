## Utility CLI to sign Rosetta payloads message with user private key


# Clone Repo
```
git clone https://github.com/MobeenAshraf/sign_payload.git
cd sign_payload
```

# Build

```
go mod download
 or
go get ./..

go mod tidy

go build
```


# Run

```
./sign_payload -p PRIVATE_KEY  -sender SENDER_ADDRESS -m MESSAGE
```

# CLI Flags / Option 

```
./sign_payload --help
```

```
Usage of ./sign_payload:
  -m string
    	Signing Message tbu for Payload
  -p string
    	Private Key Hex (32 bytes) 64 chars
  -psig string
    	Payload Signature. Default: ecdsa_recovery (default "ecdsa_recovery")
  -sender string
    	Sender Address
  -ssig string
    	Signing Signature Signature. Default: ecdsa (default "ecdsa")
```