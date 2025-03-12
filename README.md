# Ntfy Go App for ARM devices

This is a simple Go app for sending Ntfy messages from ARM devices that **do not have** `curl` or `wget` installed, or those apps **do not support HTTPS connections**.

First install Ntfy app on your mobile phone. Since official Ntfy.sh server is not very reliable, I suggest using alternative Ntfy server (`https://ntfy.envs.net/`). Go to `Settings` - `Server URL` and enter `https://ntfy.envs.net/`.

Then tap the `+` button and enter the topic name you want to subscribe to. You can made up anything for topic name, however it should be unique enough to be unlikely to guess. Why? Because topics are basically public (unless you use authentication)! In this example I am assuming your topic name is `MyUniqueTopic`.

Also, I recommend to set that Ntfy has unlimited battery usage and that Android does not remove permissions if app is not used for some time.

You can **test your setup** with your Linux computer with the following commands:
```
curl -d "TEST message!" https://ntfy.envs.net/MyUniqueTopic.
```
or:
```
wget --post-data="TEST message!" https://ntfy.envs.net/MyUniqueTopic -O - &> /dev/null
```

If everything works, run the Go app. Please note that server should be without trailing slash:
```
go run send_ntfy.go -m "My first test message from Go" -t MyUniqueTopic -s https://ntfy.envs.net
```

Now you can cross compile Go app for Arm64 architecture:
```
set GOOS=linux
set GOARCH=arm64
go build -o send_ntfy_arm64 send_ntfy.go
```

Copy the file `send_ntfy_arm64` to your ARM device, make it executable (`chmod +x send_ntfy_arm64`) and run it:
```
./send_ntfy_arm64 -m "Test message from ARM device" -t MyUniqueTopic -s https://ntfy.envs.net
```

**Please note that this example does not support authentication for Ntfy servers.**
