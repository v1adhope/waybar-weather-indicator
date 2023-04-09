package wttrstruct

type Data struct {
	CurrentCondition []struct {
		FeelsLikeC       string `json:"FeelsLikeC"`
		FeelsLikeF       string `json:"FeelsLikeF"`
		CloudCover       string `json:"cloudcover"`
		Humidity         string `json:"humidity"`
		LocalObsDateTime string `json:"localObsDateTime"`
		ObservationTime  string `json:"observation_time"`
		PrecipInches     string `json:"precipInches"`
		PrecipMM         string `json:"precipMM"`
		Pressure         string `json:"pressure"`
		PressureInches   string `json:"pressureInches"`
		TempC            string `json:"temp_C"`
		TempF            string `json:"temp_F"`
		UvIndex          string `json:"uvIndex"`
		Visibility       string `json:"visibility"`
		VisibilityMiles  string `json:"visibilityMiles"`
		WeatherCode      string `json:"weatherCode"`
		WeatherDesc      []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WeatherIconUrl []struct {
			Value string `json:"value"`
		} `json:"weatherIconUrl"`
		WindDir16Point string `json:"winddir16Point"`
		WindDirDegree  string `json:"winddirDegree"`
		WindspeedKmph  string `json:"windspeedKmph"`
		WindspeedMiles string `json:"windspeedMiles"`
	} `json:"current_condition"`

	NearestArea []struct {
		AreaName []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Population string `json:"population"`
		Region     []struct {
			Value string `json:"value"`
		} `json:"region"`
		WeatherURL []struct {
			Value string `json:"value"`
		} `json:"weatherUrl"`
	} `json:"nearest_area"`

	Request []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"request"`

	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination"`
			MoonPhase        string `json:"moon_phase"`
			MoonRise         string `json:"moonrise"`
			MoonSet          string `json:"moonset"`
			SunRise          string `json:"sunrise"`
			SunSet           string `json:"sunset"`
		} `json:"astronomy"`
		AVGTempC string `json:"avgtempC"`
		AVGTempF string `json:"avgtempF"`
		Date     string `json:"date"`
		Hourly   []struct {
			DewPointC        string `json:"DewPointC"`
			DewPointF        string `json:"DewPointF"`
			FeelsLikeC       string `json:"FeelsLikeC"`
			FeelsLikeF       string `json:"FeelsLikeF"`
			HeatIndexC       string `json:"HeatIndexC"`
			HeatIndexF       string `json:"HeatIndexF"`
			WindChillC       string `json:"WindChillC"`
			WindChillF       string `json:"WindChillF"`
			WindGustKmph     string `json:"WindGustKmph"`
			WindGustMiles    string `json:"WindGustMiles"`
			ChanceOfFog      string `json:"chanceoffog"`
			ChanceOfFrost    string `json:"chanceoffrost"`
			ChanceOfHighTemp string `json:"chanceofhightemp"`
			ChanceOfOverCast string `json:"chanceofovercast"`
			ChanceOfRain     string `json:"chanceofrain"`
			ChanceOfRemDry   string `json:"chanceofremdry"`
			ChanceOfSnow     string `json:"chanceofsnow"`
			ChanceOfSunshine string `json:"chanceofsunshine"`
			ChanceOfThunder  string `json:"chanceofthunder"`
			ChanceOfWindy    string `json:"chanceofwindy"`
			CloudCover       string `json:"cloudcover"`
			Humidity         string `json:"humidity"`
			PrecipInches     string `json:"precipInches"`
			PrecipMM         string `json:"precipMM"`
			Pressure         string `json:"pressure"`
			PressureInches   string `json:"pressureInches"`
			TempC            string `json:"tempC"`
			TempF            string `json:"tempF"`
			Time             string `json:"time"`
			UVIndex          string `json:"uvIndex"`
			Visibility       string `json:"visibility"`
			VisibilityMiles  string `json:"visibilityMiles"`
			WeatherCode      string `json:"weatherCode"`
			WeatherDesc      []struct {
				Value string `json:"value"`
			} `json:"weatherDesc"`
			WeatherIconUrl []struct {
				Value string `json:"value"`
			} `json:"weatherIconUrl"`
			WindDir16Point string `json:"winddir16Point"`
			WindDirDegree  string `json:"winddirDegree"`
			WindSpeedKmph  string `json:"windspeedKmph"`
			WindSpeedMiles string `json:"windspeedMiles"`
		} `json:"hourly"`
		MaxTempC    string `json:"maxtempC"`
		MaxTempF    string `json:"maxtempF"`
		MinTempC    string `json:"mintempC"`
		MinTempF    string `json:"mintempF"`
		SunHour     string `json:"sunHour"`
		TotalSnowCM string `json:"totalSnow_cm"`
		UVIndex     string `json:"uvIndex"`
	} `json:"weather"`
}
