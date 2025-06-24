package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

type UserProfile struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	Username    string     `json:"username,omitempty"`
	FirstName   string     `json:"first_name,omitempty"`
	LastName    string     `json:"last_name,omitempty"`
	DisplayName string     `json:"display_name,omitempty"`
	Avatar      string     `json:"avatar,omitempty"`
	Locale      string     `json:"locale,omitempty"`
	Timezone    string     `json:"timezone,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
}

// get from user service through rest api. Enpoint: http://localhost:3000/auth/me/profile
func GetUserProfile(userID string) (*UserProfile, error) {
	req, err := http.NewRequest("GET", "http://localhost:3000/auth/user/"+userID+"/profile", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Service-Token", os.Getenv("SERVICE_TOKEN"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Check for non-2xx status code
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user profile: " + res.Status)
	}

	var profile UserProfile
	if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}
