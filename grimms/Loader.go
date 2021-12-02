package grimms

//GrimmSpeUsed represents a character's semblance after use
type GrimmSpeUsed struct {
	Damage     int
	Heal       int
	StunChance int
}

//GrimmSpe represents a character's semblance, can store data and store functions
type GrimmSpe struct {
	Main           func(*GrimmStatsStruct, *GrimmSpe) GrimmSpeUsed //Called when used
	Attacked       func(*GrimmStatsStruct, int, *GrimmSpe)         //Called when attacking
	GotAttacked    func(*GrimmStatsStruct, int, *GrimmSpe)         //Called when attacked
	Passive        func(*GrimmStatsStruct, *GrimmSpe)              //Called each turn
	Every          int
	Desc           string
	CustomFunction map[string]interface{}
	CustomData     map[string]interface{}
}

//GrimmStruct represents a character
type GrimmStruct struct {
	Name         string               `json:"Name"`
	CustomID     string               `json:"ID"`
	Skin         string               `json:"-"`
	Parent       string               `json:"-"`
	Rarity       int                  `json:"Rarity"`
	Level        int                  `json:"Level"`
	IconURL      string               `json:"IconURL"`
	WinURL       string               `json:"-"`
	ImageFile    string               `json:"-"`
	ImageAuthors string               `json:"-"`
	Stats        GrimmStatsStruct     `json:"Stats"`
	Buffs        int                  `json:"Buffs"`
	Attacks      []GrimmAttacksStruct `json:"-"`
	SpePriority  int                  `json:"-"`
	Special      GrimmSpe             `json:"-"`
	Limited      bool                 `json:"-"`
	Category     string               `json:"-"`
	NotLootable  bool                 `json:"-"`
}

//GrimmAttacksStruct represents an attack of a character
type GrimmAttacksStruct struct {
	Name       string `json:"-"`
	Speed      int    `json:"-"`
	StunChance int    `json:"-"`
	Damages    int    `json:"-"`
	Heal       int    `json:"-"`
	LastUsed   int    `json:"-"`
	Every      int    `json:"-"`
}

//GrimmStatsStruct represents the stats of a character
type GrimmStatsStruct struct {
	CharID      string  `json:"grimm_id" db:"grimm_id"`
	Value       float64 `json:"value" db:"value"`
	Health      int     `json:"health"  db:"health"`
	Armor       int     `json:"armor" db:"armor"`
	Damage      int     `json:"damage" db:"damage"`
	Healing     int     `json:"healing" db:"healing"`
	DodgeChance int     `json:"dodge_chance" db:"dodge_chance"`
}

//BaseGrimms represents all the grimms in the game
var BaseGrimms = []GrimmStruct{
	beowolf,
	boarbatusk,
	nevermore,
	nuckelavee,
	petrageist,
	queenlancer,
	regulardeathstalker,
	sabyr,
	teryx,
	wyvern,
	cenitaur,
}
