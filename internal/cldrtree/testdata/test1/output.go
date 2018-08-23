// Code generated by running "go generate" in github.com/sunnogo/text. DO NOT EDIT.

package test

import "github.com/sunnogo/text/internal/cldrtree"

var tree = &cldrtree.Tree{locales, indices, buckets}

// Path values:
// <context>
//   - format
//   - stand-alone
// <width>
//   - wAbbreviated
//   - wNarrow
//   - wWide
// <month>
//   - leap7
//   - 1..12
//
// - calendars
//   - chinese
//   - dangi
//   - gregorian
//     - months
//       - <context>
//         - <width>
//           - <month>
//     - filler
//       - 0
//
// Nr elem:            39
// uniqued size:       912
// total string size:  912
// bucket waste:       0

// context specifies a property of a CLDR field.
type context uint16

// width specifies a property of a CLDR field.
type width uint16

// month specifies a property of a CLDR field.
type month uint16

const (
	calendars            = 0 // calendars
	chinese              = 0 // chinese
	dangi                = 1 // dangi
	gregorian            = 2 // gregorian
	months               = 0 // months
	filler               = 1 // filler
	format       context = 0 // format
	standAlone   context = 1 // stand-alone
	wAbbreviated width   = 0 // wAbbreviated
	wNarrow      width   = 1 // wNarrow
	wWide        width   = 2 // wWide
	leap7        month   = 0 // leap7
)

var locales = []uint32{ // 775 elements
	// Entry 0 - 1F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 20 - 3F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 40 - 5F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 60 - 7F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 80 - 9F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry A0 - BF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry C0 - DF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry E0 - FF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 100 - 11F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 120 - 13F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 140 - 15F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 160 - 17F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 180 - 19F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 1A0 - 1BF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 1C0 - 1DF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 1E0 - 1FF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 200 - 21F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 220 - 23F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 240 - 25F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 260 - 27F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 280 - 29F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 2A0 - 2BF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 2C0 - 2DF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 2E0 - 2FF
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 300 - 31F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000,
} // Size: xxxx bytes

var indices = []uint16{ // 86 elements
	// Entry 0 - 3F
	0x0001, 0x0002, 0x0003, 0x0006, 0x0021, 0x0027, 0x0002, 0x0009,
	0x001e, 0x0001, 0x000b, 0x0003, 0x8002, 0x8002, 0x000f, 0x000d,
	0x0000, 0xffff, 0x0000, 0x0005, 0x000a, 0x000f, 0x0014, 0x0019,
	0x001e, 0x0023, 0x0028, 0x002d, 0x0032, 0x0037, 0x0001, 0x0000,
	0x003c, 0x0002, 0x9000, 0x0024, 0x0001, 0x0000, 0x013b, 0x0002,
	0x002a, 0x0053, 0x0002, 0x002d, 0x0040, 0x0003, 0x8002, 0x9001,
	0x0031, 0x000d, 0x0000, 0xffff, 0x023a, 0x023f, 0x0244, 0x0249,
	0x024e, 0x0253, 0x0258, 0x025d, 0x0262, 0x0267, 0x026c, 0x0271,
	// Entry 40 - 7F
	0x0003, 0x9000, 0x0044, 0x9000, 0x000d, 0x0000, 0xffff, 0x0276,
	0x0278, 0x027a, 0x027c, 0x027e, 0x0280, 0x0282, 0x0284, 0x0286,
	0x0288, 0x028b, 0x028e, 0x0001, 0x0000, 0x0291,
} // Size: xxxx bytes

var buckets = []string{
	bucket0,
}

