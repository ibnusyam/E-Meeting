package model

type RoomDashboard struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Omzet             float64 `json:"omzet"`
	PercentageOfUsage float64 `json:"percentageOfUsage"`
}

type DashboardResponse struct {
	TotalRoom        int             `json:"totalRoom"`
	TotalVisitor     int             `json:"totalVisitor"`
	TotalReservation int             `json:"totalReservation"`
	TotalOmzet       float64         `json:"totalOmzet"`
	Rooms            []RoomDashboard `json:"rooms"`
}
