package currency

import (
	"github.com/jasonlvhit/gocron"
	"strings"
)

type Currency struct {
	Code string
	SCode string
	Name string
	Country string
}

var Currencies []Currency

func load() {
	Currencies = []Currency{
		{Code: "971", SCode: "AFN", Name: "Afghani", Country: "AFGHANISTAN"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "ÅLAND ISLANDS"},
		{Code: "008", SCode: "ALL", Name: "Lek", Country: "ALBANIA"},
		{Code: "012", SCode: "DZD", Name: "Algerian Dinar", Country: "ALGERIA"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "AMERICAN SAMOA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "ANDORRA"},
		{Code: "973", SCode: "AOA", Name: "Kwanza", Country: "ANGOLA"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "ANGUILLA"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "ANTIGUA AND BARBUDA"},
		{Code: "032", SCode: "ARS", Name: "Argentine Peso", Country: "ARGENTINA"},
		{Code: "051", SCode: "AMD", Name: "Armenian Dram", Country: "ARMENIA"},
		{Code: "533", SCode: "AWG", Name: "Aruban Florin", Country: "ARUBA"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "AUSTRALIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "AUSTRIA"},
		{Code: "944", SCode: "AZN", Name: "Azerbaijan Manat", Country: "AZERBAIJAN"},
		{Code: "044", SCode: "BSD", Name: "Bahamian Dollar", Country: "BAHAMAS (THE)"},
		{Code: "048", SCode: "BHD", Name: "Bahraini Dinar", Country: "BAHRAIN"},
		{Code: "050", SCode: "BDT", Name: "Taka", Country: "BANGLADESH"},
		{Code: "052", SCode: "BBD", Name: "Barbados Dollar", Country: "BARBADOS"},
		{Code: "933", SCode: "BYN", Name: "Belarusian Ruble", Country: "BELARUS"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "BELGIUM"},
		{Code: "084", SCode: "BZD", Name: "Belize Dollar", Country: "BELIZE"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "BENIN"},
		{Code: "060", SCode: "BMD", Name: "Bermudian Dollar", Country: "BERMUDA"},
		{Code: "356", SCode: "INR", Name: "Indian Rupee", Country: "BHUTAN"},
		{Code: "064", SCode: "BTN", Name: "Ngultrum", Country: "BHUTAN"},
		{Code: "068", SCode: "BOB", Name: "Boliviano", Country: "BOLIVIA (PLURINATIONAL STATE OF)"},
		{Code: "984", SCode: "BOV", Name: "Mvdol", Country: "BOLIVIA (PLURINATIONAL STATE OF)"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "BONAIRE, SINT EUSTATIUS AND SABA"},
		{Code: "977", SCode: "BAM", Name: "Convertible Mark", Country: "BOSNIA AND HERZEGOVINA"},
		{Code: "072", SCode: "BWP", Name: "Pula", Country: "BOTSWANA"},
		{Code: "578", SCode: "NOK", Name: "Norwegian Krone", Country: "BOUVET ISLAND"},
		{Code: "986", SCode: "BRL", Name: "Brazilian Real", Country: "BRAZIL"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "BRITISH INDIAN OCEAN TERRITORY (THE)"},
		{Code: "096", SCode: "BND", Name: "Brunei Dollar", Country: "BRUNEI DARUSSALAM"},
		{Code: "975", SCode: "BGN", Name: "Bulgarian Lev", Country: "BULGARIA"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "BURKINA FASO"},
		{Code: "108", SCode: "BIF", Name: "Burundi Franc", Country: "BURUNDI"},
		{Code: "132", SCode: "CVE", Name: "Cabo Verde Escudo", Country: "CABO VERDE"},
		{Code: "116", SCode: "KHR", Name: "Riel", Country: "CAMBODIA"},
		{Code: "950", SCode: "XAF", Name: "CFA Franc BEAC", Country: "CAMEROON"},
		{Code: "124", SCode: "CAD", Name: "Canadian Dollar", Country: "CANADA"},
		{Code: "136", SCode: "KYD", Name: "Cayman Islands Dollar", Country: "CAYMAN ISLANDS (THE)"},
		{Code: "950", SCode: "XAF", Name: "CFA Franc BEAC", Country: "CENTRAL AFRICAN REPUBLIC (THE)"},
		{Code: "950", SCode: "XAF", Name: "CFA Franc BEAC", Country: "CHAD"},
		{Code: "152", SCode: "CLP", Name: "Chilean Peso", Country: "CHILE"},
		{Code: "990", SCode: "CLF", Name: "Unidad de Fomento", Country: "CHILE"},
		{Code: "156", SCode: "CNY", Name: "Yuan Renminbi", Country: "CHINA"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "CHRISTMAS ISLAND"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "COCOS (KEELING) ISLANDS (THE)"},
		{Code: "170", SCode: "COP", Name: "Colombian Peso", Country: "COLOMBIA"},
		{Code: "970", SCode: "COU", Name: "Unidad de Valor Real", Country: "COLOMBIA"},
		{Code: "174", SCode: "KMF", Name: "Comorian Franc ", Country: "COMOROS (THE)"},
		{Code: "976", SCode: "CDF", Name: "Congolese Franc", Country: "CONGO (THE DEMOCRATIC REPUBLIC OF THE)"},
		{Code: "950", SCode: "XAF", Name: "CFA Franc BEAC", Country: "CONGO (THE)"},
		{Code: "554", SCode: "NZD", Name: "New Zealand Dollar", Country: "COOK ISLANDS (THE)"},
		{Code: "188", SCode: "CRC", Name: "Costa Rican Colon", Country: "COSTA RICA"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "CÔTE D'IVOIRE"},
		{Code: "191", SCode: "HRK", Name: "Kuna", Country: "CROATIA"},
		{Code: "192", SCode: "CUP", Name: "Cuban Peso", Country: "CUBA"},
		{Code: "931", SCode: "CUC", Name: "Peso Convertible", Country: "CUBA"},
		{Code: "532", SCode: "ANG", Name: "Netherlands Antillean Guilder", Country: "CURAÇAO"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "CYPRUS"},
		{Code: "203", SCode: "CZK", Name: "Czech Koruna", Country: "CZECHIA"},
		{Code: "208", SCode: "DKK", Name: "Danish Krone", Country: "DENMARK"},
		{Code: "262", SCode: "DJF", Name: "Djibouti Franc", Country: "DJIBOUTI"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "DOMINICA"},
		{Code: "214", SCode: "DOP", Name: "Dominican Peso", Country: "DOMINICAN REPUBLIC (THE)"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "ECUADOR"},
		{Code: "818", SCode: "EGP", Name: "Egyptian Pound", Country: "EGYPT"},
		{Code: "222", SCode: "SVC", Name: "El Salvador Colon", Country: "EL SALVADOR"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "EL SALVADOR"},
		{Code: "950", SCode: "XAF", Name: "CFA Franc BEAC", Country: "EQUATORIAL GUINEA"},
		{Code: "232", SCode: "ERN", Name: "Nakfa", Country: "ERITREA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "ESTONIA"},
		{Code: "748", SCode: "SZL", Name: "Lilangeni", Country: "ESWATINI"},
		{Code: "230", SCode: "ETB", Name: "Ethiopian Birr", Country: "ETHIOPIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "EUROPEAN UNION"},
		{Code: "238", SCode: "FKP", Name: "Falkland Islands Pound", Country: "FALKLAND ISLANDS (THE) [MALVINAS]"},
		{Code: "208", SCode: "DKK", Name: "Danish Krone", Country: "FAROE ISLANDS (THE)"},
		{Code: "242", SCode: "FJD", Name: "Fiji Dollar", Country: "FIJI"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "FINLAND"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "FRANCE"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "FRENCH GUIANA"},
		{Code: "953", SCode: "XPF", Name: "CFP Franc", Country: "FRENCH POLYNESIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "FRENCH SOUTHERN TERRITORIES (THE)"},
		{Code: "950", SCode: "XAF", Name: "CFA Franc BEAC", Country: "GABON"},
		{Code: "270", SCode: "GMD", Name: "Dalasi", Country: "GAMBIA (THE)"},
		{Code: "981", SCode: "GEL", Name: "Lari", Country: "GEORGIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "GERMANY"},
		{Code: "936", SCode: "GHS", Name: "Ghana Cedi", Country: "GHANA"},
		{Code: "292", SCode: "GIP", Name: "Gibraltar Pound", Country: "GIBRALTAR"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "GREECE"},
		{Code: "208", SCode: "DKK", Name: "Danish Krone", Country: "GREENLAND"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "GRENADA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "GUADELOUPE"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "GUAM"},
		{Code: "320", SCode: "GTQ", Name: "Quetzal", Country: "GUATEMALA"},
		{Code: "826", SCode: "GBP", Name: "Pound Sterling", Country: "GUERNSEY"},
		{Code: "324", SCode: "GNF", Name: "Guinean Franc", Country: "GUINEA"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "GUINEA-BISSAU"},
		{Code: "328", SCode: "GYD", Name: "Guyana Dollar", Country: "GUYANA"},
		{Code: "332", SCode: "HTG", Name: "Gourde", Country: "HAITI"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "HAITI"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "HEARD ISLAND AND McDONALD ISLANDS"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "HOLY SEE (THE)"},
		{Code: "340", SCode: "HNL", Name: "Lempira", Country: "HONDURAS"},
		{Code: "344", SCode: "HKD", Name: "Hong Kong Dollar", Country: "HONG KONG"},
		{Code: "348", SCode: "HUF", Name: "Forint", Country: "HUNGARY"},
		{Code: "352", SCode: "ISK", Name: "Iceland Krona", Country: "ICELAND"},
		{Code: "356", SCode: "INR", Name: "Indian Rupee", Country: "INDIA"},
		{Code: "360", SCode: "IDR", Name: "Rupiah", Country: "INDONESIA"},
		{Code: "960", SCode: "XDR", Name: "SDR (Special Drawing Right)", Country: "INTERNATIONAL MONETARY FUND (IMF) "},
		{Code: "364", SCode: "IRR", Name: "Iranian Rial", Country: "IRAN (ISLAMIC REPUBLIC OF)"},
		{Code: "368", SCode: "IQD", Name: "Iraqi Dinar", Country: "IRAQ"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "IRELAND"},
		{Code: "826", SCode: "GBP", Name: "Pound Sterling", Country: "ISLE OF MAN"},
		{Code: "376", SCode: "ILS", Name: "New Israeli Sheqel", Country: "ISRAEL"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "ITALY"},
		{Code: "388", SCode: "JMD", Name: "Jamaican Dollar", Country: "JAMAICA"},
		{Code: "392", SCode: "JPY", Name: "Yen", Country: "JAPAN"},
		{Code: "826", SCode: "GBP", Name: "Pound Sterling", Country: "JERSEY"},
		{Code: "400", SCode: "JOD", Name: "Jordanian Dinar", Country: "JORDAN"},
		{Code: "398", SCode: "KZT", Name: "Tenge", Country: "KAZAKHSTAN"},
		{Code: "404", SCode: "KES", Name: "Kenyan Shilling", Country: "KENYA"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "KIRIBATI"},
		{Code: "408", SCode: "KPW", Name: "North Korean Won", Country: "KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)"},
		{Code: "410", SCode: "KRW", Name: "Won", Country: "KOREA (THE REPUBLIC OF)"},
		{Code: "414", SCode: "KWD", Name: "Kuwaiti Dinar", Country: "KUWAIT"},
		{Code: "417", SCode: "KGS", Name: "Som", Country: "KYRGYZSTAN"},
		{Code: "418", SCode: "LAK", Name: "Lao Kip", Country: "LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "LATVIA"},
		{Code: "422", SCode: "LBP", Name: "Lebanese Pound", Country: "LEBANON"},
		{Code: "426", SCode: "LSL", Name: "Loti", Country: "LESOTHO"},
		{Code: "710", SCode: "ZAR", Name: "Rand", Country: "LESOTHO"},
		{Code: "430", SCode: "LRD", Name: "Liberian Dollar", Country: "LIBERIA"},
		{Code: "434", SCode: "LYD", Name: "Libyan Dinar", Country: "LIBYA"},
		{Code: "756", SCode: "CHF", Name: "Swiss Franc", Country: "LIECHTENSTEIN"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "LITHUANIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "LUXEMBOURG"},
		{Code: "446", SCode: "MOP", Name: "Pataca", Country: "MACAO"},
		{Code: "807", SCode: "MKD", Name: "Denar", Country: "NORTH MACEDONIA"},
		{Code: "969", SCode: "MGA", Name: "Malagasy Ariary", Country: "MADAGASCAR"},
		{Code: "454", SCode: "MWK", Name: "Malawi Kwacha", Country: "MALAWI"},
		{Code: "458", SCode: "MYR", Name: "Malaysian Ringgit", Country: "MALAYSIA"},
		{Code: "462", SCode: "MVR", Name: "Rufiyaa", Country: "MALDIVES"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "MALI"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "MALTA"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "MARSHALL ISLANDS (THE)"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "MARTINIQUE"},
		{Code: "929", SCode: "MRU", Name: "Ouguiya", Country: "MAURITANIA"},
		{Code: "480", SCode: "MUR", Name: "Mauritius Rupee", Country: "MAURITIUS"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "MAYOTTE"},
		{Code: "965", SCode: "XUA", Name: "ADB Unit of Account", Country: "MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP"},
		{Code: "484", SCode: "MXN", Name: "Mexican Peso", Country: "MEXICO"},
		{Code: "979", SCode: "MXV", Name: "Mexican Unidad de Inversion (UDI)", Country: "MEXICO"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "MICRONESIA (FEDERATED STATES OF)"},
		{Code: "498", SCode: "MDL", Name: "Moldovan Leu", Country: "MOLDOVA (THE REPUBLIC OF)"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "MONACO"},
		{Code: "496", SCode: "MNT", Name: "Tugrik", Country: "MONGOLIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "MONTENEGRO"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "MONTSERRAT"},
		{Code: "504", SCode: "MAD", Name: "Moroccan Dirham", Country: "MOROCCO"},
		{Code: "943", SCode: "MZN", Name: "Mozambique Metical", Country: "MOZAMBIQUE"},
		{Code: "104", SCode: "MMK", Name: "Kyat", Country: "MYANMAR"},
		{Code: "516", SCode: "NAD", Name: "Namibia Dollar", Country: "NAMIBIA"},
		{Code: "710", SCode: "ZAR", Name: "Rand", Country: "NAMIBIA"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "NAURU"},
		{Code: "524", SCode: "NPR", Name: "Nepalese Rupee", Country: "NEPAL"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "NETHERLANDS (THE)"},
		{Code: "953", SCode: "XPF", Name: "CFP Franc", Country: "NEW CALEDONIA"},
		{Code: "554", SCode: "NZD", Name: "New Zealand Dollar", Country: "NEW ZEALAND"},
		{Code: "558", SCode: "NIO", Name: "Cordoba Oro", Country: "NICARAGUA"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "NIGER (THE)"},
		{Code: "566", SCode: "NGN", Name: "Naira", Country: "NIGERIA"},
		{Code: "554", SCode: "NZD", Name: "New Zealand Dollar", Country: "NIUE"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "NORFOLK ISLAND"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "NORTHERN MARIANA ISLANDS (THE)"},
		{Code: "578", SCode: "NOK", Name: "Norwegian Krone", Country: "NORWAY"},
		{Code: "512", SCode: "OMR", Name: "Rial Omani", Country: "OMAN"},
		{Code: "586", SCode: "PKR", Name: "Pakistan Rupee", Country: "PAKISTAN"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "PALAU"},
		{Code: "590", SCode: "PAB", Name: "Balboa", Country: "PANAMA"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "PANAMA"},
		{Code: "598", SCode: "PGK", Name: "Kina", Country: "PAPUA NEW GUINEA"},
		{Code: "600", SCode: "PYG", Name: "Guarani", Country: "PARAGUAY"},
		{Code: "604", SCode: "PEN", Name: "Sol", Country: "PERU"},
		{Code: "608", SCode: "PHP", Name: "Philippine Peso", Country: "PHILIPPINES (THE)"},
		{Code: "554", SCode: "NZD", Name: "New Zealand Dollar", Country: "PITCAIRN"},
		{Code: "985", SCode: "PLN", Name: "Zloty", Country: "POLAND"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "PORTUGAL"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "PUERTO RICO"},
		{Code: "634", SCode: "QAR", Name: "Qatari Rial", Country: "QATAR"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "RÉUNION"},
		{Code: "946", SCode: "RON", Name: "Romanian Leu", Country: "ROMANIA"},
		{Code: "643", SCode: "RUB", Name: "Russian Ruble", Country: "RUSSIAN FEDERATION (THE)"},
		{Code: "646", SCode: "RWF", Name: "Rwanda Franc", Country: "RWANDA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "SAINT BARTHÉLEMY"},
		{Code: "654", SCode: "SHP", Name: "Saint Helena Pound", Country: "SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "SAINT KITTS AND NEVIS"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "SAINT LUCIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "SAINT MARTIN (FRENCH PART)"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "SAINT PIERRE AND MIQUELON"},
		{Code: "951", SCode: "XCD", Name: "East Caribbean Dollar", Country: "SAINT VINCENT AND THE GRENADINES"},
		{Code: "882", SCode: "WST", Name: "Tala", Country: "SAMOA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "SAN MARINO"},
		{Code: "930", SCode: "STN", Name: "Dobra", Country: "SAO TOME AND PRINCIPE"},
		{Code: "682", SCode: "SAR", Name: "Saudi Riyal", Country: "SAUDI ARABIA"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "SENEGAL"},
		{Code: "941", SCode: "RSD", Name: "Serbian Dinar", Country: "SERBIA"},
		{Code: "690", SCode: "SCR", Name: "Seychelles Rupee", Country: "SEYCHELLES"},
		{Code: "694", SCode: "SLL", Name: "Leone", Country: "SIERRA LEONE"},
		{Code: "702", SCode: "SGD", Name: "Singapore Dollar", Country: "SINGAPORE"},
		{Code: "532", SCode: "ANG", Name: "Netherlands Antillean Guilder", Country: "SINT MAARTEN (DUTCH PART)"},
		{Code: "994", SCode: "XSU", Name: "Sucre", Country: "SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS SUCRE"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "SLOVAKIA"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "SLOVENIA"},
		{Code: "090", SCode: "SBD", Name: "Solomon Islands Dollar", Country: "SOLOMON ISLANDS"},
		{Code: "706", SCode: "SOS", Name: "Somali Shilling", Country: "SOMALIA"},
		{Code: "710", SCode: "ZAR", Name: "Rand", Country: "SOUTH AFRICA"},
		{Code: "728", SCode: "SSP", Name: "South Sudanese Pound", Country: "SOUTH SUDAN"},
		{Code: "978", SCode: "EUR", Name: "Euro", Country: "SPAIN"},
		{Code: "144", SCode: "LKR", Name: "Sri Lanka Rupee", Country: "SRI LANKA"},
		{Code: "938", SCode: "SDG", Name: "Sudanese Pound", Country: "SUDAN (THE)"},
		{Code: "968", SCode: "SRD", Name: "Surinam Dollar", Country: "SURINAME"},
		{Code: "578", SCode: "NOK", Name: "Norwegian Krone", Country: "SVALBARD AND JAN MAYEN"},
		{Code: "752", SCode: "SEK", Name: "Swedish Krona", Country: "SWEDEN"},
		{Code: "756", SCode: "CHF", Name: "Swiss Franc", Country: "SWITZERLAND"},
		{Code: "947", SCode: "CHE", Name: "WIR Euro", Country: "SWITZERLAND"},
		{Code: "948", SCode: "CHW", Name: "WIR Franc", Country: "SWITZERLAND"},
		{Code: "760", SCode: "SYP", Name: "Syrian Pound", Country: "SYRIAN ARAB REPUBLIC"},
		{Code: "901", SCode: "TWD", Name: "New Taiwan Dollar", Country: "TAIWAN (PROVINCE OF CHINA)"},
		{Code: "972", SCode: "TJS", Name: "Somoni", Country: "TAJIKISTAN"},
		{Code: "834", SCode: "TZS", Name: "Tanzanian Shilling", Country: "TANZANIA, UNITED REPUBLIC OF"},
		{Code: "764", SCode: "THB", Name: "Baht", Country: "THAILAND"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "TIMOR-LESTE"},
		{Code: "952", SCode: "XOF", Name: "CFA Franc BCEAO", Country: "TOGO"},
		{Code: "554", SCode: "NZD", Name: "New Zealand Dollar", Country: "TOKELAU"},
		{Code: "776", SCode: "TOP", Name: "Pa’anga", Country: "TONGA"},
		{Code: "780", SCode: "TTD", Name: "Trinidad and Tobago Dollar", Country: "TRINIDAD AND TOBAGO"},
		{Code: "788", SCode: "TND", Name: "Tunisian Dinar", Country: "TUNISIA"},
		{Code: "949", SCode: "TRY", Name: "Turkish Lira", Country: "TURKEY"},
		{Code: "934", SCode: "TMT", Name: "Turkmenistan New Manat", Country: "TURKMENISTAN"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "TURKS AND CAICOS ISLANDS (THE)"},
		{Code: "036", SCode: "AUD", Name: "Australian Dollar", Country: "TUVALU"},
		{Code: "800", SCode: "UGX", Name: "Uganda Shilling", Country: "UGANDA"},
		{Code: "980", SCode: "UAH", Name: "Hryvnia", Country: "UKRAINE"},
		{Code: "784", SCode: "AED", Name: "UAE Dirham", Country: "UNITED ARAB EMIRATES (THE)"},
		{Code: "826", SCode: "GBP", Name: "Pound Sterling", Country: "UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "UNITED STATES MINOR OUTLYING ISLANDS (THE)"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "UNITED STATES OF AMERICA (THE)"},
		{Code: "997", SCode: "USN", Name: "US Dollar (Next day)", Country: "UNITED STATES OF AMERICA (THE)"},
		{Code: "858", SCode: "UYU", Name: "Peso Uruguayo", Country: "URUGUAY"},
		{Code: "940", SCode: "UYI", Name: "Uruguay Peso en Unidades Indexadas (UI)", Country: "URUGUAY"},
		{Code: "927", SCode: "UYW", Name: "Unidad Previsional", Country: "URUGUAY"},
		{Code: "860", SCode: "UZS", Name: "Uzbekistan Sum", Country: "UZBEKISTAN"},
		{Code: "548", SCode: "VUV", Name: "Vatu", Country: "VANUATU"},
		{Code: "928", SCode: "VES", Name: "Bolívar Soberano", Country: "VENEZUELA (BOLIVARIAN REPUBLIC OF)"},
		{Code: "704", SCode: "VND", Name: "Dong", Country: "VIET NAM"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "VIRGIN ISLANDS (BRITISH)"},
		{Code: "840", SCode: "USD", Name: "US Dollar", Country: "VIRGIN ISLANDS (U.S.)"},
		{Code: "953", SCode: "XPF", Name: "CFP Franc", Country: "WALLIS AND FUTUNA"},
		{Code: "504", SCode: "MAD", Name: "Moroccan Dirham", Country: "WESTERN SAHARA"},
		{Code: "886", SCode: "YER", Name: "Yemeni Rial", Country: "YEMEN"},
		{Code: "967", SCode: "ZMW", Name: "Zambian Kwacha", Country: "ZAMBIA"},
		{Code: "932", SCode: "ZWL", Name: "Zimbabwe Dollar", Country: "ZIMBABWE"},
		{Code: "955", SCode: "XBA", Name: "Bond Markets Unit European Composite Unit (EURCO)", Country: "ZZ01_Bond Markets Unit European_EURCO"},
		{Code: "956", SCode: "XBB", Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", Country: "ZZ02_Bond Markets Unit European_EMU-6"},
		{Code: "957", SCode: "XBC", Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", Country: "ZZ03_Bond Markets Unit European_EUA-9"},
		{Code: "958", SCode: "XBD", Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", Country: "ZZ04_Bond Markets Unit European_EUA-17"},
		{Code: "963", SCode: "XTS", Name: "Codes specifically reserved for testing purposes", Country: "ZZ06_Testing_Code"},
		{Code: "999", SCode: "XXX", Name: "The codes assigned for transactions where no currency is involved", Country: "ZZ07_No_Currency"},
		{Code: "959", SCode: "XAU", Name: "Gold", Country: "ZZ08_Gold"},
		{Code: "964", SCode: "XPD", Name: "Palladium", Country: "ZZ09_Palladium"},
		{Code: "962", SCode: "XPT", Name: "Platinum", Country: "ZZ10_Platinum"},
		{Code: "961", SCode: "XAG", Name: "Silver", Country: "ZZ11_Silver"},
	}
}

func init(){
	load()
	gocron.Every(2).Second().Do(load) //Обновляем токен по расписанию
	gocron.Start()
}

//GetCurrencyByCode
func GetCurrencyByCode(code string) *Currency {
	for _, c := range Currencies {
		if c.Code ==  code {
			return &c
		}
	}
	return nil
}

func GetCurrencyBySCode(sCode string) *Currency {
	for _, c := range Currencies {
		if strings.ToUpper(c.SCode) ==  strings.ToUpper(sCode) {
			return &c
		}
	}
	return nil
}

func GetCurrencyByName(name string) *Currency {
	for _, c := range Currencies {
		if strings.ToUpper(c.Name) ==  strings.ToUpper(name) {
			return &c
		}
	}
	return nil
}

func GetCurrencyByCountry(country string) *Currency {
	for _, c := range Currencies {
		if strings.ToUpper(c.Country) == strings.ToUpper(country) {
			return &c
		}
	}
	return nil
}

func GetCountrySCodeByCode(code string) string {
	for _, c := range Currencies {
		if c.Code ==  code {
			return c.SCode
		}
	}
	return ""
}

func GetCountryCodeBySCode(sCode string) string {
	for _, c := range Currencies {
		if strings.ToUpper(c.SCode) ==  strings.ToUpper(sCode) {
			return c.Code
		}
	}
	return ""
}