package request

// RegisterThrowRequest represents the request to register a throw
type RegisterThrowRequest struct {
	GameID     int `param:"game_id" description:"Game ID to register throw for"`
	FrameCount int `json:"frame_count" example:"1" description:"Frame number (1-10)"`
	ThrowCount int `json:"throw_count" example:"1" description:"Throw number within the frame (1-2, or 3 for 10th frame)"`
	Pin1       int `json:"pin_1" example:"1" description:"Pin 1 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin2       int `json:"pin_2" example:"0" description:"Pin 2 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin3       int `json:"pin_3" example:"1" description:"Pin 3 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin4       int `json:"pin_4" example:"0" description:"Pin 4 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin5       int `json:"pin_5" example:"1" description:"Pin 5 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin6       int `json:"pin_6" example:"0" description:"Pin 6 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin7       int `json:"pin_7" example:"1" description:"Pin 7 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin8       int `json:"pin_8" example:"0" description:"Pin 8 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin9       int `json:"pin_9" example:"1" description:"Pin 9 status (0=standing, 1=1st throw, 2=2nd throw)"`
	Pin10      int `json:"pin_10" example:"0" description:"Pin 10 status (0=standing, 1=1st throw, 2=2nd throw)"`
}