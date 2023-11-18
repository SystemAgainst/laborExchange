package domain

type Vacancy struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Salary       string `json:"salary"`
	CityID       int    `json:"city_id"`
	ExperienceID int    `json:"experience_id"`
	SpecialityID int    `json:"speciality_id"`
	Description  string `json:"description"`
}
