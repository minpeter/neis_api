package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func Time() string {

	base_url := "https://open.neis.go.kr/hub/"
	sub_url := "hisTimetable"
	query_url := "?ATPT_OFCDC_SC_CODE=T10&SD_SCHUL_CODE=7003713"
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
	var TimeInfo_data TimeDate

	json.Unmarshal([]byte(data), &TimeInfo_data)

	string_data := TimeInfo_data.HisTimetable[1].Row[0].ItrtCntnt
	characters := "1234567890."
	for _, character := range characters {
		string_data = strings.Replace(string_data, string(character), "", -1)
	}
	return strings.Replace(string_data, "<br/>", "\n", -1)

}

type TimeDate struct {
	HisTimetable []struct {
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
			Ay              string `json:"AY"`
			Sem             string `json:"SEM"`
			AllTiYmd        string `json:"ALL_TI_YMD"`
			DghtCrseScNm    string `json:"DGHT_CRSE_SC_NM"`
			OrdScNm         string `json:"ORD_SC_NM"`
			DddepNm         string `json:"DDDEP_NM"`
			Grade           string `json:"GRADE"`
			ClrmNm          string `json:"CLRM_NM"`
			ClassNm         string `json:"CLASS_NM"`
			Perio           string `json:"PERIO"`
			ItrtCntnt       string `json:"ITRT_CNTNT"`
			LoadDtm         string `json:"LOAD_DTM"`
		} `json:"row,omitempty"`
	} `json:"hisTimetable"`
}
