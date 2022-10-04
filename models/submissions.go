package models

import (
	"rwby-adventures/config"
)

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
	SubmissionID string `gorm:"primary_key;column:id"`
	DiscordID    string `gorm:"primary_key;column:discord_id" json:"-"`

	Name      string `gorm:"column:name;not null" json:"name"`
	Color     string `gorm:"column:color;not null" json:"color"`
	ShortDesc string `gorm:"column:short_desc;not null" json:"short_desc"`
	LongDesc  string `gorm:"column:long_desc;not null" json:"long_desc"`
	Author    string `gorm:"column:author;not null" json:"author"`

	Icon  *SubmissionFile   `gorm:"-" json:"icon"`
	Files []*SubmissionFile `gorm:"-" json:"files"`
	Votes []*SubmissionVote `gorm:"-" json:"votes"`
}

// Many variables for ease of use
type SubmissionFile struct {
	FileID       int    `gorm:"primary_key;column:id"`
	SubmissionID string `gorm:"primary_key;column:submission_id"`

	Name string `gorm:"column:name;not null" json:"name"`
	URI  string `gorm:"column:uri;not null" json:"uri"`
	Path string `gorm:"column:path;not null" json:"path"`
}

type SubmissionVote struct {
	SubmissionID string `gorm:"primary_key;column:submission_id"`
	DiscordID    string `gorm:"primary_key;column:discord_id" json:"-"`
}

func GetAmountOfSubmissions() int {
	var count int64
	config.Database.Model(&Submission{}).Count(&count)
	return int(count)
}

func GetAmountOfVote(ID string) int {
	var count int64
	config.Database.Model(&SubmissionVote{}).Where("discord_id = ?", ID).Count(&count)
	return int(count)
}

func GetSubmission(id string) *Submission {
	s := &Submission{
		SubmissionID: id,
	}
	config.Database.Find(&s)
	s.Votes = GetSubmissionVotes(s.SubmissionID)
	s.Files = GetSubmissionFiles(s.SubmissionID)
	s.Icon = s.Files[0]
	return s
}

func GetSubmissions(amount int, offset int) []*Submission {
	var submissions []*Submission
	config.Database.Limit(amount).Offset(offset).Find(&submissions)
	for _, s := range submissions {
		s.Votes = GetSubmissionVotes(s.SubmissionID)
		s.Files = GetSubmissionFiles(s.SubmissionID)
		s.Icon = s.Files[0]
	}
	return submissions
}

func GetUserSubmissions(ID string) []*Submission {
	var submissions []*Submission
	config.Database.Where("discord_id = ?", ID).
		Find(&submissions)
	for _, s := range submissions {
		s.Votes = GetSubmissionVotes(s.SubmissionID)
		s.Files = GetSubmissionFiles(s.SubmissionID)
		s.Icon = s.Files[0]
	}
	return submissions
}

func GetSubmissionVotes(ID string) []*SubmissionVote {
	var votes []*SubmissionVote
	config.Database.Where("submission_id = ?", ID).
		Find(&votes)
	return votes
}

func GetSubmissionFiles(ID string) []*SubmissionFile {
	var files []*SubmissionFile
	config.Database.Where("submission_id = ?", ID).
		Find(&files)
	return files
}

func (s *Submission) Save() {
	config.Database.Save(s)
	for _, f := range s.Files {
		f.SubmissionID = s.SubmissionID
		config.Database.Save(f)
	}
}

func (v *SubmissionVote) Save() {
	config.Database.Save(v)
}
