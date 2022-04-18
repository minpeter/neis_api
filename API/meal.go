package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func Meal() string {

	base_url := "https://open.neis.go.kr/hub/"
	sub_url := "mealServiceDietInfo"
	query_url := "?ATPT_OFCDC_SC_CODE=T10&SD_SCHUL_CODE=9296071"
	query_type := "&Type=json"
	query_pSize := "&pSize=1"
	api_key := "7976bde8f0664cbbac1e76a1779062e3"
	URL := base_url + sub_url + query_url + query_type + query_pSize + "&KEY=" + api_key
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var MealInfo_data MealDate

	json.Unmarshal([]byte(data), &MealInfo_data)

	string_data := MealInfo_data.MealServiceDietInfo[1].Row[0].DdishNm
	characters := "1234567890."
	for _, character := range characters {
		string_data = strings.Replace(string_data, string(character), "", -1)
	}
	return strings.Replace(string_data, "<br/>", "\n", -1)

}

type MealDate struct {
	MealServiceDietInfo []struct {
		Head []struct {
			ListTotalCount int `json:"list_total_count,omitempty"`
			Result         struct {
				Code    string `json:"CODE"`
				Message string `json:"MESSAGE"`
			} `json:"RESULT,omitempty"`
		} `json:"head,omitempty"`
		Row []struct {
			AtptOfcdcScCode string `json:"ATPT_OFCDC_SC_CODE"`
			AtptOfcdcScNm   string `json:"ATPT_OFCDC_SC_NM"`
			SdSchulCode     string `json:"SD_SCHUL_CODE"`
			SchulNm         string `json:"SCHUL_NM"`
			MmealScCode     string `json:"MMEAL_SC_CODE"`
			MmealScNm       string `json:"MMEAL_SC_NM"`
			MlsvYmd         string `json:"MLSV_YMD"`
			MlsvFgr         string `json:"MLSV_FGR"`
			DdishNm         string `json:"DDISH_NM"`
			OrplcInfo       string `json:"ORPLC_INFO"`
			CalInfo         string `json:"CAL_INFO"`
			NtrInfo         string `json:"NTR_INFO"`
			MlsvFromYmd     string `json:"MLSV_FROM_YMD"`
			MlsvToYmd       string `json:"MLSV_TO_YMD"`
		} `json:"row,omitempty"`
	} `json:"mealServiceDietInfo"`
}
