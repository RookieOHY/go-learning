package _0_complex_json_model

/*
	goland自带json转结构体
*/
type ResponseStruct struct {
	Data []struct {
		Sence       string `json:"sence"`
		CurMetaData struct {
			Replicates int     `json:"replicates"`
			CpuUtil    float64 `json:"cpuUtil"`
		} `json:"curMetaData"`
		BaseLine struct {
			CoreUtilThreshold float64 `json:"coreUtilThreshold"`
			MaxReplicates     int     `json:"maxReplicates"`
			MinReplicates     int     `json:"minReplicates"`
			MinSupportQps     int     `json:"minSupportQps"`
			Transfer          int     `json:"transfer"`
			Interval          float64 `json:"interval"`
			XOrigin           int     `json:"x_origin"`
		} `json:"baseLine"`
		PredictSeries []int `json:"predictSeries"`
		Tab3          struct {
			Sence       string `json:"sence"`
			CurMetaData struct {
				Replicates int     `json:"replicates"`
				CpuUtil    float64 `json:"cpuUtil"`
			} `json:"curMetaData"`
			BaseLine struct {
				CoreUtilThreshold float64 `json:"coreUtilThreshold"`
				MaxReplicates     int     `json:"maxReplicates"`
				MinReplicates     int     `json:"minReplicates"`
				MinSupportQps     int     `json:"minSupportQps"`
				Transfer          int     `json:"transfer"`
				Interval          float64 `json:"interval"`
				XOrigin           int     `json:"x_origin"`
			} `json:"baseLine"`
			PredictSeries []int `json:"predictSeries"`
		} `json:"tab3"`
	} `json:"data"`
}
