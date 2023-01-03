
outdir=out

module=github.com/coming-chat/go-dmens-sdk
walletModule=github.com/coming-chat/wallet-SDK


pkgDmens=$(module)/dmens
pkgBase=$(walletModule)/core/base
pkgSui=$(walletModule)/core/sui

pkgAll=$(pkgDmens) $(pkgBase) $(pkgSui) 

buildAllSDKAndroid:
	gomobile bind -ldflags "-s -w" -target=android/arm,android/arm64 -o=${outdir}/dmens.aar ${pkgAll}

buildAllSDKIOS:
	GOOS=ios gomobile bind -ldflags "-s -w" -v -target=ios/arm64  -o=${outdir}/Dmens.xcframework ${pkgAll}

gogetGomobile:
	go get golang.org/x/mobile/cmd/gomobile & go get golang.org/x/mobile/bind