var bucket0 string = "" + // Size: xxxx bytes
	"\x04cM01\x04cM02\x04cM03\x04cM04\x04cM05\x04cM06\x04cM07\x04cM08\x04cM09" +
	"\x04cM10\x04cM11\x04cM12\xfe\x01\x94\xfd\xc2\xfa/\xfc\xc0A\xd3\xff\x12" +
	"\x04[s\xc8nO\xf9_\xf6b\xa5\xee\xe8*\xbd\xf4J-\x0bu\xfb\x18\x0d\xafH\xa7" +
	"\x9e\xe0\xb1\x0d9FQ\x85\x0fԡx\x89.\xe2\x85\xec\xe1Q\x14Ux\x08u\xd6N\xe2" +
	"\xd3\xd0\xd0\xdek\xf8\xf9\xb4L\xe8_\xf0DƱ\xf8;\x8e\x88;\xbf\x85z\xab\x99" +
	"ŲR\xc7B\x9c2\U000e8bb7\x9e\xf8V\xf6Y\xc1\x8f\x0d\xce\xccw\xc7^z\x81\xbf" +
	"\xde'_g\xcf\xe2B\xcf<\xc3T\xf3\xed\xe2־\xccN\xa3\xae^\x88Rj\x9fJW\x8b˞" +
	"\xf2ԦS\x14v\x8dm)\x97a\xea\x9eOZ\xa6\xae\xc3\xfcxƪ\xe0\x81\xac\x81 \xc7 " +
	"\xef\xcdlꄶ\x92^`{\xe0cqo\x96\xdd\xcd\xd0\x1du\x04\\?\x00\x0f\x8ayk\xcelQ" +
	",8\x01\xaa\xca\xee߭[Pfd\xe8\xc0\xe4\xa7q\xecื\xc1\x96]\x91\x81%\x1b|\x9c" +
	"\x9c\xa5 Z\xfc\x16\xa26\xa2\xef\xcd\xd2\xd1-*y\xd0\xfet\xa8(\x0a\xe9C" +
	"\x9e\xb0֮\xca\x08#\xae\x02\xd6}\x86j\xc2\xc4\xfeJrPS\xda\x11\x9b\x9dOQQ@" +
	"\xa2\xd7#\x9c@\xb4ZÕ\x0d\x94\x1f\xc4\xfe\x1c\x0c\xb9j\xd3\x22\xd6\x22" +
	"\x82)_\xbf\xe1\x1e&\xa43\x07m\xb5\xc1DL:4\xd3*\\J\x7f\xfb\xe8с\xf7\xed;" +
	"\x8c\xfe\x90O\x93\xf8\xf0m)\xbc\xd9\xed\x84{\x18.\x04d\x10\xf4Kİ\xf3\xf0" +
	":\x0d\x06\x82\x0a0\xf2W\xf8\x11A0g\x8a\xc0E\x86\xc1\xe3\xc94,\x8b\x80U" +
	"\xc4f؆D\x1d%\x99\x06͉֚K\x96\x8a\xe9\xf0띖\\\xe6\xa4i<N\xbe\x88\x15\x01" +
	"\xb7لkf\xeb\x02\xb5~\\\xda{l\xbah\x91\xd6\x16\xbdhl7\xb84a:Ⱥ\xa2,\x00" +
	"\x8f\xfeh\x83RsJ\xe4\xe3\xf1!z\xcd_\x83'\x08\x140\x18g\xb5\xd0g\x11\xb28" +
	"\x00\x1cyW\xb2w\x19\xce?1\x88\xdf\xe5}\xee\xbfo\x82YZ\x10\xf7\xbbV,\xa0M" +
	"\\='\x04gM01\x04gM02\x04gM03\x04gM04\x04gM05\x04gM06\x04gM07\x04gM08\x04" +
	"gM09\x04gM10\x04gM11\x04gM12\x011\x012\x013\x014\x015\x016\x017\x018\x01" +
	"9\x0210\x0211\x0212\xfe\x94)X\xc6\xdb2bg\x06I\xf3\xbc\x97٢1g5\xed悥\xdf" +
	"\xe6\xf1\xa0\x11\xfbɊ\xd0\xfb\xe7\x90\x00<\x01\xe8\xe9\x96w\x03\xaff^" +
	"\x9fr@\x7fK\x03\xd4\xfd\xb4t\xaa\xfe\x8a\x0d>\x05\x15\xddFP\xcfQ\x17+" +
	"\x81$\x8b\xcb\x7f\x96\x9e@\x0bl[\x12wh\xb1\xc4\x12\xfa\xe9\x8c\xf5v1\xcf" +
	"7\x03;KJ\xba}~\xd3\x19\xba\x14rI\xc9\x08\xacp\xd1\xc4\x06\xda\xde\x0e" +
	"\x82\x8e\xb6\xba\x0dʨ\x82\x85T>\x10!<d?\xc8`;X`#fp\xba\xbc\xad\x0b\xd7" +
	"\xf4\xc4\x19\x0e26#\xa8h\xd1\xea\xe1v\x9f@\xa2f1C\x1b;\xd5!V\x05\xd2\x08" +
	"o\xeaԙ\xacc\xa4e=\x12(=V\x01\x9c7\x95\xa9\x8a\x12m\x09\xcf\xcb\xe3l\xdc" +
	"\xc97\x88\xa5@\x9f\x8bnB\xc2݃\xaaFa\x18R\xad\x0bP(w\\w\x16\x90\xb6\x85N" +
	"\x05\xb3w$\x1es\xa8\x83\xddw\xaf\xf00,m\xa8f\\B4\x1d\xdaJ\xda\xea"

var enumMap = map[string]uint16{
	"":             0,
	"calendars":    0,
	"chinese":      0,
	"dangi":        1,
	"gregorian":    2,
	"months":       0,
	"filler":       1,
	"format":       0,
	"stand-alone":  1,
	"wAbbreviated": 0,
	"wNarrow":      1,
	"wWide":        2,
	"leap7":        0,
}

// Total table size: xxxx bytes (4KiB); checksum: 7EA7DE6
