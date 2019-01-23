package common

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"
	"unsafe"

	//"gitlab.33.cn/wallet/privacystockapi/common"

	. "github.com/33cn/plugin/plugin/dapp/privacy/crypto"
	sccrypto "github.com/NebulousLabs/Sia/crypto"
	"github.com/33cn/chain33/common"
	"github.com/stretchr/testify/require"
)



var (

	pub string = "1D9xKRnLvV2zMtSxSx33ow1GF4pcbLcNRt"
	pri string = "0xa0c6f46de8d275ce21e935afa5363e9b8a087fe604e05f7a9eef1258dc781c3a"
)




func TestHash(t *testing.T){
	privByte, _ := common.FromHex(pri)
	println(len(privByte))
	privKey := (*[KeyLen32]byte)(unsafe.Pointer(&privByte[0]))

	hash := sccrypto.HashAll(*privKey)


	rootPrivacy := &Privacy{}
	var err error

	if err = generateKeyPairWithPrivKey((*[KeyLen32]byte)(unsafe.Pointer(&hash[0])), &rootPrivacy.SpendPrivKey, &rootPrivacy.SpendPubkey); err != nil {
		fmt.Println(err.Error())
	}

	hashViewPriv := sccrypto.HashAll(rootPrivacy.SpendPrivKey[0:KeyLen32])
	if err = generateKeyPairWithPrivKey((*[KeyLen32]byte)(unsafe.Pointer(&hashViewPriv[0])), &rootPrivacy.ViewPrivKey, &rootPrivacy.ViewPubkey); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("keyPair", "pub A", common.Bytes2Hex(rootPrivacy.ViewPubkey[:]))
	fmt.Println("keyPair", "priv a", common.Bytes2Hex(rootPrivacy.ViewPrivKey[:]))


	start := time.Now().UnixNano()
	uid := "137184"
	privacy1 := &Privacy{}
	for i := 0; i < 1; i++ {

		hash = sccrypto.HashAll(rootPrivacy.ViewPrivKey[0:KeyLen32], []byte(uid))
		if err = generateKeyPairWithPrivKey((*[KeyLen32]byte)(unsafe.Pointer(&hash[0])), &privacy1.ViewPrivKey, &privacy1.ViewPubkey); err != nil {
			fmt.Println(err.Error())
		}
	}
	end := time.Now().UnixNano()

	fmt.Println("keyPair", "pub", common.Bytes2Hex(privacy1.ViewPubkey[:]))
	fmt.Println("keyPair", "priv", common.Bytes2Hex(privacy1.ViewPrivKey[:]))
	fmt.Println("keyPair", "hash", common.Bytes2Hex(hash[:]))
	fmt.Println("ExcuteTime", "microSeconds", float32(end - start)/1000/1)

	start = time.Now().UnixNano()

	uid = "183284"
	for i := 0; i < 1; i++{

		pk := &PubKeyPrivacy{}
		sk := &PrivKeyPrivacy{}
		GenerateKeyPair(sk, pk)

		oneTimePub, _ := GenerateOneTimeAddr((*[32]byte)(unsafe.Pointer(&rootPrivacy.ViewPubkey[0])),
			(*[32]byte)(unsafe.Pointer(&rootPrivacy.SpendPubkey[0])), (*[32]byte)(unsafe.Pointer(&sk[0])), int64(i))

		hex.EncodeToString(oneTimePub[:])
	}
	end = time.Now().UnixNano()
	fmt.Println("ExcuteTime", "microSeconds", float32(end - start)/1000/1)
}

func TestNewPrivacyWithPrivKeyEx(t *testing.T) {
	uid := "123456"
	privByte, err := common.FromHex(pri)
	require.NoError(t, err)

	pritest := (*[KeyLen32]byte)(unsafe.Pointer(&privByte[0]))
	privacy1, err := NewPrivacyWithPrivKeyEx(pritest, []byte(uid))
	fmt.Println("keyPair", "pub", common.Bytes2Hex(privacy1.ViewPubkey[:]))
	fmt.Println("keyPair", "priv", common.Bytes2Hex(privacy1.ViewPrivKey[:]))
	require.NoError(t, err)


	hash := sccrypto.HashAll(*pritest, []byte(uid))
	rootPrivacy := &Privacy{}
	err = generateKeyPairWithPrivKey((*[KeyLen32]byte)(unsafe.Pointer(&hash[0])), &rootPrivacy.SpendPrivKey, &rootPrivacy.SpendPubkey)
	require.NoError(t, err)
	hashViewPriv := sccrypto.HashAll(rootPrivacy.SpendPrivKey[0:KeyLen32])
	err = generateKeyPairWithPrivKey((*[KeyLen32]byte)(unsafe.Pointer(&hashViewPriv[0])), &rootPrivacy.ViewPrivKey, &rootPrivacy.ViewPubkey)
	fmt.Println("keyPair", "pub", common.Bytes2Hex(rootPrivacy.ViewPubkey[:]))
	fmt.Println("keyPair", "priv", common.Bytes2Hex(rootPrivacy.ViewPrivKey[:]))
	require.NoError(t, err)
	require.Equal(t, privacy1, rootPrivacy)
}

func TestDivisionAmount(t *testing.T) {
	amount := Coin * 99
	amounts := DivisionAmount(amount)
	require.Equal(t, len(amounts), 1)
	require.Equal(t, amount, amounts[0])

	amount = Coin * 100
	amounts = DivisionAmount(amount)
	require.Equal(t, len(amounts), 1)
	require.Equal(t, amount, amounts[0])

	amount = Coin * 101
	amounts = DivisionAmount(amount)
	require.Equal(t, len(amounts), 2)
	require.Equal(t, 100 * Coin, amounts[0])
	require.Equal(t, 1 * Coin, amounts[1])

	amount = Coin * 1000
	amounts = DivisionAmount(amount)
	require.Equal(t, len(amounts), 1)
	require.Equal(t, 1000 * Coin, amounts[0])

	amount = Coin * 1001
	amounts = DivisionAmount(amount)
	require.Equal(t, len(amounts), 2)
	require.Equal(t, 1000 * Coin, amounts[0])
	require.Equal(t, 1 * Coin, amounts[1])

	amount = Coin * 999
	amounts = DivisionAmount(amount)
	require.Equal(t, len(amounts), 10)
	require.Equal(t, 100 * Coin, amounts[0])
	require.Equal(t, 100 * Coin, amounts[6])
	require.Equal(t, 99 * Coin, amounts[len(amounts) - 1])

	amount = Coin * 100111
	amounts = DivisionAmount(amount)
	require.Equal(t, len(amounts), 102)
	require.Equal(t, 1000 * Coin, amounts[0])
	require.Equal(t, 1000 * Coin, amounts[6])
	require.Equal(t, 1000 * Coin, amounts[50])
	require.Equal(t, 1000 * Coin, amounts[99])
	require.Equal(t, 100 * Coin, amounts[len(amounts) - 2])
	require.Equal(t, 11 * Coin, amounts[len(amounts) - 1])

}