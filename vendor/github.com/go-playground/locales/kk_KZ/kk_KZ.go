package kk_KZ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type kk_KZ struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'kk_KZ' locale
func New() locales.Translator {
	return &kk_KZ{
		locale:                 "kk_KZ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "қаң.", "ақп.", "нау.", "сәу.", "мам.", "мау.", "шіл.", "там.", "қыр.", "қаз.", "қар.", "жел."},
		monthsNarrow:           []string{"", "Қ", "А", "Н", "С", "М", "М", "Ш", "Т", "Қ", "Қ", "Қ", "Ж"},
		monthsWide:             []string{"", "қаңтар", "ақпан", "наурыз", "сәуір", "мамыр", "маусым", "шілде", "тамыз", "қыркүйек", "қазан", "қараша", "желтоқсан"},
		daysAbbreviated:        []string{"Жс", "Дс", "Сс", "Ср", "Бс", "Жм", "Сб"},
		daysNarrow:             []string{"Ж", "Д", "С", "С", "Б", "Ж", "С"},
		daysShort:              []string{"Жс", "Дс", "Сс", "Ср", "Бс", "Жм", "Сб"},
		daysWide:               []string{"жексенбі", "дүйсенбі", "сейсенбі", "сәрсенбі", "бейсенбі", "жұма", "сенбі"},
		periodsAbbreviated:     []string{"таңғы", "түскі/кешкі"},
		periodsNarrow:          []string{"таңғы", "түскі/кешкі"},
		periodsWide:            []string{"таңғы", "түскі/кешкі"},
		erasAbbreviated:        []string{"б.з.д.", "б.з."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Біздің заманымызға дейін", "Біздің заманымыз"},
		timezones:              map[string]string{"CST": "Солтүстік Америка стандартты орталық уақыты", "CDT": "Солтүстік Америка жазғы орталық уақыты", "CLT": "Чили стандартты уақыты", "WIT": "Шығыс Индонезия уақыты", "COT": "Колумбия стандартты уақыты", "ECT": "Эквадор уақыты", "IST": "Үндістан стандартты уақыты", "ADT": "Атлантика жазғы уақыты", "WIB": "Батыс Индонезия уақыты", "AEDT": "Австралия жазғы шығыс уақыты", "MDT": "Солтүстік Америка жазғы тау уақыты", "GMT": "Гринвич уақыты", "OESZ": "Шығыс Еуропа жазғы уақыты", "VET": "Венесуэла уақыты", "WART": "Батыс Аргентина стандартты уақыты", "HADT": "Гавай және Алеут аралдары жазғы уақыты", "MYT": "Малайзия уақыты", "WAST": "Батыс Африка жазғы уақыты", "AKST": "Аляска стандартты уақыты", "∅∅∅": "Бразилия жазғы уақыты", "LHDT": "Лорд-Хау жазғы уақыты", "COST": "Колумбия жазғы уақыты", "AEST": "Австралия стандартты шығыс уақыты", "ACST": "Австралия стандартты орталық уақыты", "ARST": "Аргентина жазғы уақыты", "WAT": "Батыс Африка стандартты уақыты", "EST": "Солтүстік Америка стандартты шығыс уақыты", "HKST": "Гонконг жазғы уақыты", "CHADT": "Чатем жазғы уақыты", "SRT": "Суринам уақыты", "OEZ": "Шығыс Еуропа стандартты уақыты", "TMT": "Түрікменстан стандартты уақыты", "BOT": "Боливия уақыты", "NZDT": "Жаңа Зеландия жазғы уақыты", "BT": "Бутан уақыты", "UYST": "Уругвай жазғы уақыты", "JDT": "Жапония жазғы уақыты", "UYT": "Уругвай стандартты уақыты", "JST": "Жапония стандартты уақыты", "MESZ": "Орталық Еуропа жазғы уақыты", "EDT": "Солтүстік Америка жазғы шығыс уақыты", "CAT": "Орталық Африка уақыты", "LHST": "Лорд-Хау стандартты уақыты", "GFT": "Француз Гвианасы уақыты", "ACDT": "Австралия жазғы орталық уақыты", "NZST": "Жаңа Зеландия стандартты уақыты", "WARST": "Батыс Аргентина жазғы уақыты", "AWDT": "Австралия жазғы батыс уақыты", "CLST": "Чили жазғы уақыты", "TMST": "Түрікменстан жазғы уақыты", "AWST": "Австралия стандартты батыс уақыты", "WESZ": "Батыс Еуропа жазғы уақыты", "SGT": "Сингапур уақыты", "PST": "Солтүстік Америка стандартты Тынық мұхиты уақыты", "EAT": "Шығыс Африка уақыты", "MST": "Солтүстік Америка стандартты тау уақыты", "ChST": "Чаморро стандартты уақыты", "HNT": "Ньюфаундленд стандартты уақыты", "CHAST": "Чатем стандартты уақыты", "ACWDT": "Австралия жазғы орталық-батыс уақыты", "PDT": "Солтүстік Америка жазғы Тынық мұхиты уақыты", "WITA": "Орталық Индонезия уақыты", "WEZ": "Батыс Еуропа стандартты уақыты", "ART": "Аргентина стандартты уақыты", "HAT": "Ньюфаундленд жазғы уақыты", "HKT": "Гонконг стандартты уақыты", "GYT": "Гайана уақыты", "MEZ": "Орталық Еуропа стандартты уақыты", "AKDT": "Аляска жазғы уақыты", "AST": "Атлантика стандартты уақыты", "HAST": "Гавай және Алеут аралдары стандартты уақыты", "ACWST": "Австралия стандартты орталық-батыс уақыты", "SAST": "Оңтүстік Африка уақыты"},
	}
}

// Locale returns the current translators string locale
func (kk *kk_KZ) Locale() string {
	return kk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kk_KZ'
func (kk *kk_KZ) PluralsCardinal() []locales.PluralRule {
	return kk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kk_KZ'
func (kk *kk_KZ) PluralsOrdinal() []locales.PluralRule {
	return kk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kk_KZ'
func (kk *kk_KZ) PluralsRange() []locales.PluralRule {
	return kk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kk_KZ'
func (kk *kk_KZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kk_KZ'
func (kk *kk_KZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)

	if (nMod10 == 6) || (nMod10 == 9) || (nMod10 == 0 && n != 0) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kk_KZ'
func (kk *kk_KZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := kk.CardinalPluralRule(num1, v1)
	end := kk.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kk *kk_KZ) MonthAbbreviated(month time.Month) string {
	return kk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kk *kk_KZ) MonthsAbbreviated() []string {
	return kk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kk *kk_KZ) MonthNarrow(month time.Month) string {
	return kk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kk *kk_KZ) MonthsNarrow() []string {
	return kk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kk *kk_KZ) MonthWide(month time.Month) string {
	return kk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kk *kk_KZ) MonthsWide() []string {
	return kk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kk *kk_KZ) WeekdayAbbreviated(weekday time.Weekday) string {
	return kk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kk *kk_KZ) WeekdaysAbbreviated() []string {
	return kk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kk *kk_KZ) WeekdayNarrow(weekday time.Weekday) string {
	return kk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kk *kk_KZ) WeekdaysNarrow() []string {
	return kk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kk *kk_KZ) WeekdayShort(weekday time.Weekday) string {
	return kk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kk *kk_KZ) WeekdaysShort() []string {
	return kk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kk *kk_KZ) WeekdayWide(weekday time.Weekday) string {
	return kk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kk *kk_KZ) WeekdaysWide() []string {
	return kk.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kk_KZ' and handles both Whole and Real numbers based on 'v'
func (kk *kk_KZ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kk.group) - 1; j >= 0; j-- {
					b = append(b, kk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kk_KZ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kk *kk_KZ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kk_KZ'
func (kk *kk_KZ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kk.group) - 1; j >= 0; j-- {
					b = append(b, kk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, kk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kk_KZ'
// in accounting notation.
func (kk *kk_KZ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kk.group) - 1; j >= 0; j-- {
					b = append(b, kk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, kk.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, kk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, kk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kk.monthsAbbreviated[t.Month()]...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kk.monthsWide[t.Month()]...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20, 0xd0, 0xb6}...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, kk.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kk_KZ'
func (kk *kk_KZ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
