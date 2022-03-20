package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase5 struct {
	Input          string
	PossibleResult []string
}

func Test5(t *testing.T) {
	testcases := []Testcase5{
		{
			Input:          "a",
			PossibleResult: []string{"a"},
		}, {
			Input:          "ab",
			PossibleResult: []string{"a", "b"},
		}, {
			Input:          "babad",
			PossibleResult: []string{"bab", "aba"},
		},  {
			Input:          "bababbababbabab",
			PossibleResult: []string{"bababbababbabab"},
		}, {
			Input:          "cbbd",
			PossibleResult: []string{"bb"},
		}, {
			Input:          "ccc",
			PossibleResult: []string{"ccc"},
		}, {
			Input:          "zbkksfgesmfyuedjxdtknclymgskfjduhfocipzjqnmvcodjlvlagmhokqfeudickyeoobmkerjdeloxfbauryanltprloaeboavxzltgcurgbtgtpygpjizoopwmwjixaowppdvkferupefhwombszifyliidrxpxgcpbfzqtxdfiwfmtgvjiccrigwljrlvhaegvbitngckdnsfcnqlgykwjmsifcttcbeummaoidrrhkxmxctugcrlpbiolzqnjtwhzreruglrdvzioewcgvjjwfyqmhupusktptvfpcqxkvpbrlzchtacmlzgeejnvjzzhcegwtwqhimwooflzeiomeqyrnaeiquolmsunqrgffkpljewyritkivdrfnovbatdstypzsmbjdrromcqexnmjcuqpjzzjpqucjmnxeqcmorrdjbmszpytsdtabvonfrdviktirywejlpkffgrqnusmlouqieanryqemoiezlfoowmihqwtwgechzzjvnjeegzlmcathczlrbpvkxqcpfvtptksupuhmqyfwjjvgcweoizvdrlgurerzhwtjnqzloibplrcgutcxmxkhrrdioammuebcttcfismjwkyglqncfsndkcgntibvgeahvlrjlwgirccijvgtmfwifdxtqzfbpcgxpxrdiilyfizsbmowhfepurefkvdppwoaxijwmwpoozijpgyptgtbgrucgtlzxvaobeaolrptlnayruabfxoledjrekmbooeykciduefqkohmgalvljdocvmnqjzpicofhudjfksgmylcnktdxjdeuyfmsegfskkbz",
			PossibleResult: []string{"zbkksfgesmfyuedjxdtknclymgskfjduhfocipzjqnmvcodjlvlagmhokqfeudickyeoobmkerjdeloxfbauryanltprloaeboavxzltgcurgbtgtpygpjizoopwmwjixaowppdvkferupefhwombszifyliidrxpxgcpbfzqtxdfiwfmtgvjiccrigwljrlvhaegvbitngckdnsfcnqlgykwjmsifcttcbeummaoidrrhkxmxctugcrlpbiolzqnjtwhzreruglrdvzioewcgvjjwfyqmhupusktptvfpcqxkvpbrlzchtacmlzgeejnvjzzhcegwtwqhimwooflzeiomeqyrnaeiquolmsunqrgffkpljewyritkivdrfnovbatdstypzsmbjdrromcqexnmjcuqpjzzjpqucjmnxeqcmorrdjbmszpytsdtabvonfrdviktirywejlpkffgrqnusmlouqieanryqemoiezlfoowmihqwtwgechzzjvnjeegzlmcathczlrbpvkxqcpfvtptksupuhmqyfwjjvgcweoizvdrlgurerzhwtjnqzloibplrcgutcxmxkhrrdioammuebcttcfismjwkyglqncfsndkcgntibvgeahvlrjlwgirccijvgtmfwifdxtqzfbpcgxpxrdiilyfizsbmowhfepurefkvdppwoaxijwmwpoozijpgyptgtbgrucgtlzxvaobeaolrptlnayruabfxoledjrekmbooeykciduefqkohmgalvljdocvmnqjzpicofhudjfksgmylcnktdxjdeuyfmsegfskkbz"},
		},
	}

	for _, testcase := range testcases {
		result := longestPalindrome(testcase.Input)
		assert.Contains(t, testcase.PossibleResult, result)
	}
}

type IsSubstringPalindromeTestcase struct {
	Str        string
	LeftIndex  int
	RightIndex int
	Expected   bool
}

func TestIsSubstringPalindrome(t *testing.T) {
	testcases := []IsSubstringPalindromeTestcase{
		{
			Str:        "a",
			LeftIndex:  0,
			RightIndex: 0,
			Expected:   true,
		}, {
			Str:        "abcdcdz",
			LeftIndex:  0,
			RightIndex: 3,
			Expected:   false,
		}, {
			Str:        "abcdcbz",
			LeftIndex:  1,
			RightIndex: 5,
			Expected:   true,
		}, {
			Str:        "cbbd",
			LeftIndex:  1,
			RightIndex: 2,
			Expected:   true,
		}, {
			Str:        "ccc",
			LeftIndex:  0,
			RightIndex: 2,
			Expected:   true,
		},
	}
	for _, testcase := range testcases {
		checker := newPalindromeChecker()
		actual := checker.isSubstringPalindrome(testcase.Str, testcase.LeftIndex, testcase.RightIndex)
		assert.Equalf(t, testcase.Expected, actual, "%s %d %d", testcase.Str, testcase.LeftIndex, testcase.RightIndex)
	}
}
