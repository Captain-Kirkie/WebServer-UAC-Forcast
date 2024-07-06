package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Region string

// logan, ogden, uintas, salt-lake, provo, skyline, moab, abajos, southwest
const (
	Logan Region = "logan"
    Ogden Region = "ogden"
    Uintas Region = "uintas"
    SaltLake Region = "salt-lake"
    Provo Region = "provo"
    Skyline Region = "skyline"
    Moab Region = "moab"
    Abajos Region = "abajos"
    Southwest Region = "southwest"
)

var (
    regionMap = map[string]Region{
        "logan": Logan,
		"ogden": Ogden,
		"uintas": Uintas,
		"salt-lake": SaltLake,
		"provo": Provo,
		"skyline": Skyline,
		"moab": Moab,
		"abajos": Abajos,
		"southwest": Southwest,
    }
)

func validateQueryParam(str string) (Region, bool) {
	c, ok := regionMap[str]
	return c, ok
}

type Advisory struct {
    DateIssued               string `json:"date_issued"`
    DateIssuedTimestamp      string `json:"date_issued_timestamp"`
    SpecialAnnouncement      string `json:"special_announcement"`
    GeneralAnnouncements     string `json:"general_announcements"`
    CurrentConditions        string `json:"current_conditions"`
    MountainWeather          string `json:"mountain_weather"`
    RecentActivity           string `json:"recent_activity"`
    AvalancheProblem1        string `json:"avalanche_problem_1"`
    AvalancheProblem1Desc    string `json:"avalanche_problem_1_description"`
    AvalancheProblem2        string `json:"avalanche_problem_2"`
    AvalancheProblem2Desc    string `json:"avalanche_problem_2_description"`
    AvalancheProblem3        string `json:"avalanche_problem_3"`
    AvalancheProblem3Desc    string `json:"avalanche_problem_3_description"`
    BottomLine               string `json:"bottom_line"`
    DangerRose1              string `json:"danger_rose_1"`
    DangerRose2              string `json:"danger_rose_2"`
    DangerRose3              string `json:"danger_rose_3"`
    OverallDangerRose        string `json:"overall_danger_rose"`
    OverallDangerRoseImage   string `json:"overall_danger_rose_image"`
    OverallDangerRating      string `json:"overall_danger_rating"`
    Region                   string `json:"region"` 
}

type AdvisoryWrapper struct {
    Advisory Advisory `json:"advisory"`
}

type Advisories struct {
    Advisories []AdvisoryWrapper `json:"advisories"`
}

func FetchUACForcast(c *gin.Context) {
	// grab the query param
	query := c.Query("region") 
	
	region, okQuery := validateQueryParam(query)
	
	if !okQuery {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request. Use valid query param."})
	}
	// TODO: pull area into an enum
	url := "https://utahavalanchecenter.org/forecast/" + string(region) +"/json"
	
	fmt.Println("this is the url %s", url)
	
	res, err := http.Get(url)
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}
	resBody, err := io.ReadAll(res.Body)
	print(resBody)

	// Deserialization
	var apiResponse Advisories

	err = json.Unmarshal(resBody, &apiResponse)
	if err != nil {
		fmt.Println("Error deserializing:", err)
		return
	}

	fmt.Println(apiResponse)

	// body, err := base64.StdEncoding.DecodeString(string(bytes)) // response body is []byte
    // fmt.Println(body) 
	fmt.Println("fetching services", err)
	c.JSON(res.StatusCode, "You did it")
}