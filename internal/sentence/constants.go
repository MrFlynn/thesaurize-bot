package sentence

// List of words that should not be looked up in thesaurus.
var ignoreWords = map[string]bool{
	"the":   true,
	"of":    true,
	"and":   true,
	"a":     true,
	"to":    true,
	"in":    true,
	"is":    true,
	"you":   true,
	"that":  true,
	"it":    true,
	"he":    true,
	"was":   true,
	"for":   true,
	"on":    true,
	"are":   true,
	"as":    true,
	"with":  true,
	"his:":  true,
	"they":  true,
	"at":    true,
	"be":    true,
	"this":  true,
	"have":  true,
	"form":  true,
	"or":    true,
	"one":   true,
	"had":   true,
	"by":    true,
	"but":   true,
	"what":  true,
	"were":  true,
	"we":    true,
	"when":  true,
	"your":  true,
	"can":   true,
	"said":  true,
	"there": true,
	"an":    true,
	"which": true,
	"she":   true,
	"do":    true,
	"how":   true,
	"their": true,
	"if":    true,
	"will":  true,
	"out":   true,
	"then":  true,
	"them":  true,
	"these": true,
	"so":    true,
	"some":  true,
	"her":   true,
	"would": true,
	"him":   true,
	"into":  true,
	"has":   true,
	"no":    true,
	"way":   true,
	"could": true,
	"my":    true,
	"than":  true,
	"been":  true,
	"who":   true,
	"its":   true,
	"now":   true,
	"did":   true,
	"get":   true,
	"come":  true,
	"may":   true,
	"part":  true,
	"i":     true,
	"me":    true,
	"his":   true,
	"us":    true,
}