# go-dmens-sdk

go sdk for [Dmens](https://github.com/coming-chat/Dmens) used by ComingChat App

## Usage

Maybe you have a sui account and a sui chain
```
let suiAccount = SuiNewAccount(mnemonic)
let suiChain = SuiNewChain(rpcUrl)
```

* New poster
```
let configuration = DevnetConfig()
let poster = NewPosterWithAddress(suiAccount.Address(), configuration)
```

* Build action's transaction

```go
// register or update dmens user info
let txn = Register(Profile{
    Name: "MyName"
	Bio: ""
	Avatar: "https://xxxx.xxx"
})

// post a new note
let txn = poster.DmensPost("note text content")

// replay a note
let txn = poster.DmensPostWithRef(ACTION_REPLY, "reply text content", refNoteId)

// like/repost/quote a note
let txn = poster.DmensPostWithRef(ACTION_LIKE, "", refNoteId)
let txn = poster.DmensPostWithRef(ACTION_REPOST, "", refNoteId)
let txn = poster.DmensPostWithRef(ACTION_QUOTE_POST, "", refNoteId)

// follow & unfollow other users
let txn = poster.DmensFollow([address1, address2, address3, ...])
let txn = poster.DmensUnfollow([address1, address2, address3, ...])
```

* Get max gas budget
```go
let maxGasBudget = txn.maxGasBudget
```


* Estimate transaction gas fee
```go
let gasFee = suiChain.EstimateGasFee(txn)
print("estimate transaction gas fee = " gasFee.Value)
```

* Sign & Send transaction
```go
let signedTxn = txn.SignWithAccount(suiAccount)

let txnHash = suiChain.SendRawTransaction(signedTxn.Value)

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

* Query data
  ........

* Following & Follower

  ```go
  // query the following status of a specified user.
  let isFollowing = poster.IsMyFollowing(specifiedUser)
  
  // batch query the following status of all users in a specified list.
  let userPage: *UserPage = ...
  err = poster.BatchQueryIsFollowingStatus(userPage)
  
  // query following list
  let users = poster.QueryUserFollowing("", pageSize, cursor)
  
  // query follower list
  let users = poster.QueryUserFollowers("", pageSize, cursor)
  
  // get follow count
  let counter = poster.QueryUserFollowCount("")
  print(counter.FollowerCount)
  print(counter.FollowingCount)
  ```
  


## Include content

- [x] Call Dmens contract function
- [x] fetch Dmens poster and tweets by GraphQl