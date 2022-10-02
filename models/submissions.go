package models

// {
// 	name: "Submission Name",
// 	color: "Color",
// 	short_desc: "This is a short description of the submission.",
// 	long_desc: "This is a long description of the submission.",
// 	icon: "https://img.rwbyadventures.com/Janina_Wolf/Xmas_Icon.webp",
// 	author: "@Someone#0000",
// 	votes: 51,
// 	files:[
// 		{
// 			name:"Image name.png",
// 			uri:"image file.png"
// 		}
// 	],
// },

type Submission struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	Name      string           `gorm:"column:name;not null" json:"name"`
	Color     string           `gorm:"column:color;not null" json:"color"`
	ShortDesc string           `gorm:"column:short_desc;not null" json:"short_desc"`
	LongDesc  string           `gorm:"column:long_desc;not null" json:"long_desc"`
	Icon      string           `gorm:"column:icon;not null" json:"icon"`
	Author    string           `gorm:"column:author;not null" json:"author"`
	Votes     int              `gorm:"column:votes;not null" json:"votes"`
	Files     []SubmissionFile `gorm:"column:files;not null" json:"files"`
}

type SubmissionFile struct {
	ID   string `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name;not null" json:"name"`
	URI  string `gorm:"column:uri;not null" json:"uri"`
}
