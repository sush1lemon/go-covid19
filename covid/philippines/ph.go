package philippines

const baseApi = "https://services5.arcgis.com/mnYJ21GiFTR97WFg/arcgis/rest/services/"
const currentStats = "slide_fig/FeatureServer/0/query?"
const hospitalPUI = "PUI_fac_tracing/FeatureServer/0/query?"
const ageGroup = "age_group/FeatureServer/0/query?"

// Params
type phParams struct {
	F              string `json:"f" url:"f"`
	Where          string `json:"where" url:"where"`
	ReturnGeometry string `json:"returnGeometry" url:"returnGeometry"`
	SpatialRel     string `json:"spatialRel" url:"spatialRel"`
	OutFields      string `json:"outFields" url:"outFields"`
	OutStatistics  string `json:"outStatistics, omitempty" url:"outStatistics, omitempty"`
	CacheHint      string `json:"cacheHint" url:"cacheHint"`
}

// Params
type outStatistics struct {
	StatisticType         string `json:"statisticType" url:"statisticType"`
	OnStatisticField      string `json:"onStatisticField" url:"onStatisticField"`
	OutStatisticFieldName string `json:"outStatisticFieldName" url:"outStatisticFieldName"`
}

// Api response
type statsBase struct {
	Features []statsFeature `json:"features"`
}

// Api response
type statsFeature struct {
	Attributes StatsAttributes `json:"attributes"`
}

// Api response
type StatsAttributes struct {
	Day       int64       `json:"day"`
	Confirmed int64       `json:"confirmed"`
	PUIs      interface{} `json:"PUIs"`
	PUMs      interface{} `json:"PUMs"`
	Recovered int64       `json:"recovered"`
	Deaths    int64       `json:"deaths"`
	Tests     interface{} `json:"tests"`
	ObjectID  int64       `json:"ObjectId"`
}

// Api response
type hsPUIBase struct {
	Features []hsPUIFeature `json:"features"`
}

// Api response
type hsPUIFeature struct {
	Attributes HsPUIsAttributes `json:"attributes"`
}

// Api response
type HsPUIsAttributes struct {
	Region    string  `json:"region"`
	HF        string  `json:"hf"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	PUIs      int64   `json:"PUIs"`
	ObjectID  int64   `json:"ObjectId"`
}
