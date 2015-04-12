package goics

/*
GetJapanTimezoneTemplate 日本用タイムゾーンテンプレート
 */
func GetJapanTimezoneTemplate() *Component {
	return &Component{
		Name: VTIMEZONE,
		Values: map[string]string{
			TZID: "Japan",
		},
		SubComponents: []*Component{
			&Component{
				Name: STANDARD,
				Values: map[string]string{
					DTSTART: "19390101T000000",
					TZOFFSETTO: "+0900",
					TZOFFSETFROM: "+0900",
					TZNAME: "JST",
				},
			},
		},
	}
}