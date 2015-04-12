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

/*
GetICSTemplateJapan タイムゾーン日本で必要な情報をあらかじめ入れた状態で返す
 */
func GetICSTemplateJapan(comps... *Component) *ICS {
	return &ICS{
		RootComponent: &Component{
			Name: VCALENDAR,
			Values: map[string]string{
				PRODID: "-//mk2//ICSMailer//EN",
				VERSION: "2.0",
				METHOD: "PUBLISH",
			},
			SubComponents:
				append([]*Component{GetJapanTimezoneTemplate()}, comps...),
		},
	}
}