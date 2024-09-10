package taphone

import (
	"regexp"
	"strings"
)

var vowels = map[string]string{
	"அ": "A", "ஆ": "A", "இ": "I", "ஈ": "I", "உ": "U", "ஊ": "U",
	"எ": "E", "ஏ": "E", "ஐ": "AI", "ஒ": "O", "ஓ": "O", "ஔ": "O",
}

var consonants = map[string]string{
	"க": "K", "ங": "NG",
	"ச": "C", "ஞ": "NJ",
	"ட": "T", "ண": "N1",
	"த": "T", "ந": "N",
	"ப": "P", "ம": "M",
	"ய": "Y", "ர": "R", "ல": "L", "வ": "V",
	"ழ": "Z", "ள": "L1",
	"ற": "R1", "ன": "N2",
	"ஶ": "S", "ஜ": "J", "ஷ": "SH", "ஸ": "S", "ஹ": "H",
}

var compounds = map[string]string{
	"க்க": "K2", "ங்க": "NGK", "ச்ச": "CH",
	"ஞ்ச": "NJC", "ட்ட": "T2", "ண்ட": "N1T",
	"த்த": "T2", "ந்த": "NT", "ப்ப": "P2",
	"ம்ப": "MB", "ய்ய": "YY", "ல்ல": "L2",
	"வ்வ": "VV", "ழ்ழ": "Z2", "ள்ள": "L12",
	"ற்ற": "R2", "ன்ன": "N22",
}

var modifiers = map[string]string{
	"ா": "3", "ி": "4", "ீ": "4", "ு": "5", "ூ": "5",
	"ெ": "6", "ே": "6", "ை": "7", "ொ": "8", "ோ": "8", "ௌ": "9",
	"்": "", "ஃ": "H",
}

var (
	regexKey0, _     = regexp.Compile(`[1-9]`)
	regexKey1, _     = regexp.Compile(`[2-9]`)
	regexNonTamil, _ = regexp.Compile(`[\P{Tamil}]`)
	regexAlphaNum, _ = regexp.Compile(`[^0-9A-Z]`)
)

// TAphone is the Tamil-phone tokenizer.
type TAphone struct {
	modCompounds  *regexp.Regexp
	modConsonants *regexp.Regexp
	modVowels     *regexp.Regexp
}

// New returns a new instance of the TAPhone tokenizer.
func New() *TAphone {
	var (
		glyphs []string
		mods   []string
		tn     = &TAphone{}
	)

	// modifiers.
	for k := range modifiers {
		mods = append(mods, k)
	}

	// compounds.
	for k := range compounds {
		glyphs = append(glyphs, k)
	}
	tn.modCompounds, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	// consonants.
	glyphs = []string{}
	for k := range consonants {
		glyphs = append(glyphs, k)
	}
	tn.modConsonants, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	// vowels.
	glyphs = []string{}
	for k := range vowels {
		glyphs = append(glyphs, k)
	}
	tn.modVowels, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	return tn
}

// Encode encodes a unicode Tamil string to its Roman TAPhone hash.
// Ideally, words should be encoded one at a time, and not as phrases
// or sentences.
func (t *TAphone) Encode(input string) (string, string, string) {
	// key2 accounts for hard and modified sounds.
	key2 := t.process(input)

	// key1 loses numeric modifiers that denote phonetic modifiers.
	key1 := regexKey1.ReplaceAllString(key2, "")

	// key0 loses numeric modifiers that denote hard sounds, doubled sounds,
	// and phonetic modifiers.
	key0 := regexKey0.ReplaceAllString(key2, "")

	return key0, key1, key2
}

func (t *TAphone) process(input string) string {
	// Remove all non-Tamil characters.
	input = regexNonTamil.ReplaceAllString(strings.Trim(input, ""), "")

	// All character replacements are grouped between { and } to maintain
	// separatability till the final step.

	// Replace and group modified compounds.
	input = t.replaceModifiedGlyphs(input, compounds, t.modCompounds)

	// Replace and group unmodified compounds.
	for k, v := range compounds {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace and group modified consonants and vowels.
	input = t.replaceModifiedGlyphs(input, consonants, t.modConsonants)
	input = t.replaceModifiedGlyphs(input, vowels, t.modVowels)

	// Replace and group unmodified consonants.
	for k, v := range consonants {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace and group unmodified vowels.
	for k, v := range vowels {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace all modifiers.
	for k, v := range modifiers {
		input = strings.ReplaceAll(input, k, v)
	}

	// Remove non alpha numeric characters (losing the bracket grouping).
	return regexAlphaNum.ReplaceAllString(input, "")
}

func (t *TAphone) replaceModifiedGlyphs(input string, glyphs map[string]string, r *regexp.Regexp) string {
	for _, matches := range r.FindAllStringSubmatch(input, -1) {
		for _, m := range matches {
			if rep, ok := glyphs[m]; ok {
				input = strings.ReplaceAll(input, m, rep)
			}
		}
	}
	return input
}
