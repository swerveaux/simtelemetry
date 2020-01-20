This is a small Go app that starts a UDP server on 4843 an a webserver on 10001.   It's designed to take telemetry data from Forza 7 and Forza Horizon 4 and display it on a webpage.

Both games are in UWP sandboxes and have to be taken out of those sandboxes to work.   I wish someone had told me this.

Fire up Powershell and run 

```CheckNetIsolation LoopbackExempt -a -n=Microsoft.SunriseBaseGame_8wekyb3d8bbwe```

to allow Forza Horizon 4 to talk to loopback (127.0.0.1) and

```CheckNetIsolation LoopbackExempt -a -n=Microsoft.ApolloBaseGame_8wekyb3d8bbwe```

for Forza 7.   Otherwise, it will not emit data to your local machine and lead me to hours of frustration.

It looks like FH4 does not send the "V2" data, which includes things like what gear you're in or how hard you have the gas/brake/whatever pressed.   That's a bummer.

