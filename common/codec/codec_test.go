package codec

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
)

// var (
// 	veryBigInt = new(big.Int).Add(
// 		big.NewInt(0).Lsh(big.NewInt(0xFFFFFFFFFFFFFF), 16),
// 		big.NewInt(0xFFFF),
// 	)
// 	veryVeryBigInt = new(big.Int).Exp(veryBigInt, big.NewInt(8), nil)
// )

type namedByteType byte
type RawValue []byte
type simplestruct struct {
	A uint
	B string
}

func unhex(str string) []byte {
	b, err := hex.DecodeString(strings.Replace(str, " ", "", -1))
	if err != nil {
		panic(fmt.Sprintf("invalid hex string: %q", str))
	}
	return b
}

type encTest struct {
	val           interface{}
	output, error string
}

var encTests = []encTest{
	// booleans
	{val: true, output: "01"},
	{val: false, output: "80"},

	// // integers
	// {val: uint32(0), output: "80"},
	// {val: uint32(127), output: "7F"},
	// {val: uint32(128), output: "8180"},
	// {val: uint32(256), output: "820100"},
	// {val: uint32(1024), output: "820400"},
	// {val: uint32(0xFFFFFF), output: "83FFFFFF"},
	// {val: uint32(0xFFFFFFFF), output: "84FFFFFFFF"},
	// {val: uint64(0xFFFFFFFF), output: "84FFFFFFFF"},
	// {val: uint64(0xFFFFFFFFFF), output: "85FFFFFFFFFF"},
	// {val: uint64(0xFFFFFFFFFFFF), output: "86FFFFFFFFFFFF"},
	// {val: uint64(0xFFFFFFFFFFFFFF), output: "87FFFFFFFFFFFFFF"},
	// {val: uint64(0xFFFFFFFFFFFFFFFF), output: "88FFFFFFFFFFFFFFFF"},

	// // // big integers (should match uint for small values)
	// // {val: big.NewInt(0), output: "80"},
	// // {val: big.NewInt(1), output: "01"},
	// // {val: big.NewInt(127), output: "7F"},
	// // {val: big.NewInt(128), output: "8180"},
	// // {val: big.NewInt(256), output: "820100"},
	// // {val: big.NewInt(1024), output: "820400"},
	// // {val: big.NewInt(0xFFFFFF), output: "83FFFFFF"},
	// // {val: big.NewInt(0xFFFFFFFF), output: "84FFFFFFFF"},
	// // {val: big.NewInt(0xFFFFFFFFFF), output: "85FFFFFFFFFF"},
	// // {val: big.NewInt(0xFFFFFFFFFFFF), output: "86FFFFFFFFFFFF"},
	// // {val: big.NewInt(0xFFFFFFFFFFFFFF), output: "87FFFFFFFFFFFFFF"},
	// // {
	// // 	val:    big.NewInt(0).SetBytes(unhex("102030405060708090A0B0C0D0E0F2")),
	// // 	output: "8F102030405060708090A0B0C0D0E0F2",
	// // },
	// // {
	// // 	val:    big.NewInt(0).SetBytes(unhex("0100020003000400050006000700080009000A000B000C000D000E01")),
	// // 	output: "9C0100020003000400050006000700080009000A000B000C000D000E01",
	// // },
	// // {
	// // 	val:    big.NewInt(0).SetBytes(unhex("010000000000000000000000000000000000000000000000000000000000000000")),
	// // 	output: "A1010000000000000000000000000000000000000000000000000000000000000000",
	// // },
	// // {
	// // 	val:    veryBigInt,
	// // 	output: "89FFFFFFFFFFFFFFFFFF",
	// // },
	// // {
	// // 	val:    veryVeryBigInt,
	// // 	output: "B848FFFFFFFFFFFFFFFFF800000000000000001BFFFFFFFFFFFFFFFFC8000000000000000045FFFFFFFFFFFFFFFFC800000000000000001BFFFFFFFFFFFFFFFFF8000000000000000001",
	// // },

	// // // non-pointer big.Int
	// // {val: *big.NewInt(0), output: "80"},
	// // {val: *big.NewInt(0xFFFFFF), output: "83FFFFFF"},

	// // // negative ints are not supported
	// // {val: big.NewInt(-1), error: "rlp: cannot encode negative *big.Int"},

	// // byte arrays
	// {val: [0]byte{}, output: "80"},
	// {val: [1]byte{0}, output: "00"},
	// {val: [1]byte{1}, output: "01"},
	// {val: [1]byte{0x7F}, output: "7F"},
	// {val: [1]byte{0x80}, output: "8180"},
	// {val: [1]byte{0xFF}, output: "81FF"},
	// {val: [3]byte{1, 2, 3}, output: "83010203"},
	// {val: [57]byte{1, 2, 3}, output: "B839010203000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},

	// // named byte type arrays
	// {val: [0]namedByteType{}, output: "80"},
	// {val: [1]namedByteType{0}, output: "00"},
	// {val: [1]namedByteType{1}, output: "01"},
	// {val: [1]namedByteType{0x7F}, output: "7F"},
	// {val: [1]namedByteType{0x80}, output: "8180"},
	// {val: [1]namedByteType{0xFF}, output: "81FF"},
	// {val: [3]namedByteType{1, 2, 3}, output: "83010203"},
	// {val: [57]namedByteType{1, 2, 3}, output: "B839010203000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},

	// // byte slices
	// {val: []byte{}, output: "80"},
	// {val: []byte{0}, output: "00"},
	// {val: []byte{0x7E}, output: "7E"},
	// {val: []byte{0x7F}, output: "7F"},
	// {val: []byte{0x80}, output: "8180"},
	// {val: []byte{1, 2, 3}, output: "83010203"},

	// // named byte type slices
	// {val: []namedByteType{}, output: "80"},
	// {val: []namedByteType{0}, output: "00"},
	// {val: []namedByteType{0x7E}, output: "7E"},
	// {val: []namedByteType{0x7F}, output: "7F"},
	// {val: []namedByteType{0x80}, output: "8180"},
	// {val: []namedByteType{1, 2, 3}, output: "83010203"},

	// // strings
	// {val: "", output: "80"},
	// {val: "\x7E", output: "7E"},
	// {val: "\x7F", output: "7F"},
	// {val: "\x80", output: "8180"},
	// {val: "dog", output: "83646F67"},
	// {
	// 	val:    "Lorem ipsum dolor sit amet, consectetur adipisicing eli",
	// 	output: "B74C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E7365637465747572206164697069736963696E6720656C69",
	// },
	// {
	// 	val:    "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
	// 	output: "B8384C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E7365637465747572206164697069736963696E6720656C6974",
	// },
	// {
	// 	val:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur mauris magna, suscipit sed vehicula non, iaculis faucibus tortor. Proin suscipit ultricies malesuada. Duis tortor elit, dictum quis tristique eu, ultrices at risus. Morbi a est imperdiet mi ullamcorper aliquet suscipit nec lorem. Aenean quis leo mollis, vulputate elit varius, consequat enim. Nulla ultrices turpis justo, et posuere urna consectetur nec. Proin non convallis metus. Donec tempor ipsum in mauris congue sollicitudin. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Suspendisse convallis sem vel massa faucibus, eget lacinia lacus tempor. Nulla quis ultricies purus. Proin auctor rhoncus nibh condimentum mollis. Aliquam consequat enim at metus luctus, a eleifend purus egestas. Curabitur at nibh metus. Nam bibendum, neque at auctor tristique, lorem libero aliquet arcu, non interdum tellus lectus sit amet eros. Cras rhoncus, metus ac ornare cursus, dolor justo ultrices metus, at ullamcorper volutpat",
	// 	output: "B904004C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E73656374657475722061646970697363696E6720656C69742E20437572616269747572206D6175726973206D61676E612C20737573636970697420736564207665686963756C61206E6F6E2C20696163756C697320666175636962757320746F72746F722E2050726F696E20737573636970697420756C74726963696573206D616C6573756164612E204475697320746F72746F7220656C69742C2064696374756D2071756973207472697374697175652065752C20756C7472696365732061742072697375732E204D6F72626920612065737420696D70657264696574206D6920756C6C616D636F7270657220616C6971756574207375736369706974206E6563206C6F72656D2E2041656E65616E2071756973206C656F206D6F6C6C69732C2076756C70757461746520656C6974207661726975732C20636F6E73657175617420656E696D2E204E756C6C6120756C74726963657320747572706973206A7573746F2C20657420706F73756572652075726E6120636F6E7365637465747572206E65632E2050726F696E206E6F6E20636F6E76616C6C6973206D657475732E20446F6E65632074656D706F7220697073756D20696E206D617572697320636F6E67756520736F6C6C696369747564696E2E20566573746962756C756D20616E746520697073756D207072696D697320696E206661756369627573206F726369206C756374757320657420756C74726963657320706F737565726520637562696C69612043757261653B2053757370656E646973736520636F6E76616C6C69732073656D2076656C206D617373612066617563696275732C2065676574206C6163696E6961206C616375732074656D706F722E204E756C6C61207175697320756C747269636965732070757275732E2050726F696E20617563746F722072686F6E637573206E69626820636F6E64696D656E74756D206D6F6C6C69732E20416C697175616D20636F6E73657175617420656E696D206174206D65747573206C75637475732C206120656C656966656E6420707572757320656765737461732E20437572616269747572206174206E696268206D657475732E204E616D20626962656E64756D2C206E6571756520617420617563746F72207472697374697175652C206C6F72656D206C696265726F20616C697175657420617263752C206E6F6E20696E74657264756D2074656C6C7573206C65637475732073697420616D65742065726F732E20437261732072686F6E6375732C206D65747573206163206F726E617265206375727375732C20646F6C6F72206A7573746F20756C747269636573206D657475732C20617420756C6C616D636F7270657220766F6C7574706174",
	// },

	// // slices
	// {val: []uint{}, output: "C0"},
	// {val: []uint{1, 2, 3}, output: "C3010203"},
	// {
	// 	// [ [], [[]], [ [], [[]] ] ]
	// 	val:    []interface{}{[]interface{}{}, [][]interface{}{{}}, []interface{}{[]interface{}{}, [][]interface{}{{}}}},
	// 	output: "C7C0C1C0C3C0C1C0",
	// },
	// {
	// 	val:    []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii", "jjj", "kkk", "lll", "mmm", "nnn", "ooo"},
	// 	output: "F83C836161618362626283636363836464648365656583666666836767678368686883696969836A6A6A836B6B6B836C6C6C836D6D6D836E6E6E836F6F6F",
	// },
	// {
	// 	val:    []interface{}{uint(1), uint(0xFFFFFF), []interface{}{[]uint{4, 5, 5}}, "abc"},
	// 	output: "CE0183FFFFFFC4C304050583616263",
	// },
	// {
	// 	val: [][]string{
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 		{"asdf", "qwer", "zxcv"},
	// 	},
	// 	output: "F90200CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376CF84617364668471776572847A786376",
	// },

	// // RawValue
	// {val: RawValue(unhex("01")), output: "01"},
	// {val: RawValue(unhex("82FFFF")), output: "8382FFFF"},
	// {val: []RawValue{unhex("01"), unhex("02")}, output: "C20102"},

	// // structs
	// {val: simplestruct{}, output: "C28080"},
	// {val: simplestruct{A: 3, B: "foo"}, output: "C50383666F6F"},

	// nil
	{val: (*uint)(nil), output: "80"},
	{val: (*string)(nil), output: "80"},
	{val: (*[]byte)(nil), output: "80"},
	{val: (*[10]byte)(nil), output: "80"},
	{val: (*big.Int)(nil), output: "80"},
	{val: (*[]string)(nil), output: "C0"},
	{val: (*[10]string)(nil), output: "C0"},
	{val: (*[]interface{})(nil), output: "C0"},
	{val: (*[]struct{ uint })(nil), output: "C0"},
	{val: (*interface{})(nil), output: "C0"},

	// nil struct fields
	{
		val: struct {
			X *[]byte
		}{},
		output: "C180",
	},
	{
		val: struct {
			X *[2]byte
		}{},
		output: "C180",
	},
	{
		val: struct {
			X *uint64
		}{},
		output: "C180",
	},
	{
		val: struct {
			X *[]uint64
		}{},
		output: "C1C0",
	},
}

type TestCase struct {
	In  interface{} `json:"in"`
	Out string      `json:"out"`
}

type TestCases map[string]TestCase

func runEncTests(t *testing.T, f func(val interface{}) ([]byte, error)) {
	for i, test := range encTests {
		output, err := f(test.val)
		if err != nil && test.error == "" {
			t.Errorf("test %d: unexpected error: %v\nvalue %#v\ntype %T",
				i, err, test.val, test.val)
			continue
		}
		if test.error != "" && fmt.Sprint(err) != test.error {
			t.Errorf("test %d: error mismatch\ngot   %v\nwant  %v\nvalue %#v\ntype  %T",
				i, err, test.error, test.val, test.val)
			continue
		}
		if err == nil && !bytes.Equal(output, unhex(test.output)) {
			t.Errorf("test %d: output mismatch:\ngot   %X\nwant  %s\nvalue %#v\ntype  %T",
				i, output, test.output, test.val, test.val)
		}
	}
}

func TestEncodeToBytes(t *testing.T) {
	runEncTests(t, RLP.MarshalToBytes)
}
