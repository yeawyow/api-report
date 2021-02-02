package model

type Employee struct {
	UID                uint   `json:"uid"`
	Username           string `json:"username"`
	Pass               string `json:"pass"`
	PnameID            int    `json:"pname_id"`
	Fname              string `json:"fname"`
	Lname              string `json:"lname"`
	GroupUserID        int    `json:"group_user_id"`
	HospitalDepartment int    `json:"hospital_department"`
	PositionID         int    `json:"position_id"`
	Status             string `json:"status"`
}