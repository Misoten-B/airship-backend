package dto

// 		string id PK
//     	references user FK
// 		references three_dimentional_model FK "3Dモデル"
// 		references business_card_parts_coordinate FK　"名刺のテンプレート"
// 		references business_card_background FK "名刺の背景"
// 		string speaking_description "喋らせたい内容"
// 		string speaking_audio_path "喋らせたい内容の音声パス"
//     	string business_card_name "名刺名"
//     	string display_name "表示名"
// 		string company_name "会社名"
// 		string department "部署"
// 		string official_position "役職"
// 		string phone_number "電話番号"
// 		string email "メールアドレス"
// 		string postal_code "郵便番号"
// 		text address "住所"
// 		string qrcode_image_path ""
//     	timestamp created_at
// 		int access_count

type BusinessCardResponse struct {
	ID                          string
	ThreeDimentionalModel       string
	BusinessCardPartsCoordinate string
	BusinessCardBackground      string
	SpeakingDescription         string
	SpeakingAudioPath           string
	BusinessCardName            string
	DisplayName                 string
	CompanyName                 string
	Department                  string
	OfficialPosition            string
	PhoneNumber                 string
	Email                       string
	PostalCode                  string
	Address                     string
	QrcodeImagePath             string
	AccessCount                 int
}
