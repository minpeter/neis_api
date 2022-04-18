package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Schedule() string {

	base_url := "https://open.neis.go.kr/hub/"
	sub_url := "acaInsTiInfo"
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
	var Schedule_data ScheduleData

	json.Unmarshal([]byte(data), &Schedule_data)

	string_data := Schedule_data.AcaInsTiInfo[1].Row[0].RealmScNm
	return string_data
}

type ScheduleData struct {
	AcaInsTiInfo []struct {
		Head []struct {
			ListTotalCount int `json:"list_total_count,omitempty"`
			Result         struct {
				Code    string `json:"CODE"`
				Message string `json:"MESSAGE"`
			} `json:"RESULT,omitempty"`
		} `json:"head,omitempty"`
		Row []struct {
			AtptOfcdcScCode        string      `json:"ATPT_OFCDC_SC_CODE"`
			AtptOfcdcScNm          string      `json:"ATPT_OFCDC_SC_NM"`
			AdmstZoneNm            string      `json:"ADMST_ZONE_NM"`
			AcaInstiScNm           string      `json:"ACA_INSTI_SC_NM"`
			AcaAsnum               string      `json:"ACA_ASNUM"`
			AcaNm                  string      `json:"ACA_NM"`
			EstblYmd               string      `json:"ESTBL_YMD"`
			RegYmd                 string      `json:"REG_YMD"`
			RegSttusNm             string      `json:"REG_STTUS_NM"`
			CaaBeginYmd            interface{} `json:"CAA_BEGIN_YMD"`
			CaaEndYmd              interface{} `json:"CAA_END_YMD"`
			ToforSmtot             int         `json:"TOFOR_SMTOT"`
			DtmRcptnAbltyNmprSmtot int         `json:"DTM_RCPTN_ABLTY_NMPR_SMTOT"`
			RealmScNm              string      `json:"REALM_SC_NM"`
			LeOrdNm                string      `json:"LE_ORD_NM"`
			LeCrseListNm           interface{} `json:"LE_CRSE_LIST_NM"`
			LeCrseNm               string      `json:"LE_CRSE_NM"`
			PsnbyThccCntnt         interface{} `json:"PSNBY_THCC_CNTNT"`
			ThccOthbcYn            string      `json:"THCC_OTHBC_YN"`
			BrhsAcaYn              string      `json:"BRHS_ACA_YN"`
			FaRdnzc                string      `json:"FA_RDNZC"`
			FaRdnma                string      `json:"FA_RDNMA"`
			FaRdnda                string      `json:"FA_RDNDA"`
			LoadDtm                string      `json:"LOAD_DTM"`
		} `json:"row,omitempty"`
	} `json:"acaInsTiInfo"`
}
