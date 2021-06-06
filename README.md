# gochain

This is simply an extended hello world to get in touch with go. The base structure is a Block with some fields combined
to a chain. Some sample output of the program looks like this:

```log
GOROOT=/Users/manuel/.gvm/gos/go1.16.5 #gosetup
GOPATH=/Users/manuel/workspace/private/gochain/.lgopath:/Users/manuel/.gvm/pkgsets/go1.16.5/global #gosetup
/Users/manuel/.gvm/gos/go1.16.5/bin/go build -o /private/var/folders/rx/xlrjwzts3yqcdv16vm7gqlp80000gn/T/___go_build_gochain_go /Users/manuel/workspace/private/gochain/gochain.go #gosetup
/private/var/folders/rx/xlrjwzts3yqcdv16vm7gqlp80000gn/T/___go_build_gochain_go
INFO   [2021-06-06T20:26:21+02:00] starting application                         
DEBUG  [2021-06-06T20:26:21+02:00] block0: Block(number=0,tx=Tx(source=soSource1,target=noTarget1,amount=0),comment=genesis,ts=1623003981539181000,last=0,current=8e3d8b5958f0f05bfc37efe1ae9bea4c61e8285af916cb9e3ce488279cd1fa5f57efd894d80d8234aeb57fba6f005ddebef6aab090c1b4e41ac82f0f3b508b78) 
DEBUG  [2021-06-06T20:26:21+02:00] block1: Block(number=1,tx=Tx(source=random1,target=mbo,amount=1.5),comment=first tx,ts=1623003981539376000,last=8e3d8b5958f0f05bfc37efe1ae9bea4c61e8285af916cb9e3ce488279cd1fa5f57efd894d80d8234aeb57fba6f005ddebef6aab090c1b4e41ac82f0f3b508b78,current=69a774938fa0b534877a978e737f9ee092c77ebef057fa73353b351d782c0f39036ea55db61f1a8d27b3fd691d4a79fe8cef10c2ec640a0fa62380a84c15c1a9) 
DEBUG  [2021-06-06T20:26:21+02:00] block2: Block(number=2,tx=Tx(source=random2,target=mbo,amount=2.5),comment=second tx,ts=1623003981539432000,last=69a774938fa0b534877a978e737f9ee092c77ebef057fa73353b351d782c0f39036ea55db61f1a8d27b3fd691d4a79fe8cef10c2ec640a0fa62380a84c15c1a9,current=b25f191d66095a231b65a9d88dde53b5c93f870147e0b663a466915404bee9e9b90ea38dceb8999592b78ae4078be494263ecaf6143ed6f1aa4f33cd9d398732) 
DEBUG  [2021-06-06T20:26:21+02:00] block3: Block(number=3,tx=Tx(source=random3,target=mbo,amount=3.5),comment=third tx,ts=1623003981539761000,last=b25f191d66095a231b65a9d88dde53b5c93f870147e0b663a466915404bee9e9b90ea38dceb8999592b78ae4078be494263ecaf6143ed6f1aa4f33cd9d398732,current=93ca881794ddf28741f0506bfcd3cb698f534f6795f4decb0c484eb4d1c69b6b0408eb9df0cca4d882d033562684e6c40bf1366a0f74a19926671d7d602736e7) 
INFO   [2021-06-06T20:26:21+02:00] done   
```

The sample uses the following packages:

- logrus: Logging library from external source which supports different log levels.
- crypto/sha512: to create hashes out of the blocks for the blockchain

Other libs seem to be pretty standard and not worth mentioning.

## go environment

On my mac I used gvm for being able to have multiple go environments available at the same time. For installing gvm I
had to install a go version from https://golang.org/ first because the gvm compile process used in gvm already requires
an installed go version.
