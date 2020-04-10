package jsons

import (
	"encoding/json"
)

const countries = `{
  "countries": [
    {
      "name": {
        "common": "Aruba",
        "official": "Aruba"
      },
      "tld": [
        ".aw"
      ],
      "cca2": "AW",
      "ccn3": "533",
      "cca3": "ABW",
      "cioc": "ARU"
    },
    {
      "name": {
        "common": "Afghanistan",
        "official": "Islamic Republic of Afghanistan"
      },
      "tld": [
        ".af"
      ],
      "cca2": "AF",
      "ccn3": "004",
      "cca3": "AFG",
      "cioc": "AFG"
    },
    {
      "name": {
        "common": "Angola",
        "official": "Republic of Angola"
      },
      "tld": [
        ".ao"
      ],
      "cca2": "AO",
      "ccn3": "024",
      "cca3": "AGO",
      "cioc": "ANG"
    },
    {
      "name": {
        "common": "Anguilla",
        "official": "Anguilla"
      },
      "tld": [
        ".ai"
      ],
      "cca2": "AI",
      "ccn3": "660",
      "cca3": "AIA",
      "cioc": ""
    },
    {
      "name": {
        "common": "Åland Islands",
        "official": "Åland Islands"
      },
      "tld": [
        ".ax"
      ],
      "cca2": "AX",
      "ccn3": "248",
      "cca3": "ALA",
      "cioc": ""
    },
    {
      "name": {
        "common": "Albania",
        "official": "Republic of Albania"
      },
      "tld": [
        ".al"
      ],
      "cca2": "AL",
      "ccn3": "008",
      "cca3": "ALB",
      "cioc": "ALB"
    },
    {
      "name": {
        "common": "Andorra",
        "official": "Principality of Andorra"
      },
      "tld": [
        ".ad"
      ],
      "cca2": "AD",
      "ccn3": "020",
      "cca3": "AND",
      "cioc": "AND"
    },
    {
      "name": {
        "common": "United Arab Emirates",
        "official": "United Arab Emirates"
      },
      "tld": [
        ".ae",
        "امارات."
      ],
      "cca2": "AE",
      "ccn3": "784",
      "cca3": "ARE",
      "cioc": "UAE"
    },
    {
      "name": {
        "common": "Argentina",
        "official": "Argentine Republic"
      },
      "tld": [
        ".ar"
      ],
      "cca2": "AR",
      "ccn3": "032",
      "cca3": "ARG",
      "cioc": "ARG"
    },
    {
      "name": {
        "common": "Armenia",
        "official": "Republic of Armenia"
      },
      "tld": [
        ".am"
      ],
      "cca2": "AM",
      "ccn3": "051",
      "cca3": "ARM",
      "cioc": "ARM"
    },
    {
      "name": {
        "common": "American Samoa",
        "official": "American Samoa"
      },
      "tld": [
        ".as"
      ],
      "cca2": "AS",
      "ccn3": "016",
      "cca3": "ASM",
      "cioc": "ASA"
    },
    {
      "name": {
        "common": "Antarctica",
        "official": "Antarctica"
      },
      "tld": [
        ".aq"
      ],
      "cca2": "AQ",
      "ccn3": "010",
      "cca3": "ATA",
      "cioc": ""
    },
    {
      "name": {
        "common": "French Southern and Antarctic Lands",
        "official": "Territory of the French Southern and Antarctic Lands"
      },
      "tld": [
        ".tf"
      ],
      "cca2": "TF",
      "ccn3": "260",
      "cca3": "ATF",
      "cioc": ""
    },
    {
      "name": {
        "common": "Antigua and Barbuda",
        "official": "Antigua and Barbuda"
      },
      "tld": [
        ".ag"
      ],
      "cca2": "AG",
      "ccn3": "028",
      "cca3": "ATG",
      "cioc": "ANT"
    },
    {
      "name": {
        "common": "Australia",
        "official": "Commonwealth of Australia"
      },
      "tld": [
        ".au"
      ],
      "cca2": "AU",
      "ccn3": "036",
      "cca3": "AUS",
      "cioc": "AUS"
    },
    {
      "name": {
        "common": "Austria",
        "official": "Republic of Austria"
      },
      "tld": [
        ".at"
      ],
      "cca2": "AT",
      "ccn3": "040",
      "cca3": "AUT",
      "cioc": "AUT"
    },
    {
      "name": {
        "common": "Azerbaijan",
        "official": "Republic of Azerbaijan"
      },
      "tld": [
        ".az"
      ],
      "cca2": "AZ",
      "ccn3": "031",
      "cca3": "AZE",
      "cioc": "AZE"
    },
    {
      "name": {
        "common": "Burundi",
        "official": "Republic of Burundi"
      },
      "tld": [
        ".bi"
      ],
      "cca2": "BI",
      "ccn3": "108",
      "cca3": "BDI",
      "cioc": "BDI"
    },
    {
      "name": {
        "common": "Belgium",
        "official": "Kingdom of Belgium"
      },
      "tld": [
        ".be"
      ],
      "cca2": "BE",
      "ccn3": "056",
      "cca3": "BEL",
      "cioc": "BEL"
    },
    {
      "name": {
        "common": "Benin",
        "official": "Republic of Benin"
      },
      "tld": [
        ".bj"
      ],
      "cca2": "BJ",
      "ccn3": "204",
      "cca3": "BEN",
      "cioc": "BEN"
    },
    {
      "name": {
        "common": "Burkina Faso",
        "official": "Burkina Faso"
      },
      "tld": [
        ".bf"
      ],
      "cca2": "BF",
      "ccn3": "854",
      "cca3": "BFA",
      "cioc": "BUR"
    },
    {
      "name": {
        "common": "Bangladesh",
        "official": "People's Republic of Bangladesh"
      },
      "tld": [
        ".bd"
      ],
      "cca2": "BD",
      "ccn3": "050",
      "cca3": "BGD",
      "cioc": "BAN"
    },
    {
      "name": {
        "common": "Bulgaria",
        "official": "Republic of Bulgaria"
      },
      "tld": [
        ".bg"
      ],
      "cca2": "BG",
      "ccn3": "100",
      "cca3": "BGR",
      "cioc": "BUL"
    },
    {
      "name": {
        "common": "Bahrain",
        "official": "Kingdom of Bahrain"
      },
      "tld": [
        ".bh"
      ],
      "cca2": "BH",
      "ccn3": "048",
      "cca3": "BHR",
      "cioc": "BRN"
    },
    {
      "name": {
        "common": "Bahamas",
        "official": "Commonwealth of the Bahamas"
      },
      "tld": [
        ".bs"
      ],
      "cca2": "BS",
      "ccn3": "044",
      "cca3": "BHS",
      "cioc": "BAH"
    },
    {
      "name": {
        "common": "Bosnia and Herzegovina",
        "official": "Bosnia and Herzegovina"
      },
      "tld": [
        ".ba"
      ],
      "cca2": "BA",
      "ccn3": "070",
      "cca3": "BIH",
      "cioc": "BIH"
    },
    {
      "name": {
        "common": "Saint Barthélemy",
        "official": "Collectivity of Saint Barthélemy"
      },
      "tld": [
        ".bl"
      ],
      "cca2": "BL",
      "ccn3": "652",
      "cca3": "BLM",
      "cioc": ""
    },
    {
      "name": {
        "common": "Saint Helena, Ascension and Tristan da Cunha",
        "official": "Saint Helena, Ascension and Tristan da Cunha"
      },
      "tld": [
        ".sh",
        ".ac"
      ],
      "cca2": "SH",
      "ccn3": "654",
      "cca3": "SHN",
      "cioc": ""
    },
    {
      "name": {
        "common": "Belarus",
        "official": "Republic of Belarus"
      },
      "tld": [
        ".by"
      ],
      "cca2": "BY",
      "ccn3": "112",
      "cca3": "BLR",
      "cioc": "BLR"
    },
    {
      "name": {
        "common": "Belize",
        "official": "Belize"
      },
      "tld": [
        ".bz"
      ],
      "cca2": "BZ",
      "ccn3": "084",
      "cca3": "BLZ",
      "cioc": "BIZ"
    },
    {
      "name": {
        "common": "Bermuda",
        "official": "Bermuda"
      },
      "tld": [
        ".bm"
      ],
      "cca2": "BM",
      "ccn3": "060",
      "cca3": "BMU",
      "cioc": "BER"
    },
    {
      "name": {
        "common": "Bolivia",
        "official": "Plurinational State of Bolivia"
      },
      "tld": [
        ".bo"
      ],
      "cca2": "BO",
      "ccn3": "068",
      "cca3": "BOL",
      "cioc": "BOL"
    },
    {
      "name": {
        "common": "Caribbean Netherlands",
        "official": "Bonaire, Sint Eustatius and Saba"
      },
      "tld": [
        ".bq",
        ".nl"
      ],
      "cca2": "BQ",
      "ccn3": "535",
      "cca3": "BES",
      "cioc": ""
    },
    {
      "name": {
        "common": "Brazil",
        "official": "Federative Republic of Brazil"
      },
      "tld": [
        ".br"
      ],
      "cca2": "BR",
      "ccn3": "076",
      "cca3": "BRA",
      "cioc": "BRA"
    },
    {
      "name": {
        "common": "Barbados",
        "official": "Barbados"
      },
      "tld": [
        ".bb"
      ],
      "cca2": "BB",
      "ccn3": "052",
      "cca3": "BRB",
      "cioc": "BAR"
    },
    {
      "name": {
        "common": "Brunei",
        "official": "Nation of Brunei, Abode of Peace"
      },
      "tld": [
        ".bn"
      ],
      "cca2": "BN",
      "ccn3": "096",
      "cca3": "BRN",
      "cioc": "BRU"
    },
    {
      "name": {
        "common": "Bhutan",
        "official": "Kingdom of Bhutan"
      },
      "tld": [
        ".bt"
      ],
      "cca2": "BT",
      "ccn3": "064",
      "cca3": "BTN",
      "cioc": "BHU"
    },
    {
      "name": {
        "common": "Bouvet Island",
        "official": "Bouvet Island"
      },
      "tld": [
        ".bv"
      ],
      "cca2": "BV",
      "ccn3": "074",
      "cca3": "BVT",
      "cioc": ""
    },
    {
      "name": {
        "common": "Botswana",
        "official": "Republic of Botswana"
      },
      "tld": [
        ".bw"
      ],
      "cca2": "BW",
      "ccn3": "072",
      "cca3": "BWA",
      "cioc": "BOT"
    },
    {
      "name": {
        "common": "Central African Republic",
        "official": "Central African Republic"
      },
      "tld": [
        ".cf"
      ],
      "cca2": "CF",
      "ccn3": "140",
      "cca3": "CAF",
      "cioc": "CAF"
    },
    {
      "name": {
        "common": "Canada",
        "official": "Canada"
      },
      "tld": [
        ".ca"
      ],
      "cca2": "CA",
      "ccn3": "124",
      "cca3": "CAN",
      "cioc": "CAN"
    },
    {
      "name": {
        "common": "Cocos (Keeling) Islands",
        "official": "Territory of the Cocos (Keeling) Islands"
      },
      "tld": [
        ".cc"
      ],
      "cca2": "CC",
      "ccn3": "166",
      "cca3": "CCK",
      "cioc": ""
    },
    {
      "name": {
        "common": "Switzerland",
        "official": "Swiss Confederation"
      },
      "tld": [
        ".ch"
      ],
      "cca2": "CH",
      "ccn3": "756",
      "cca3": "CHE",
      "cioc": "SUI"
    },
    {
      "name": {
        "common": "Chile",
        "official": "Republic of Chile"
      },
      "tld": [
        ".cl"
      ],
      "cca2": "CL",
      "ccn3": "152",
      "cca3": "CHL",
      "cioc": "CHI"
    },
    {
      "name": {
        "common": "China",
        "official": "People's Republic of China"
      },
      "tld": [
        ".cn",
        ".中国",
        ".中國",
        ".公司",
        ".网络"
      ],
      "cca2": "CN",
      "ccn3": "156",
      "cca3": "CHN",
      "cioc": "CHN"
    },
    {
      "name": {
        "common": "Ivory Coast",
        "official": "Republic of Côte d'Ivoire"
      },
      "tld": [
        ".ci"
      ],
      "cca2": "CI",
      "ccn3": "384",
      "cca3": "CIV",
      "cioc": "CIV"
    },
    {
      "name": {
        "common": "Cameroon",
        "official": "Republic of Cameroon"
      },
      "tld": [
        ".cm"
      ],
      "cca2": "CM",
      "ccn3": "120",
      "cca3": "CMR",
      "cioc": "CMR"
    },
    {
      "name": {
        "common": "DR Congo",
        "official": "Democratic Republic of the Congo"
      },
      "tld": [
        ".cd"
      ],
      "cca2": "CD",
      "ccn3": "180",
      "cca3": "COD",
      "cioc": "COD"
    },
    {
      "name": {
        "common": "Republic of the Congo",
        "official": "Republic of the Congo"
      },
      "tld": [
        ".cg"
      ],
      "cca2": "CG",
      "ccn3": "178",
      "cca3": "COG",
      "cioc": "CGO"
    },
    {
      "name": {
        "common": "Cook Islands",
        "official": "Cook Islands"
      },
      "tld": [
        ".ck"
      ],
      "cca2": "CK",
      "ccn3": "184",
      "cca3": "COK",
      "cioc": "COK"
    },
    {
      "name": {
        "common": "Colombia",
        "official": "Republic of Colombia"
      },
      "tld": [
        ".co"
      ],
      "cca2": "CO",
      "ccn3": "170",
      "cca3": "COL",
      "cioc": "COL"
    },
    {
      "name": {
        "common": "Comoros",
        "official": "Union of the Comoros"
      },
      "tld": [
        ".km"
      ],
      "cca2": "KM",
      "ccn3": "174",
      "cca3": "COM",
      "cioc": "COM"
    },
    {
      "name": {
        "common": "Cape Verde",
        "official": "Republic of Cabo Verde"
      },
      "tld": [
        ".cv"
      ],
      "cca2": "CV",
      "ccn3": "132",
      "cca3": "CPV",
      "cioc": "CPV"
    },
    {
      "name": {
        "common": "Costa Rica",
        "official": "Republic of Costa Rica"
      },
      "tld": [
        ".cr"
      ],
      "cca2": "CR",
      "ccn3": "188",
      "cca3": "CRI",
      "cioc": "CRC"
    },
    {
      "name": {
        "common": "Cuba",
        "official": "Republic of Cuba"
      },
      "tld": [
        ".cu"
      ],
      "cca2": "CU",
      "ccn3": "192",
      "cca3": "CUB",
      "cioc": "CUB"
    },
    {
      "name": {
        "common": "Curaçao",
        "official": "Country of Curaçao"
      },
      "tld": [
        ".cw"
      ],
      "cca2": "CW",
      "ccn3": "531",
      "cca3": "CUW",
      "cioc": ""
    },
    {
      "name": {
        "common": "Christmas Island",
        "official": "Territory of Christmas Island"
      },
      "tld": [
        ".cx"
      ],
      "cca2": "CX",
      "ccn3": "162",
      "cca3": "CXR",
      "cioc": ""
    },
    {
      "name": {
        "common": "Cayman Islands",
        "official": "Cayman Islands"
      },
      "tld": [
        ".ky"
      ],
      "cca2": "KY",
      "ccn3": "136",
      "cca3": "CYM",
      "cioc": "CAY"
    },
    {
      "name": {
        "common": "Cyprus",
        "official": "Republic of Cyprus"
      },
      "tld": [
        ".cy"
      ],
      "cca2": "CY",
      "ccn3": "196",
      "cca3": "CYP",
      "cioc": "CYP"
    },
    {
      "name": {
        "common": "Czechia",
        "official": "Czech Republic"
      },
      "tld": [
        ".cz"
      ],
      "cca2": "CZ",
      "ccn3": "203",
      "cca3": "CZE",
      "cioc": "CZE"
    },
    {
      "name": {
        "common": "Germany",
        "official": "Federal Republic of Germany"
      },
      "tld": [
        ".de"
      ],
      "cca2": "DE",
      "ccn3": "276",
      "cca3": "DEU",
      "cioc": "GER"
    },
    {
      "name": {
        "common": "Djibouti",
        "official": "Republic of Djibouti"
      },
      "tld": [
        ".dj"
      ],
      "cca2": "DJ",
      "ccn3": "262",
      "cca3": "DJI",
      "cioc": "DJI"
    },
    {
      "name": {
        "common": "Dominica",
        "official": "Commonwealth of Dominica"
      },
      "tld": [
        ".dm"
      ],
      "cca2": "DM",
      "ccn3": "212",
      "cca3": "DMA",
      "cioc": "DMA"
    },
    {
      "name": {
        "common": "Denmark",
        "official": "Kingdom of Denmark"
      },
      "tld": [
        ".dk"
      ],
      "cca2": "DK",
      "ccn3": "208",
      "cca3": "DNK",
      "cioc": "DEN"
    },
    {
      "name": {
        "common": "Dominican Republic",
        "official": "Dominican Republic"
      },
      "tld": [
        ".do"
      ],
      "cca2": "DO",
      "ccn3": "214",
      "cca3": "DOM",
      "cioc": "DOM"
    },
    {
      "name": {
        "common": "Algeria",
        "official": "People's Democratic Republic of Algeria"
      },
      "tld": [
        ".dz",
        "الجزائر."
      ],
      "cca2": "DZ",
      "ccn3": "012",
      "cca3": "DZA",
      "cioc": "ALG"
    },
    {
      "name": {
        "common": "Ecuador",
        "official": "Republic of Ecuador"
      },
      "tld": [
        ".ec"
      ],
      "cca2": "EC",
      "ccn3": "218",
      "cca3": "ECU",
      "cioc": "ECU"
    },
    {
      "name": {
        "common": "Egypt",
        "official": "Arab Republic of Egypt"
      },
      "tld": [
        ".eg",
        ".مصر"
      ],
      "cca2": "EG",
      "ccn3": "818",
      "cca3": "EGY",
      "cioc": "EGY"
    },
    {
      "name": {
        "common": "Eritrea",
        "official": "State of Eritrea"
      },
      "tld": [
        ".er"
      ],
      "cca2": "ER",
      "ccn3": "232",
      "cca3": "ERI",
      "cioc": "ERI"
    },
    {
      "name": {
        "common": "Western Sahara",
        "official": "Sahrawi Arab Democratic Republic"
      },
      "tld": [
        ".eh"
      ],
      "cca2": "EH",
      "ccn3": "732",
      "cca3": "ESH",
      "cioc": ""
    },
    {
      "name": {
        "common": "Spain",
        "official": "Kingdom of Spain"
      },
      "tld": [
        ".es"
      ],
      "cca2": "ES",
      "ccn3": "724",
      "cca3": "ESP",
      "cioc": "ESP"
    },
    {
      "name": {
        "common": "Estonia",
        "official": "Republic of Estonia"
      },
      "tld": [
        ".ee"
      ],
      "cca2": "EE",
      "ccn3": "233",
      "cca3": "EST",
      "cioc": "EST"
    },
    {
      "name": {
        "common": "Ethiopia",
        "official": "Federal Democratic Republic of Ethiopia"
      },
      "tld": [
        ".et"
      ],
      "cca2": "ET",
      "ccn3": "231",
      "cca3": "ETH",
      "cioc": "ETH"
    },
    {
      "name": {
        "common": "Finland",
        "official": "Republic of Finland"
      },
      "tld": [
        ".fi"
      ],
      "cca2": "FI",
      "ccn3": "246",
      "cca3": "FIN",
      "cioc": "FIN"
    },
    {
      "name": {
        "common": "Fiji",
        "official": "Republic of Fiji"
      },
      "tld": [
        ".fj"
      ],
      "cca2": "FJ",
      "ccn3": "242",
      "cca3": "FJI",
      "cioc": "FIJ"
    },
    {
      "name": {
        "common": "Falkland Islands",
        "official": "Falkland Islands"
      },
      "tld": [
        ".fk"
      ],
      "cca2": "FK",
      "ccn3": "238",
      "cca3": "FLK",
      "cioc": ""
    },
    {
      "name": {
        "common": "France",
        "official": "French Republic"
      },
      "tld": [
        ".fr"
      ],
      "cca2": "FR",
      "ccn3": "250",
      "cca3": "FRA",
      "cioc": "FRA"
    },
    {
      "name": {
        "common": "Faroe Islands",
        "official": "Faroe Islands"
      },
      "tld": [
        ".fo"
      ],
      "cca2": "FO",
      "ccn3": "234",
      "cca3": "FRO",
      "cioc": ""
    },
    {
      "name": {
        "common": "Micronesia",
        "official": "Federated States of Micronesia"
      },
      "tld": [
        ".fm"
      ],
      "cca2": "FM",
      "ccn3": "583",
      "cca3": "FSM",
      "cioc": "FSM"
    },
    {
      "name": {
        "common": "Gabon",
        "official": "Gabonese Republic"
      },
      "tld": [
        ".ga"
      ],
      "cca2": "GA",
      "ccn3": "266",
      "cca3": "GAB",
      "cioc": "GAB"
    },
    {
      "name": {
        "common": "United Kingdom",
        "official": "United Kingdom of Great Britain and Northern Ireland"
      },
      "tld": [
        ".uk"
      ],
      "cca2": "GB",
      "ccn3": "826",
      "cca3": "GBR",
      "cioc": "GBR"
    },
    {
      "name": {
        "common": "Georgia",
        "official": "Georgia"
      },
      "tld": [
        ".ge"
      ],
      "cca2": "GE",
      "ccn3": "268",
      "cca3": "GEO",
      "cioc": "GEO"
    },
    {
      "name": {
        "common": "Guernsey",
        "official": "Bailiwick of Guernsey"
      },
      "tld": [
        ".gg"
      ],
      "cca2": "GG",
      "ccn3": "831",
      "cca3": "GGY",
      "cioc": ""
    },
    {
      "name": {
        "common": "Ghana",
        "official": "Republic of Ghana"
      },
      "tld": [
        ".gh"
      ],
      "cca2": "GH",
      "ccn3": "288",
      "cca3": "GHA",
      "cioc": "GHA"
    },
    {
      "name": {
        "common": "Gibraltar",
        "official": "Gibraltar"
      },
      "tld": [
        ".gi"
      ],
      "cca2": "GI",
      "ccn3": "292",
      "cca3": "GIB",
      "cioc": ""
    },
    {
      "name": {
        "common": "Guinea",
        "official": "Republic of Guinea"
      },
      "tld": [
        ".gn"
      ],
      "cca2": "GN",
      "ccn3": "324",
      "cca3": "GIN",
      "cioc": "GUI"
    },
    {
      "name": {
        "common": "Guadeloupe",
        "official": "Guadeloupe"
      },
      "tld": [
        ".gp"
      ],
      "cca2": "GP",
      "ccn3": "312",
      "cca3": "GLP",
      "cioc": ""
    },
    {
      "name": {
        "common": "Gambia",
        "official": "Republic of the Gambia"
      },
      "tld": [
        ".gm"
      ],
      "cca2": "GM",
      "ccn3": "270",
      "cca3": "GMB",
      "cioc": "GAM"
    },
    {
      "name": {
        "common": "Guinea-Bissau",
        "official": "Republic of Guinea-Bissau"
      },
      "tld": [
        ".gw"
      ],
      "cca2": "GW",
      "ccn3": "624",
      "cca3": "GNB",
      "cioc": "GBS"
    },
    {
      "name": {
        "common": "Equatorial Guinea",
        "official": "Republic of Equatorial Guinea"
      },
      "tld": [
        ".gq"
      ],
      "cca2": "GQ",
      "ccn3": "226",
      "cca3": "GNQ",
      "cioc": "GEQ"
    },
    {
      "name": {
        "common": "Greece",
        "official": "Hellenic Republic"
      },
      "tld": [
        ".gr"
      ],
      "cca2": "GR",
      "ccn3": "300",
      "cca3": "GRC",
      "cioc": "GRE"
    },
    {
      "name": {
        "common": "Grenada",
        "official": "Grenada"
      },
      "tld": [
        ".gd"
      ],
      "cca2": "GD",
      "ccn3": "308",
      "cca3": "GRD",
      "cioc": "GRN"
    },
    {
      "name": {
        "common": "Greenland",
        "official": "Greenland"
      },
      "tld": [
        ".gl"
      ],
      "cca2": "GL",
      "ccn3": "304",
      "cca3": "GRL",
      "cioc": ""
    },
    {
      "name": {
        "common": "Guatemala",
        "official": "Republic of Guatemala"
      },
      "tld": [
        ".gt"
      ],
      "cca2": "GT",
      "ccn3": "320",
      "cca3": "GTM",
      "cioc": "GUA"
    },
    {
      "name": {
        "common": "French Guiana",
        "official": "Guiana"
      },
      "tld": [
        ".gf"
      ],
      "cca2": "GF",
      "ccn3": "254",
      "cca3": "GUF",
      "cioc": ""
    },
    {
      "name": {
        "common": "Guam",
        "official": "Guam"
      },
      "tld": [
        ".gu"
      ],
      "cca2": "GU",
      "ccn3": "316",
      "cca3": "GUM",
      "cioc": "GUM"
    },
    {
      "name": {
        "common": "Guyana",
        "official": "Co-operative Republic of Guyana"
      },
      "tld": [
        ".gy"
      ],
      "cca2": "GY",
      "ccn3": "328",
      "cca3": "GUY",
      "cioc": "GUY"
    },
    {
      "name": {
        "common": "Hong Kong",
        "official": "Hong Kong Special Administrative Region of the People's Republic of China"
      },
      "tld": [
        ".hk",
        ".香港"
      ],
      "cca2": "HK",
      "ccn3": "344",
      "cca3": "HKG",
      "cioc": "HKG"
    },
    {
      "name": {
        "common": "Heard Island and McDonald Islands",
        "official": "Heard Island and McDonald Islands"
      },
      "tld": [
        ".hm",
        ".aq"
      ],
      "cca2": "HM",
      "ccn3": "334",
      "cca3": "HMD",
      "cioc": ""
    },
    {
      "name": {
        "common": "Honduras",
        "official": "Republic of Honduras"
      },
      "tld": [
        ".hn"
      ],
      "cca2": "HN",
      "ccn3": "340",
      "cca3": "HND",
      "cioc": "HON"
    },
    {
      "name": {
        "common": "Croatia",
        "official": "Republic of Croatia"
      },
      "tld": [
        ".hr"
      ],
      "cca2": "HR",
      "ccn3": "191",
      "cca3": "HRV",
      "cioc": "CRO"
    },
    {
      "name": {
        "common": "Haiti",
        "official": "Republic of Haiti"
      },
      "tld": [
        ".ht"
      ],
      "cca2": "HT",
      "ccn3": "332",
      "cca3": "HTI",
      "cioc": "HAI"
    },
    {
      "name": {
        "common": "Hungary",
        "official": "Hungary"
      },
      "tld": [
        ".hu"
      ],
      "cca2": "HU",
      "ccn3": "348",
      "cca3": "HUN",
      "cioc": "HUN"
    },
    {
      "name": {
        "common": "Indonesia",
        "official": "Republic of Indonesia"
      },
      "tld": [
        ".id"
      ],
      "cca2": "ID",
      "ccn3": "360",
      "cca3": "IDN",
      "cioc": "INA"
    },
    {
      "name": {
        "common": "Isle of Man",
        "official": "Isle of Man"
      },
      "tld": [
        ".im"
      ],
      "cca2": "IM",
      "ccn3": "833",
      "cca3": "IMN",
      "cioc": ""
    },
    {
      "name": {
        "common": "India",
        "official": "Republic of India"
      },
      "tld": [
        ".in"
      ],
      "cca2": "IN",
      "ccn3": "356",
      "cca3": "IND",
      "cioc": "IND"
    },
    {
      "name": {
        "common": "British Indian Ocean Territory",
        "official": "British Indian Ocean Territory"
      },
      "tld": [
        ".io"
      ],
      "cca2": "IO",
      "ccn3": "086",
      "cca3": "IOT",
      "cioc": ""
    },
    {
      "name": {
        "common": "Ireland",
        "official": "Republic of Ireland"
      },
      "tld": [
        ".ie"
      ],
      "cca2": "IE",
      "ccn3": "372",
      "cca3": "IRL",
      "cioc": "IRL"
    },
    {
      "name": {
        "common": "Iran",
        "official": "Islamic Republic of Iran"
      },
      "tld": [
        ".ir",
        "ایران."
      ],
      "cca2": "IR",
      "ccn3": "364",
      "cca3": "IRN",
      "cioc": "IRI"
    },
    {
      "name": {
        "common": "Iraq",
        "official": "Republic of Iraq"
      },
      "tld": [
        ".iq"
      ],
      "cca2": "IQ",
      "ccn3": "368",
      "cca3": "IRQ",
      "cioc": "IRQ"
    },
    {
      "name": {
        "common": "Iceland",
        "official": "Iceland"
      },
      "tld": [
        ".is"
      ],
      "cca2": "IS",
      "ccn3": "352",
      "cca3": "ISL",
      "cioc": "ISL"
    },
    {
      "name": {
        "common": "Israel",
        "official": "State of Israel"
      },
      "tld": [
        ".il"
      ],
      "cca2": "IL",
      "ccn3": "376",
      "cca3": "ISR",
      "cioc": "ISR"
    },
    {
      "name": {
        "common": "Italy",
        "official": "Italian Republic"
      },
      "tld": [
        ".it"
      ],
      "cca2": "IT",
      "ccn3": "380",
      "cca3": "ITA",
      "cioc": "ITA"
    },
    {
      "name": {
        "common": "Jamaica",
        "official": "Jamaica"
      },
      "tld": [
        ".jm"
      ],
      "cca2": "JM",
      "ccn3": "388",
      "cca3": "JAM",
      "cioc": "JAM"
    },
    {
      "name": {
        "common": "Jersey",
        "official": "Bailiwick of Jersey"
      },
      "tld": [
        ".je"
      ],
      "cca2": "JE",
      "ccn3": "832",
      "cca3": "JEY",
      "cioc": ""
    },
    {
      "name": {
        "common": "Jordan",
        "official": "Hashemite Kingdom of Jordan"
      },
      "tld": [
        ".jo",
        "الاردن."
      ],
      "cca2": "JO",
      "ccn3": "400",
      "cca3": "JOR",
      "cioc": "JOR"
    },
    {
      "name": {
        "common": "Japan",
        "official": "Japan"
      },
      "tld": [
        ".jp",
        ".みんな"
      ],
      "cca2": "JP",
      "ccn3": "392",
      "cca3": "JPN",
      "cioc": "JPN"
    },
    {
      "name": {
        "common": "Kazakhstan",
        "official": "Republic of Kazakhstan"
      },
      "tld": [
        ".kz",
        ".қаз"
      ],
      "cca2": "KZ",
      "ccn3": "398",
      "cca3": "KAZ",
      "cioc": "KAZ"
    },
    {
      "name": {
        "common": "Kenya",
        "official": "Republic of Kenya"
      },
      "tld": [
        ".ke"
      ],
      "cca2": "KE",
      "ccn3": "404",
      "cca3": "KEN",
      "cioc": "KEN"
    },
    {
      "name": {
        "common": "Kyrgyzstan",
        "official": "Kyrgyz Republic"
      },
      "tld": [
        ".kg"
      ],
      "cca2": "KG",
      "ccn3": "417",
      "cca3": "KGZ",
      "cioc": "KGZ"
    },
    {
      "name": {
        "common": "Cambodia",
        "official": "Kingdom of Cambodia"
      },
      "tld": [
        ".kh"
      ],
      "cca2": "KH",
      "ccn3": "116",
      "cca3": "KHM",
      "cioc": "CAM"
    },
    {
      "name": {
        "common": "Kiribati",
        "official": "Independent and Sovereign Republic of Kiribati"
      },
      "tld": [
        ".ki"
      ],
      "cca2": "KI",
      "ccn3": "296",
      "cca3": "KIR",
      "cioc": "KIR"
    },
    {
      "name": {
        "common": "Saint Kitts and Nevis",
        "official": "Federation of Saint Christopher and Nevis"
      },
      "tld": [
        ".kn"
      ],
      "cca2": "KN",
      "ccn3": "659",
      "cca3": "KNA",
      "cioc": "SKN"
    },
    {
      "name": {
        "common": "South Korea",
        "official": "Republic of Korea"
      },
      "tld": [
        ".kr",
        ".한국"
      ],
      "cca2": "KR",
      "ccn3": "410",
      "cca3": "KOR",
      "cioc": "KOR"
    },
    {
      "name": {
        "common": "Kosovo",
        "official": "Republic of Kosovo"
      },
      "tld": [],
      "cca2": "XK",
      "ccn3": "",
      "cca3": "UNK",
      "cioc": "KOS"
    },
    {
      "name": {
        "common": "Kuwait",
        "official": "State of Kuwait"
      },
      "tld": [
        ".kw"
      ],
      "cca2": "KW",
      "ccn3": "414",
      "cca3": "KWT",
      "cioc": "KUW"
    },
    {
      "name": {
        "common": "Laos",
        "official": "Lao People's Democratic Republic"
      },
      "tld": [
        ".la"
      ],
      "cca2": "LA",
      "ccn3": "418",
      "cca3": "LAO",
      "cioc": "LAO"
    },
    {
      "name": {
        "common": "Lebanon",
        "official": "Lebanese Republic"
      },
      "tld": [
        ".lb"
      ],
      "cca2": "LB",
      "ccn3": "422",
      "cca3": "LBN",
      "cioc": "LIB"
    },
    {
      "name": {
        "common": "Liberia",
        "official": "Republic of Liberia"
      },
      "tld": [
        ".lr"
      ],
      "cca2": "LR",
      "ccn3": "430",
      "cca3": "LBR",
      "cioc": "LBR"
    },
    {
      "name": {
        "common": "Libya",
        "official": "State of Libya"
      },
      "tld": [
        ".ly"
      ],
      "cca2": "LY",
      "ccn3": "434",
      "cca3": "LBY",
      "cioc": "LBA"
    },
    {
      "name": {
        "common": "Saint Lucia",
        "official": "Saint Lucia"
      },
      "tld": [
        ".lc"
      ],
      "cca2": "LC",
      "ccn3": "662",
      "cca3": "LCA",
      "cioc": "LCA"
    },
    {
      "name": {
        "common": "Liechtenstein",
        "official": "Principality of Liechtenstein"
      },
      "tld": [
        ".li"
      ],
      "cca2": "LI",
      "ccn3": "438",
      "cca3": "LIE",
      "cioc": "LIE"
    },
    {
      "name": {
        "common": "Sri Lanka",
        "official": "Democratic Socialist Republic of Sri Lanka"
      },
      "tld": [
        ".lk",
        ".இலங்கை",
        ".ලංකා"
      ],
      "cca2": "LK",
      "ccn3": "144",
      "cca3": "LKA",
      "cioc": "SRI"
    },
    {
      "name": {
        "common": "Lesotho",
        "official": "Kingdom of Lesotho"
      },
      "tld": [
        ".ls"
      ],
      "cca2": "LS",
      "ccn3": "426",
      "cca3": "LSO",
      "cioc": "LES"
    },
    {
      "name": {
        "common": "Lithuania",
        "official": "Republic of Lithuania"
      },
      "tld": [
        ".lt"
      ],
      "cca2": "LT",
      "ccn3": "440",
      "cca3": "LTU",
      "cioc": "LTU"
    },
    {
      "name": {
        "common": "Luxembourg",
        "official": "Grand Duchy of Luxembourg"
      },
      "tld": [
        ".lu"
      ],
      "cca2": "LU",
      "ccn3": "442",
      "cca3": "LUX",
      "cioc": "LUX"
    },
    {
      "name": {
        "common": "Latvia",
        "official": "Republic of Latvia"
      },
      "tld": [
        ".lv"
      ],
      "cca2": "LV",
      "ccn3": "428",
      "cca3": "LVA",
      "cioc": "LAT"
    },
    {
      "name": {
        "common": "Macau",
        "official": "Macao Special Administrative Region of the People's Republic of China"
      },
      "tld": [
        ".mo"
      ],
      "cca2": "MO",
      "ccn3": "446",
      "cca3": "MAC",
      "cioc": ""
    },
    {
      "name": {
        "common": "Saint Martin",
        "official": "Saint Martin"
      },
      "tld": [
        ".fr",
        ".gp"
      ],
      "cca2": "MF",
      "ccn3": "663",
      "cca3": "MAF",
      "cioc": ""
    },
    {
      "name": {
        "common": "Morocco",
        "official": "Kingdom of Morocco"
      },
      "tld": [
        ".ma",
        "المغرب."
      ],
      "cca2": "MA",
      "ccn3": "504",
      "cca3": "MAR",
      "cioc": "MAR"
    },
    {
      "name": {
        "common": "Monaco",
        "official": "Principality of Monaco"
      },
      "tld": [
        ".mc"
      ],
      "cca2": "MC",
      "ccn3": "492",
      "cca3": "MCO",
      "cioc": "MON"
    },
    {
      "name": {
        "common": "Moldova",
        "official": "Republic of Moldova"
      },
      "tld": [
        ".md"
      ],
      "cca2": "MD",
      "ccn3": "498",
      "cca3": "MDA",
      "cioc": "MDA"
    },
    {
      "name": {
        "common": "Madagascar",
        "official": "Republic of Madagascar"
      },
      "tld": [
        ".mg"
      ],
      "cca2": "MG",
      "ccn3": "450",
      "cca3": "MDG",
      "cioc": "MAD"
    },
    {
      "name": {
        "common": "Maldives",
        "official": "Republic of the Maldives"
      },
      "tld": [
        ".mv"
      ],
      "cca2": "MV",
      "ccn3": "462",
      "cca3": "MDV",
      "cioc": "MDV"
    },
    {
      "name": {
        "common": "Mexico",
        "official": "United Mexican States"
      },
      "tld": [
        ".mx"
      ],
      "cca2": "MX",
      "ccn3": "484",
      "cca3": "MEX",
      "cioc": "MEX"
    },
    {
      "name": {
        "common": "Marshall Islands",
        "official": "Republic of the Marshall Islands"
      },
      "tld": [
        ".mh"
      ],
      "cca2": "MH",
      "ccn3": "584",
      "cca3": "MHL",
      "cioc": "MHL"
    },
    {
      "name": {
        "common": "North Macedonia",
        "official": "Republic of North Macedonia"
      },
      "tld": [
        ".mk"
      ],
      "cca2": "MK",
      "ccn3": "807",
      "cca3": "MKD",
      "cioc": "MKD"
    },
    {
      "name": {
        "common": "Mali",
        "official": "Republic of Mali"
      },
      "tld": [
        ".ml"
      ],
      "cca2": "ML",
      "ccn3": "466",
      "cca3": "MLI",
      "cioc": "MLI"
    },
    {
      "name": {
        "common": "Malta",
        "official": "Republic of Malta"
      },
      "tld": [
        ".mt"
      ],
      "cca2": "MT",
      "ccn3": "470",
      "cca3": "MLT",
      "cioc": "MLT"
    },
    {
      "name": {
        "common": "Myanmar",
        "official": "Republic of the Union of Myanmar"
      },
      "tld": [
        ".mm"
      ],
      "cca2": "MM",
      "ccn3": "104",
      "cca3": "MMR",
      "cioc": "MYA"
    },
    {
      "name": {
        "common": "Montenegro",
        "official": "Montenegro"
      },
      "tld": [
        ".me"
      ],
      "cca2": "ME",
      "ccn3": "499",
      "cca3": "MNE",
      "cioc": "MNE"
    },
    {
      "name": {
        "common": "Mongolia",
        "official": "Mongolia"
      },
      "tld": [
        ".mn"
      ],
      "cca2": "MN",
      "ccn3": "496",
      "cca3": "MNG",
      "cioc": "MGL"
    },
    {
      "name": {
        "common": "Northern Mariana Islands",
        "official": "Commonwealth of the Northern Mariana Islands"
      },
      "tld": [
        ".mp"
      ],
      "cca2": "MP",
      "ccn3": "580",
      "cca3": "MNP",
      "cioc": ""
    },
    {
      "name": {
        "common": "Mozambique",
        "official": "Republic of Mozambique"
      },
      "tld": [
        ".mz"
      ],
      "cca2": "MZ",
      "ccn3": "508",
      "cca3": "MOZ",
      "cioc": "MOZ"
    },
    {
      "name": {
        "common": "Mauritania",
        "official": "Islamic Republic of Mauritania"
      },
      "tld": [
        ".mr"
      ],
      "cca2": "MR",
      "ccn3": "478",
      "cca3": "MRT",
      "cioc": "MTN"
    },
    {
      "name": {
        "common": "Montserrat",
        "official": "Montserrat"
      },
      "tld": [
        ".ms"
      ],
      "cca2": "MS",
      "ccn3": "500",
      "cca3": "MSR",
      "cioc": ""
    },
    {
      "name": {
        "common": "Martinique",
        "official": "Martinique"
      },
      "tld": [
        ".mq"
      ],
      "cca2": "MQ",
      "ccn3": "474",
      "cca3": "MTQ",
      "cioc": ""
    },
    {
      "name": {
        "common": "Mauritius",
        "official": "Republic of Mauritius"
      },
      "tld": [
        ".mu"
      ],
      "cca2": "MU",
      "ccn3": "480",
      "cca3": "MUS",
      "cioc": "MRI"
    },
    {
      "name": {
        "common": "Malawi",
        "official": "Republic of Malawi"
      },
      "tld": [
        ".mw"
      ],
      "cca2": "MW",
      "ccn3": "454",
      "cca3": "MWI",
      "cioc": "MAW"
    },
    {
      "name": {
        "common": "Malaysia",
        "official": "Malaysia"
      },
      "tld": [
        ".my"
      ],
      "cca2": "MY",
      "ccn3": "458",
      "cca3": "MYS",
      "cioc": "MAS"
    },
    {
      "name": {
        "common": "Mayotte",
        "official": "Department of Mayotte"
      },
      "tld": [
        ".yt"
      ],
      "cca2": "YT",
      "ccn3": "175",
      "cca3": "MYT",
      "cioc": ""
    },
    {
      "name": {
        "common": "Namibia",
        "official": "Republic of Namibia"
      },
      "tld": [
        ".na"
      ],
      "cca2": "NA",
      "ccn3": "516",
      "cca3": "NAM",
      "cioc": "NAM"
    },
    {
      "name": {
        "common": "New Caledonia",
        "official": "New Caledonia"
      },
      "tld": [
        ".nc"
      ],
      "cca2": "NC",
      "ccn3": "540",
      "cca3": "NCL",
      "cioc": ""
    },
    {
      "name": {
        "common": "Niger",
        "official": "Republic of Niger"
      },
      "tld": [
        ".ne"
      ],
      "cca2": "NE",
      "ccn3": "562",
      "cca3": "NER",
      "cioc": "NIG"
    },
    {
      "name": {
        "common": "Norfolk Island",
        "official": "Territory of Norfolk Island"
      },
      "tld": [
        ".nf"
      ],
      "cca2": "NF",
      "ccn3": "574",
      "cca3": "NFK",
      "cioc": ""
    },
    {
      "name": {
        "common": "Nigeria",
        "official": "Federal Republic of Nigeria"
      },
      "tld": [
        ".ng"
      ],
      "cca2": "NG",
      "ccn3": "566",
      "cca3": "NGA",
      "cioc": "NGR"
    },
    {
      "name": {
        "common": "Nicaragua",
        "official": "Republic of Nicaragua"
      },
      "tld": [
        ".ni"
      ],
      "cca2": "NI",
      "ccn3": "558",
      "cca3": "NIC",
      "cioc": "NCA"
    },
    {
      "name": {
        "common": "Niue",
        "official": "Niue"
      },
      "tld": [
        ".nu"
      ],
      "cca2": "NU",
      "ccn3": "570",
      "cca3": "NIU",
      "cioc": ""
    },
    {
      "name": {
        "common": "Netherlands",
        "official": "Kingdom of the Netherlands"
      },
      "tld": [
        ".nl"
      ],
      "cca2": "NL",
      "ccn3": "528",
      "cca3": "NLD",
      "cioc": "NED"
    },
    {
      "name": {
        "common": "Norway",
        "official": "Kingdom of Norway"
      },
      "tld": [
        ".no"
      ],
      "cca2": "NO",
      "ccn3": "578",
      "cca3": "NOR",
      "cioc": "NOR"
    },
    {
      "name": {
        "common": "Nepal",
        "official": "Federal Democratic Republic of Nepal"
      },
      "tld": [
        ".np"
      ],
      "cca2": "NP",
      "ccn3": "524",
      "cca3": "NPL",
      "cioc": "NEP"
    },
    {
      "name": {
        "common": "Nauru",
        "official": "Republic of Nauru"
      },
      "tld": [
        ".nr"
      ],
      "cca2": "NR",
      "ccn3": "520",
      "cca3": "NRU",
      "cioc": "NRU"
    },
    {
      "name": {
        "common": "New Zealand",
        "official": "New Zealand"
      },
      "tld": [
        ".nz"
      ],
      "cca2": "NZ",
      "ccn3": "554",
      "cca3": "NZL",
      "cioc": "NZL"
    },
    {
      "name": {
        "common": "Oman",
        "official": "Sultanate of Oman"
      },
      "tld": [
        ".om"
      ],
      "cca2": "OM",
      "ccn3": "512",
      "cca3": "OMN",
      "cioc": "OMA"
    },
    {
      "name": {
        "common": "Pakistan",
        "official": "Islamic Republic of Pakistan"
      },
      "tld": [
        ".pk"
      ],
      "cca2": "PK",
      "ccn3": "586",
      "cca3": "PAK",
      "cioc": "PAK"
    },
    {
      "name": {
        "common": "Panama",
        "official": "Republic of Panama"
      },
      "tld": [
        ".pa"
      ],
      "cca2": "PA",
      "ccn3": "591",
      "cca3": "PAN",
      "cioc": "PAN"
    },
    {
      "name": {
        "common": "Pitcairn Islands",
        "official": "Pitcairn Group of Islands"
      },
      "tld": [
        ".pn"
      ],
      "cca2": "PN",
      "ccn3": "612",
      "cca3": "PCN",
      "cioc": ""
    },
    {
      "name": {
        "common": "Peru",
        "official": "Republic of Peru"
      },
      "tld": [
        ".pe"
      ],
      "cca2": "PE",
      "ccn3": "604",
      "cca3": "PER",
      "cioc": "PER"
    },
    {
      "name": {
        "common": "Philippines",
        "official": "Republic of the Philippines"
      },
      "tld": [
        ".ph"
      ],
      "cca2": "PH",
      "ccn3": "608",
      "cca3": "PHL",
      "cioc": "PHI"
    },
    {
      "name": {
        "common": "Palau",
        "official": "Republic of Palau"
      },
      "tld": [
        ".pw"
      ],
      "cca2": "PW",
      "ccn3": "585",
      "cca3": "PLW",
      "cioc": "PLW"
    },
    {
      "name": {
        "common": "Papua New Guinea",
        "official": "Independent State of Papua New Guinea"
      },
      "tld": [
        ".pg"
      ],
      "cca2": "PG",
      "ccn3": "598",
      "cca3": "PNG",
      "cioc": "PNG"
    },
    {
      "name": {
        "common": "Poland",
        "official": "Republic of Poland"
      },
      "tld": [
        ".pl"
      ],
      "cca2": "PL",
      "ccn3": "616",
      "cca3": "POL",
      "cioc": "POL"
    },
    {
      "name": {
        "common": "Puerto Rico",
        "official": "Commonwealth of Puerto Rico"
      },
      "tld": [
        ".pr"
      ],
      "cca2": "PR",
      "ccn3": "630",
      "cca3": "PRI",
      "cioc": "PUR"
    },
    {
      "name": {
        "common": "North Korea",
        "official": "Democratic People's Republic of Korea"
      },
      "tld": [
        ".kp"
      ],
      "cca2": "KP",
      "ccn3": "408",
      "cca3": "PRK",
      "cioc": "PRK"
    },
    {
      "name": {
        "common": "Portugal",
        "official": "Portuguese Republic"
      },
      "tld": [
        ".pt"
      ],
      "cca2": "PT",
      "ccn3": "620",
      "cca3": "PRT",
      "cioc": "POR"
    },
    {
      "name": {
        "common": "Paraguay",
        "official": "Republic of Paraguay"
      },
      "tld": [
        ".py"
      ],
      "cca2": "PY",
      "ccn3": "600",
      "cca3": "PRY",
      "cioc": "PAR"
    },
    {
      "name": {
        "common": "Palestine",
        "official": "State of Palestine"
      },
      "tld": [
        ".ps",
        "فلسطين."
      ],
      "cca2": "PS",
      "ccn3": "275",
      "cca3": "PSE",
      "cioc": "PLE"
    },
    {
      "name": {
        "common": "French Polynesia",
        "official": "French Polynesia"
      },
      "tld": [
        ".pf"
      ],
      "cca2": "PF",
      "ccn3": "258",
      "cca3": "PYF",
      "cioc": ""
    },
    {
      "name": {
        "common": "Qatar",
        "official": "State of Qatar"
      },
      "tld": [
        ".qa",
        "قطر."
      ],
      "cca2": "QA",
      "ccn3": "634",
      "cca3": "QAT",
      "cioc": "QAT"
    },
    {
      "name": {
        "common": "Réunion",
        "official": "Réunion Island"
      },
      "tld": [
        ".re"
      ],
      "cca2": "RE",
      "ccn3": "638",
      "cca3": "REU",
      "cioc": ""
    },
    {
      "name": {
        "common": "Romania",
        "official": "Romania"
      },
      "tld": [
        ".ro"
      ],
      "cca2": "RO",
      "ccn3": "642",
      "cca3": "ROU",
      "cioc": "ROU"
    },
    {
      "name": {
        "common": "Russia",
        "official": "Russian Federation"
      },
      "tld": [
        ".ru",
        ".su",
        ".рф"
      ],
      "cca2": "RU",
      "ccn3": "643",
      "cca3": "RUS",
      "cioc": "RUS"
    },
    {
      "name": {
        "common": "Rwanda",
        "official": "Republic of Rwanda"
      },
      "tld": [
        ".rw"
      ],
      "cca2": "RW",
      "ccn3": "646",
      "cca3": "RWA",
      "cioc": "RWA"
    },
    {
      "name": {
        "common": "Saudi Arabia",
        "official": "Kingdom of Saudi Arabia"
      },
      "tld": [
        ".sa",
        ".السعودية"
      ],
      "cca2": "SA",
      "ccn3": "682",
      "cca3": "SAU",
      "cioc": "KSA"
    },
    {
      "name": {
        "common": "Sudan",
        "official": "Republic of the Sudan"
      },
      "tld": [
        ".sd"
      ],
      "cca2": "SD",
      "ccn3": "729",
      "cca3": "SDN",
      "cioc": "SUD"
    },
    {
      "name": {
        "common": "Senegal",
        "official": "Republic of Senegal"
      },
      "tld": [
        ".sn"
      ],
      "cca2": "SN",
      "ccn3": "686",
      "cca3": "SEN",
      "cioc": "SEN"
    },
    {
      "name": {
        "common": "Singapore",
        "official": "Republic of Singapore"
      },
      "tld": [
        ".sg",
        ".新加坡",
        ".சிங்கப்பூர்"
      ],
      "cca2": "SG",
      "ccn3": "702",
      "cca3": "SGP",
      "cioc": "SIN"
    },
    {
      "name": {
        "common": "South Georgia",
        "official": "South Georgia and the South Sandwich Islands"
      },
      "tld": [
        ".gs"
      ],
      "cca2": "GS",
      "ccn3": "239",
      "cca3": "SGS",
      "cioc": ""
    },
    {
      "name": {
        "common": "Svalbard and Jan Mayen",
        "official": "Svalbard og Jan Mayen"
      },
      "tld": [
        ".sj"
      ],
      "cca2": "SJ",
      "ccn3": "744",
      "cca3": "SJM",
      "cioc": ""
    },
    {
      "name": {
        "common": "Solomon Islands",
        "official": "Solomon Islands"
      },
      "tld": [
        ".sb"
      ],
      "cca2": "SB",
      "ccn3": "090",
      "cca3": "SLB",
      "cioc": "SOL"
    },
    {
      "name": {
        "common": "Sierra Leone",
        "official": "Republic of Sierra Leone"
      },
      "tld": [
        ".sl"
      ],
      "cca2": "SL",
      "ccn3": "694",
      "cca3": "SLE",
      "cioc": "SLE"
    },
    {
      "name": {
        "common": "El Salvador",
        "official": "Republic of El Salvador"
      },
      "tld": [
        ".sv"
      ],
      "cca2": "SV",
      "ccn3": "222",
      "cca3": "SLV",
      "cioc": "ESA"
    },
    {
      "name": {
        "common": "San Marino",
        "official": "Most Serene Republic of San Marino"
      },
      "tld": [
        ".sm"
      ],
      "cca2": "SM",
      "ccn3": "674",
      "cca3": "SMR",
      "cioc": "SMR"
    },
    {
      "name": {
        "common": "Somalia",
        "official": "Federal Republic of Somalia"
      },
      "tld": [
        ".so"
      ],
      "cca2": "SO",
      "ccn3": "706",
      "cca3": "SOM",
      "cioc": "SOM"
    },
    {
      "name": {
        "common": "Saint Pierre and Miquelon",
        "official": "Saint Pierre and Miquelon"
      },
      "tld": [
        ".pm"
      ],
      "cca2": "PM",
      "ccn3": "666",
      "cca3": "SPM",
      "cioc": ""
    },
    {
      "name": {
        "common": "Serbia",
        "official": "Republic of Serbia"
      },
      "tld": [
        ".rs",
        ".срб"
      ],
      "cca2": "RS",
      "ccn3": "688",
      "cca3": "SRB",
      "cioc": "SRB"
    },
    {
      "name": {
        "common": "South Sudan",
        "official": "Republic of South Sudan"
      },
      "tld": [
        ".ss"
      ],
      "cca2": "SS",
      "ccn3": "728",
      "cca3": "SSD",
      "cioc": ""
    },
    {
      "name": {
        "common": "São Tomé and Príncipe",
        "official": "Democratic Republic of São Tomé and Príncipe"
      },
      "tld": [
        ".st"
      ],
      "cca2": "ST",
      "ccn3": "678",
      "cca3": "STP",
      "cioc": "STP"
    },
    {
      "name": {
        "common": "Suriname",
        "official": "Republic of Suriname"
      },
      "tld": [
        ".sr"
      ],
      "cca2": "SR",
      "ccn3": "740",
      "cca3": "SUR",
      "cioc": "SUR"
    },
    {
      "name": {
        "common": "Slovakia",
        "official": "Slovak Republic"
      },
      "tld": [
        ".sk"
      ],
      "cca2": "SK",
      "ccn3": "703",
      "cca3": "SVK",
      "cioc": "SVK"
    },
    {
      "name": {
        "common": "Slovenia",
        "official": "Republic of Slovenia"
      },
      "tld": [
        ".si"
      ],
      "cca2": "SI",
      "ccn3": "705",
      "cca3": "SVN",
      "cioc": "SLO"
    },
    {
      "name": {
        "common": "Sweden",
        "official": "Kingdom of Sweden"
      },
      "tld": [
        ".se"
      ],
      "cca2": "SE",
      "ccn3": "752",
      "cca3": "SWE",
      "cioc": "SWE"
    },
    {
      "name": {
        "common": "Eswatini",
        "official": "Kingdom of Eswatini"
      },
      "tld": [
        ".sz"
      ],
      "cca2": "SZ",
      "ccn3": "748",
      "cca3": "SWZ",
      "cioc": "SWZ"
    },
    {
      "name": {
        "common": "Sint Maarten",
        "official": "Sint Maarten"
      },
      "tld": [
        ".sx"
      ],
      "cca2": "SX",
      "ccn3": "534",
      "cca3": "SXM",
      "cioc": ""
    },
    {
      "name": {
        "common": "Seychelles",
        "official": "Republic of Seychelles"
      },
      "tld": [
        ".sc"
      ],
      "cca2": "SC",
      "ccn3": "690",
      "cca3": "SYC",
      "cioc": "SEY"
    },
    {
      "name": {
        "common": "Syria",
        "official": "Syrian Arab Republic"
      },
      "tld": [
        ".sy",
        "سوريا."
      ],
      "cca2": "SY",
      "ccn3": "760",
      "cca3": "SYR",
      "cioc": "SYR"
    },
    {
      "name": {
        "common": "Turks and Caicos Islands",
        "official": "Turks and Caicos Islands"
      },
      "tld": [
        ".tc"
      ],
      "cca2": "TC",
      "ccn3": "796",
      "cca3": "TCA",
      "cioc": ""
    },
    {
      "name": {
        "common": "Chad",
        "official": "Republic of Chad"
      },
      "tld": [
        ".td"
      ],
      "cca2": "TD",
      "ccn3": "148",
      "cca3": "TCD",
      "cioc": "CHA"
    },
    {
      "name": {
        "common": "Togo",
        "official": "Togolese Republic"
      },
      "tld": [
        ".tg"
      ],
      "cca2": "TG",
      "ccn3": "768",
      "cca3": "TGO",
      "cioc": "TOG"
    },
    {
      "name": {
        "common": "Thailand",
        "official": "Kingdom of Thailand"
      },
      "tld": [
        ".th",
        ".ไทย"
      ],
      "cca2": "TH",
      "ccn3": "764",
      "cca3": "THA",
      "cioc": "THA"
    },
    {
      "name": {
        "common": "Tajikistan",
        "official": "Republic of Tajikistan"
      },
      "tld": [
        ".tj"
      ],
      "cca2": "TJ",
      "ccn3": "762",
      "cca3": "TJK",
      "cioc": "TJK"
    },
    {
      "name": {
        "common": "Tokelau",
        "official": "Tokelau"
      },
      "tld": [
        ".tk"
      ],
      "cca2": "TK",
      "ccn3": "772",
      "cca3": "TKL",
      "cioc": ""
    },
    {
      "name": {
        "common": "Turkmenistan",
        "official": "Turkmenistan"
      },
      "tld": [
        ".tm"
      ],
      "cca2": "TM",
      "ccn3": "795",
      "cca3": "TKM",
      "cioc": "TKM"
    },
    {
      "name": {
        "common": "Timor-Leste",
        "official": "Democratic Republic of Timor-Leste"
      },
      "tld": [
        ".tl"
      ],
      "cca2": "TL",
      "ccn3": "626",
      "cca3": "TLS",
      "cioc": "TLS"
    },
    {
      "name": {
        "common": "Tonga",
        "official": "Kingdom of Tonga"
      },
      "tld": [
        ".to"
      ],
      "cca2": "TO",
      "ccn3": "776",
      "cca3": "TON",
      "cioc": "TGA"
    },
    {
      "name": {
        "common": "Trinidad and Tobago",
        "official": "Republic of Trinidad and Tobago"
      },
      "tld": [
        ".tt"
      ],
      "cca2": "TT",
      "ccn3": "780",
      "cca3": "TTO",
      "cioc": "TTO"
    },
    {
      "name": {
        "common": "Tunisia",
        "official": "Tunisian Republic"
      },
      "tld": [
        ".tn"
      ],
      "cca2": "TN",
      "ccn3": "788",
      "cca3": "TUN",
      "cioc": "TUN"
    },
    {
      "name": {
        "common": "Turkey",
        "official": "Republic of Turkey"
      },
      "tld": [
        ".tr"
      ],
      "cca2": "TR",
      "ccn3": "792",
      "cca3": "TUR",
      "cioc": "TUR"
    },
    {
      "name": {
        "common": "Tuvalu",
        "official": "Tuvalu"
      },
      "tld": [
        ".tv"
      ],
      "cca2": "TV",
      "ccn3": "798",
      "cca3": "TUV",
      "cioc": "TUV"
    },
    {
      "name": {
        "common": "Taiwan",
        "official": "Republic of China (Taiwan)"
      },
      "tld": [
        ".tw",
        ".台灣",
        ".台湾"
      ],
      "cca2": "TW",
      "ccn3": "158",
      "cca3": "TWN",
      "cioc": "TPE"
    },
    {
      "name": {
        "common": "Tanzania",
        "official": "United Republic of Tanzania"
      },
      "tld": [
        ".tz"
      ],
      "cca2": "TZ",
      "ccn3": "834",
      "cca3": "TZA",
      "cioc": "TAN"
    },
    {
      "name": {
        "common": "Uganda",
        "official": "Republic of Uganda"
      },
      "tld": [
        ".ug"
      ],
      "cca2": "UG",
      "ccn3": "800",
      "cca3": "UGA",
      "cioc": "UGA"
    },
    {
      "name": {
        "common": "Ukraine",
        "official": "Ukraine"
      },
      "tld": [
        ".ua",
        ".укр"
      ],
      "cca2": "UA",
      "ccn3": "804",
      "cca3": "UKR",
      "cioc": "UKR"
    },
    {
      "name": {
        "common": "United States Minor Outlying Islands",
        "official": "United States Minor Outlying Islands"
      },
      "tld": [
        ".us"
      ],
      "cca2": "UM",
      "ccn3": "581",
      "cca3": "UMI",
      "cioc": ""
    },
    {
      "name": {
        "common": "Uruguay",
        "official": "Oriental Republic of Uruguay"
      },
      "tld": [
        ".uy"
      ],
      "cca2": "UY",
      "ccn3": "858",
      "cca3": "URY",
      "cioc": "URU"
    },
    {
      "name": {
        "common": "United States",
        "official": "United States of America"
      },
      "tld": [
        ".us"
      ],
      "cca2": "US",
      "ccn3": "840",
      "cca3": "USA",
      "cioc": "USA"
    },
    {
      "name": {
        "common": "Uzbekistan",
        "official": "Republic of Uzbekistan"
      },
      "tld": [
        ".uz"
      ],
      "cca2": "UZ",
      "ccn3": "860",
      "cca3": "UZB",
      "cioc": "UZB"
    },
    {
      "name": {
        "common": "Vatican City",
        "official": "Vatican City State"
      },
      "tld": [
        ".va"
      ],
      "cca2": "VA",
      "ccn3": "336",
      "cca3": "VAT",
      "cioc": ""
    },
    {
      "name": {
        "common": "Saint Vincent and the Grenadines",
        "official": "Saint Vincent and the Grenadines"
      },
      "tld": [
        ".vc"
      ],
      "cca2": "VC",
      "ccn3": "670",
      "cca3": "VCT",
      "cioc": "VIN"
    },
    {
      "name": {
        "common": "Venezuela",
        "official": "Bolivarian Republic of Venezuela"
      },
      "tld": [
        ".ve"
      ],
      "cca2": "VE",
      "ccn3": "862",
      "cca3": "VEN",
      "cioc": "VEN"
    },
    {
      "name": {
        "common": "British Virgin Islands",
        "official": "Virgin Islands"
      },
      "tld": [
        ".vg"
      ],
      "cca2": "VG",
      "ccn3": "092",
      "cca3": "VGB",
      "cioc": "IVB"
    },
    {
      "name": {
        "common": "United States Virgin Islands",
        "official": "Virgin Islands of the United States"
      },
      "tld": [
        ".vi"
      ],
      "cca2": "VI",
      "ccn3": "850",
      "cca3": "VIR",
      "cioc": "ISV"
    },
    {
      "name": {
        "common": "Vietnam",
        "official": "Socialist Republic of Vietnam"
      },
      "tld": [
        ".vn"
      ],
      "cca2": "VN",
      "ccn3": "704",
      "cca3": "VNM",
      "cioc": "VIE"
    },
    {
      "name": {
        "common": "Vanuatu",
        "official": "Republic of Vanuatu"
      },
      "tld": [
        ".vu"
      ],
      "cca2": "VU",
      "ccn3": "548",
      "cca3": "VUT",
      "cioc": "VAN"
    },
    {
      "name": {
        "common": "Wallis and Futuna",
        "official": "Territory of the Wallis and Futuna Islands"
      },
      "tld": [
        ".wf"
      ],
      "cca2": "WF",
      "ccn3": "876",
      "cca3": "WLF",
      "cioc": ""
    },
    {
      "name": {
        "common": "Samoa",
        "official": "Independent State of Samoa"
      },
      "tld": [
        ".ws"
      ],
      "cca2": "WS",
      "ccn3": "882",
      "cca3": "WSM",
      "cioc": "SAM"
    },
    {
      "name": {
        "common": "Yemen",
        "official": "Republic of Yemen"
      },
      "tld": [
        ".ye"
      ],
      "cca2": "YE",
      "ccn3": "887",
      "cca3": "YEM",
      "cioc": "YEM"
    },
    {
      "name": {
        "common": "South Africa",
        "official": "Republic of South Africa"
      },
      "tld": [
        ".za"
      ],
      "cca2": "ZA",
      "ccn3": "710",
      "cca3": "ZAF",
      "cioc": "RSA"
    },
    {
      "name": {
        "common": "Zambia",
        "official": "Republic of Zambia"
      },
      "tld": [
        ".zm"
      ],
      "cca2": "ZM",
      "ccn3": "894",
      "cca3": "ZMB",
      "cioc": "ZAM"
    },
    {
      "name": {
        "common": "Zimbabwe",
        "official": "Republic of Zimbabwe"
      },
      "tld": [
        ".zw"
      ],
      "cca2": "ZW",
      "ccn3": "716",
      "cca3": "ZWE",
      "cioc": "ZIM"
    }
  ]
}`

//CountryList struct
type CountryList struct {
	Countries []Country `json:"countries"`
}

//Country struct
type Country struct {
	Name CountryName `json:"name"`
	TLD  []string    `json:"tld"`
	Cca2 string      `json:"cca2"`
	Ccn3 string      `json:"ccn3"`
	Cca3 string      `json:"cca3"`
	Cioc string      `json:"cioc"`
}

//CountryName struct
type CountryName struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

//JsonCountries list of countries
func JsonCountries() (*CountryList, error) {
	var countryList CountryList
	err := json.Unmarshal([]byte(countries), &countryList)
	if err != nil {
		return nil, err
	}

	return &countryList, err
}
