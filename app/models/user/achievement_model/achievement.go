package achievement_model

import "sport_kld/app/models"

type UserAchievement struct {
	UID models.UID `json:"uid"`
	UserUID models.UID `json:"userUid"`
	Name string `json:"name"`
	Description string `json:"description"`
	AchievementDate string `json:"achievementDate"`
}
