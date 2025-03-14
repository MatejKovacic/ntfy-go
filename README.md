# Ntfy Go App for ARM devices

This is a simple Go app for sending Ntfy messages from ARM devices that **do not have** `curl` or `wget` installed, or those apps **do not support HTTPS connections**.

**Usage**: `send_ntfy -s <server> -t <topic> -m <message> [-u <username> -p <password>]`

**Precompiled binaries**:
- x86 architecture: [send_ntfy_amd64](send_ntfy_amd64)
- ARM64 architecture: [send_ntfy_arm64](send_ntfy_arm64)
- ARM32 architecture: [send_ntfy_arm](send_ntfy_arm)

## How to test

### Setup app on your phone

First install Ntfy app on your mobile phone. Since official Ntfy.sh server is not very reliable, I suggest using alternative Ntfy server (`https://ntfy.envs.net/`). Go to `Settings` - `Server URL` and enter `https://ntfy.envs.net/`.

Then tap the `+` button and enter the topic name you want to subscribe to. You can made up anything for topic name, however it should be unique enough to be unlikely to guess. Why? Because topics are basically public (unless you use authentication)! In this example I am assuming your topic name is `MyUniqueTopic`.

Also, I recommend to set that Ntfy has unlimited battery usage and that Android does not remove permissions if app is not used for some time.

### Test if receiving notifications is working

You can **test your setup** with your Linux computer with the following commands:
```
curl -d "TEST message!" https://ntfy.envs.net/MyUniqueTopic.
```
or:
```
wget --post-data="TEST message!" https://ntfy.envs.net/MyUniqueTopic -O - &> /dev/null
```

### Test Go app on your computer
If everything works, run the Go app:
```
go run send_ntfy.go -m "My first test message from Go" -t MyUniqueTopic -s https://ntfy.envs.net
```
If you are using authentication, use `-u` (username) and `-p` (password) parameters.

## Compile ARM versions

Now you can cross compile Go app for **ARM64** architecture:
```
GOOS=linux GOARCH=arm64 go build -o send_ntfy_arm64 send_ntfy.go
```

If your device is **32-bit ARM** (for instance `armv7l` - please use `uname -a` command to identify architecture of your device), you should cross compile Go app for **ARM32** architecture (this is the case for *TP-Link M7350* device):
```
GOOS=linux GOARCH=arm GOARM=7 go build -o send_ntfy_arm send_ntfy.go
```

Now copy the file `send_ntfy_arm64` (or `send_ntfy_arm`) to your ARM device, make it executable (`chmod +x send_ntfy_arm64` or `chmod +x send_ntfy_arm`) and run it:
```
./send_ntfy_arm64 -m "Test message from ARM device" -t MyUniqueTopic -s https://ntfy.envs.net
```
or (for 32-bit version):
```
./send_ntfy_arm -m "Test message from ARM device" -t MyUniqueTopic -s https://ntfy.envs.net
```
