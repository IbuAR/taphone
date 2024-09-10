package taphone

import (
	"testing"
)

type testVal struct {
	word     string
	expected expected
}

type expected struct {
	val1, val2, val3 string
}

func TestTAPhone(t *testing.T) {
	phone := New()
	testStrings := []testVal{
		{
			word: "நிலம்",
			expected: expected{
				"NLM",
				"NLM",
				"N4LM",
			},
		},
		{
			word: "சப்தம்",
			expected: expected{
				"CPTM",
				"CPTM",
				"CPTM",
			},
		},
		{
			word: "பச்சை",
			expected: expected{
				"PCH",
				"PCH",
				"PCH7",
			},
		},
		{
			word: "கலம்",
			expected: expected{
				"KLM",
				"KLM",
				"KLM",
			},
		},
		{
			word: "காலம்",
			expected: expected{
				"KLM",
				"KLM",
				"K3LM",
			},
		}, {
			word: "மச்சம்",
			expected: expected{
				"MCHM",
				"MCHM",
				"MCHM",
			},
		},
		{
			word: "பச்சரிசி",
			expected: expected{
				"PCHRC",
				"PCHRC",
				"PCHR4C4",
			},
		},
		{
			word: "பத்ரகாளி",
			expected: expected{
				"PTRKL",
				"PTRKL1",
				"PTRK3L14",
			},
		},
		{
			word: "கற்க",
			expected: expected{
				"KRK",
				"KR1K",
				"KR1K",
			},
		},
		{
			word: "எண்ணியல்",
			expected: expected{
				"ENNYL",
				"EN1N1YL",
				"EN1N14YL",
			},
		},
	}

	for _, v := range testStrings {
		out1, out2, out3 := phone.Encode(v.word)

		expect(t, v.expected.val1, out1, v.word)
		expect(t, v.expected.val2, out2, v.word)
		expect(t, v.expected.val3, out3, v.word)
	}
}

func expect(t *testing.T, actual, expected, word string) {
	if actual != expected {
		t.Errorf("Expected %v but got %v for %v", expected, actual, word)
	}
}
