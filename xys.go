package xs

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var commonBase64 = base64.NewEncoding("ZmserbBoHQtNP+wOcza/LpngG8yJq42KWYj0DSfdikx3VT16IlUAFM97hECvuRX5").WithPadding(base64.NoPadding)

type xys struct {
	X0 string `json:"x0"`
	X1 string `json:"x1"`
	X2 string `json:"x2"`
	X3 string `json:"x3"`
	X4 string `json:"x4"`
}

func XYS(params, a1 string) string {
	platform := "xhs-pc-web"
	hash := md5Hash(params)
	x4 := ""
	if strings.Contains(params, "{") {
		x4 = "object"
	}
	xys := &xys{
		X0: "4.2.1",
		X1: platform,
		X2: "Mac OS",
		X3: mns0101(hash, a1, platform, params),
		X4: x4,
	}
	js, err := json.Marshal(xys)
	if err != nil {
		return ""
	}
	ret := encryptEncodeUtf8(string(js))
	return "XYS_" + commonBase64.EncodeToString(ret)
}

func mns0101(hash, a1, platform, params string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	randNum := rand.Uint32()
	ts := time.Now().UnixMilli()
	startupTs := ts - int64(rand.Intn(4001)+1000)

	x3 := buildX3(randNum, ts, startupTs, hash, a1, platform, params)
	data := xor(x3, 858975407)
	str := encode(data)
	return "mns0101_" + str
}

func deMns0101() (any, error) {
	data, err := decode("Q2vPHt2SEFYRY9PNnW4CKIBdBhB10QpONKp6TmX7T4h3BBa3L6z087o/xLbHlJe4aY7Ruo4PyIP3yG9cHHbKJK75hwDx+pe2dlf8kZJR5TP0+zwKZ3G0WDGtI5xKiCMlbZ7avy1TTYbHgeguZ6wWoovZp52nWmL72xuStd4BJ7")
	if err != nil {
		return nil, err
	}
	data = xor(data, 858975407)
	return nil, nil
}

func buildX3(randNum uint32, ts int64, startupTs int64, hash string, a1, platform, params string) []byte {
	arr := make([]byte, 0, 124)
	// 固定头 4 字节
	arr = append(arr, 119, 104, 96, 41)

	randData := binary.LittleEndian.AppendUint32([]byte{}, randNum)
	arr = append(arr, randData...)

	// 时间戳
	arr = append(arr, encodeTimestamp(ts, true)...)

	// 启动时间戳
	arr = binary.LittleEndian.AppendUint64(arr, uint64(startupTs))

	// 固定整数 4 1269
	arr = binary.LittleEndian.AppendUint32(arr, 4)
	arr = binary.LittleEndian.AppendUint32(arr, 1269)

	arr = binary.LittleEndian.AppendUint32(arr, uint32(len(params)))

	arr = append(arr, hashXor(hash, randData[0])...)

	// a1 和 platform
	arr = append(arr, bytesPrefixLen(a1)...)
	arr = append(arr, bytesPrefixLen(platform)...)

	// 固定
	tail := []byte{
		1, byte(rand.Intn(256)),
		249, 83, 102, 103, 201, 181, 128, 99, 94, 7, 68, 250, 132, 21,
	}
	arr = append(arr, tail...)
	return arr
}

func encodeTimestamp(ts int64, randomizeFirst bool) []byte {
	key := []byte{41, 41, 41, 41, 41, 41, 41, 41}
	arr := make([]byte, 8)
	binary.LittleEndian.PutUint64(arr, uint64(ts))
	encoded := make([]byte, 8)
	for i := 0; i < 8; i++ {
		encoded[i] = arr[i] ^ key[i]
	}
	if randomizeFirst {
		encoded[0] = byte(rand.Intn(256))
	}
	return encoded
}

func hashXor(hash string, xorKey byte) []byte {
	data, err := hex.DecodeString(hash)
	if err != nil {
		return nil
	}
	ret := make([]byte, len(data))
	for i, b := range data {
		ret[i] = b ^ xorKey
	}
	return ret[:8]
}

func bytesPrefixLen(s string) []byte {
	data := []byte(s)
	if len(data) > 255 {
		panic("string too long for 1-byte length prefix")
	}
	return append([]byte{byte(len(data))}, data...)
}

func computeValue(seed uint32) uint32 {
	s15 := seed >> 15
	s13 := seed >> 13
	s12 := seed >> 12
	s10 := seed >> 10
	xorPart := (s15 &^ s13) | (s13 &^ s15)
	return (xorPart ^ s12 ^ s10) << 31
}

func xor(arr []byte, seed uint32) []byte {
	res := make([]byte, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i] ^ byte(seed&0xFF)
		seed = computeValue(seed) | (seed >> 1)
	}
	return res
}

func md5Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	v := h.Sum(nil)
	return hex.EncodeToString(v)
}

const alphabet = "NOPQRStuvwxWXYZabcyz012DEFTKLMdefghijkl4563GHIJBC7mnop89+/"

var base = big.NewInt(58)

func encode(data []byte) string {
	num := new(big.Int).SetBytes(data)
	if num.Cmp(big.NewInt(0)) == 0 {
		return string(alphabet[0])
	}
	var result []byte
	mod := new(big.Int)
	for num.Cmp(big.NewInt(0)) > 0 {
		num.QuoRem(num, base, mod)
		result = append(result, alphabet[mod.Int64()])
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func decode(s string) ([]byte, error) {
	num := big.NewInt(0)
	for _, c := range s {
		index := strings.IndexRune(alphabet, c)
		if index == -1 {
			return nil, fmt.Errorf("invalid character: %c", c)
		}
		num.Mul(num, base)
		num.Add(num, big.NewInt(int64(index)))
	}
	return num.Bytes(), nil
}

func encryptEncodeUtf8(input string) []byte {
	encodeUri := url.PathEscape(input)
	var output []byte
	for i := 0; i < len(encodeUri); i++ {
		char := string(encodeUri[i])
		if char == "%" {
			hexx := encodeUri[i+1 : i+3]
			decimal, _ := strconv.ParseInt(hexx, 16, 0)
			output = append(output, byte(decimal&0xFF))
			i += 2
		} else {
			output = append(output, encodeUri[i])
		}
	}
	return output
}
