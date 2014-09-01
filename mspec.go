package gomspec

import (
	"strings"
)

var mspec *Mspec

var (
	// http://wiki.bash-hackers.org/scripting/terminalcodes

	reset            = "\033[0m"
	bold             = "\033[1m"
	dim              = "\033[2m"
	underline        = "\033[4m"
	underlineOff     = "\033[24m"
	blink            = "\033[5m"
	blinkOff         = "\033[25m"
	inverse          = "\033[7m"
	inverseOff       = "\033[27m"
	strikethrough    = "\033[9m"
	strikethroughOff = "\033[29m"

	black        = "\033[0;30m"
	darkGray     = "\033[1;30m"
	red          = "\033[0;31m"
	lightRed     = "\033[1;31m"
	green        = "\033[0;32m"
	lightGreen   = "\033[1;32m"
	yellow       = "\033[0;33m"
	lightYellow  = "\033[1;33m"
	blue         = "\033[0;34m"
	lightBlue    = "\033[1;34m"
	magenta      = "\033[0;35m"
	lightMagenta = "\033[1;35m"
	cyan         = "\033[0;36m"
	lightCyan    = "\033[1;36m"
	grey         = "\033[0;37m"
	white        = "\033[1;37m"
	defaultColor = "\033[39m"

	blackBg        = "\033[40m"
	regBg          = "\033[41m"
	greenBg        = "\033[42m"
	yellowBg       = "\033[43m"
	blueBg         = "\033[44m"
	magentaBg      = "\033[45m"
	cyanBg         = "\033[46m"
	whiteBg        = "\033[47m"
	defaultBgColor = "\033[49m"

	darkGrey = "\x1B[90m"
)

// Mspec defines the configurations and registrations for package.
type Mspec struct {
	AnsiOfFeature            string
	AnsiOfGiven              string
	AnsiOfWhen               string
	AnsiOfThen               string
	AnsiOfThenNotImplemented string
	AnsiOfThenWithError      string
	AnsiOfCode               string
	AnsiOfCodeError          string
	AnsiOfExpectedError      string
}

func init() {
	mspec = &Mspec{
		AnsiOfFeature:            strings.Join([]string{white}, ""),
		AnsiOfGiven:              strings.Join([]string{grey}, ""),
		AnsiOfWhen:               strings.Join([]string{lightGreen}, ""),
		AnsiOfThen:               strings.Join([]string{green}, ""),
		AnsiOfThenNotImplemented: strings.Join([]string{lightYellow}, ""),
		AnsiOfThenWithError:      strings.Join([]string{regBg, white, bold}, ""),
		AnsiOfCode:               strings.Join([]string{darkGrey}, ""),
		AnsiOfCodeError:          strings.Join([]string{white, bold}, ""),
		AnsiOfExpectedError:      strings.Join([]string{red}, ""),
	}
}
