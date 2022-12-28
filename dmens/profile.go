package dmens

const (
	profileModule    = "profile"
	FunctionRegister = "register"
)

type Profile struct {
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}

//func (p *Poster) Register(profile Profile) (*types.TransactionBytes, error) {
//	profileBytes, err := json.Marshal(profile)
//	if err != nil {
//		return nil, err
//	}
//	//TODO need use wallet sdk function
//}
