# go-dmens-sdk

go sdk for [Dmens](https://github.com/coming-chat/Dmens) used by ComingChat App

## Usage

Maybe you have a sui account and a sui chain
```golang
var suiAccount = SuiNewAccount(mnemonic)
var suiChain = SuiNewChain(rpcUrl)
```

* New poster
```golang
var configuration = DevnetConfig()
configuration.FullNodeUrl = "https://fullnode.testnet.sui.io" // Custom rpc url
var poster = NewPosterWithAddress(suiAccount.Address(), configuration)

// switch rpcUrl
poster.SwitchRpcUrl("http://....")
```

* Build action's transaction

```go
// register or update dmens user info
var txn = Register(Profile{
    Name: "MyName"
	Bio: ""
	Avatar: "https://xxxx.xxx"
})

// post a new note
var txn = poster.DmensPost("note text content")

// replay a note
var txn = poster.DmensPostWithRef(ACTION_REPLY, "reply text content", refNoteId)

// like/repost/quote a note
var txn = poster.DmensPostWithRef(ACTION_LIKE, "", refNoteId)
var txn = poster.DmensPostWithRef(ACTION_REPOST, "", refNoteId)
var txn = poster.DmensPostWithRef(ACTION_QUOTE_POST, "", refNoteId)

// follow & unfollow other users
var txn = poster.DmensFollow([address1, address2, address3, ...])
var txn = poster.DmensUnfollow([address1, address2, address3, ...])
```

* Get max gas budget
```go
var maxGasBudget = txn.maxGasBudget
```


* Estimate transaction gas fee
```go
var gasFee = suiChain.EstimateGasFee(txn)
print("estimate transaction gas fee = " gasFee.Value)
```

* Sign & Send transaction
```go
var signedTxn = txn.SignWithAccount(suiAccount)

var txnHash = suiChain.SendRawTransaction(signedTxn.Value)

print("transaction hash = ", txnHash.Value)

```

* NFT Avatar
```go
// get user's nft avatar
var user: UserInfo = ...
print(user.NFTAvatar)

// query nft avatar by nftid
var avatar = poster.QueryNFTAvatar(nftid)

// batch query nft avatar for user page
var userPage: UserPage = ...
err = poster.BatchQueryNFTAvatarForUserPage(userPage)

// Transaction
// set avatar transaction
var txn = poster.SetNftAvatar(nftid)

// remove avatar transaction
var user: UserInfo = ...
var txn = poster.RemoveNftAvatar(user.NFTAvatar)

// sign & send transaction ...

```

* Sui Name

```go
type UserInfo {
	// Only queried when call QueryUserInfoByAddress
	SuiName string `json:"suiName"`
  ......
}

var user = poster.QueryUserInfoByAddress(address)
print(user.SuiName)

// only query sui name
var name = poster.QuerySuiNameByAddress(address)
print(name)
```



* Query data
  ........

* Following & Follower

  ```go
  // query the following status of a specified user.
  var isFollowing = poster.IsMyFollowing(specifiedUser)
  
  // batch query the following status of all users in a specified list.
  var userPage: *UserPage = ...
  err = poster.BatchQueryIsFollowingStatus(userPage)
  
  // query following list
  var users = poster.QueryUserFollowing("", pageSize, cursor)
  
  // query follower list
  var users = poster.QueryUserFollowers("", pageSize, cursor)
  
  // get follow count
  var counter = poster.QueryUserFollowCount("")
  print(counter.FollowerCount)
  print(counter.FollowingCount)
  ```



## Include content

- [x] Call Dmens contract function
- [x] fetch Dmens poster and tweets by GraphQl