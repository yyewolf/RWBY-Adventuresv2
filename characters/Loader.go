package chars

//CharacterSemblanceUsed represents a character's semblance after use
type CharacterSemblanceUsed struct {
	Damage     int
	Heal       int
	StunChance int
}

//CharacterSemblance represents a character's semblance, can store data and store functions
type CharacterSemblance struct {
	Main           func(*CharacterStatsStruct, *CharacterSemblance) CharacterSemblanceUsed //Called when used
	Attacked       func(*CharacterStatsStruct, int, *CharacterSemblance)                   //Called when attacking
	GotAttacked    func(*CharacterStatsStruct, int, *CharacterSemblance)                   //Called when attacked
	Passive        func(*CharacterStatsStruct, *CharacterSemblance)                        //Called each turn
	Every          int
	Desc           string
	CustomFunction map[string]interface{}
	CustomData     map[string]interface{}
}

//CharacterStruct represents a character
type CharacterStruct struct {
	Name              string                   `json:"Name"`
	CustomID          string                   `json:"ID"`
	Skin              string                   `json:"-"`
	Parent            string                   `json:"-"`
	Weapon            string                   `json:"-"`
	Rarity            int                      `json:"Rarity"`
	Level             int                      `json:"Level"`
	IconURL           string                   `json:"IconURL"`
	WinURL            string                   `json:"-"`
	ImageFile         string                   `json:"-"`
	ImageAuthors      string                   `json:"-"`
	Stats             CharacterStatsStruct     `json:"Stats"`
	Buffs             int                      `json:"Buffs"`
	Attacks           []CharacterAttacksStruct `json:"-"`
	SemblancePriority int                      `json:"-"`
	Semblance         CharacterSemblance       `json:"-"`
	Limited           bool                     `json:"-"`
	Category          string                   `json:"-"`
	NotLootable       bool                     `json:"-"`
}

//CharacterAttacksStruct represents an attack of a character
type CharacterAttacksStruct struct {
	Name       string `json:"-"`
	Speed      int    `json:"-"`
	StunChance int    `json:"-"`
	Damages    int    `json:"-"`
	Heal       int    `json:"-"`
	LastUsed   int    `json:"-"`
	Every      int    `json:"-"`
}

//CharacterStatsStruct represents the stats of a character
type CharacterStatsStruct struct {
	CharID      string  `json:"char_id" db:"char_id"`
	Value       float64 `json:"value" db:"value"`
	Health      int     `json:"health"  db:"health"`
	Armor       int     `json:"armor" db:"armor"`
	Damage      int     `json:"damage" db:"damage"`
	Healing     int     `json:"healing" db:"healing"`
	DodgeChance int     `json:"dodge_chance" db:"dodge_chance"`
}

//BaseCharacters represents all the characters in the game
var BaseCharacters = []CharacterStruct{
	blakeBelladonna,
	lieRen,
	rubyRose,
	weissSchnee,
	noraValkyrie,
	pyrrhaNikos,
	cocoAdel,
	foxAlistair,
	velvetScarlatina,
	yatsuhashiDaichi,
	pennyPolendina,
	kobalt,
	flyntCoal,
	neonKatt,
	jauneArc,
	yangXiaoLong,
	neptuneVasilias,
	sunWukong,
	scarletDavid,
	ivori,
	nebulaViolette,
	dewGayl,
	octaviaEmber,
	sageAyana,
	gwenDarcy,
	arslanAltan,
	bolinHori,
	cardinWinchester,
	cloverEbi,
	doveBronzewing,
	elmEderne,
	harrietBree,
	marrowAmin,
	nadirShiko,
	reeseChloris,
	russelThrush,
	skyLark,
	vineZeki,
	brawnzNi,
	mayZedong,
	nolanPorfirio,
	royStallion,
	robynHill,
	fionaThyme,
	mayMarigold,
	joannaGreenleaf,
	cielSoleil,
	iliaAmitola,
	melanieMalachite,
	miltiadesMalachite,
	winterSchnee,
	whitleySchnee,
	corsacAlbain,
	fennecAlbain,
	adrianCottaArc,
	amber,
	glyndaGoodwitch,
	higanbanaWaitress,
	lilMissMalachite,
	saphronCottaArc,
	terraCottaArc,
	trifa,
	bartholomewOobleck,
	peterPort,
	jamesIronwood,
	anRen,
	kaliBelladonna,

	//SPECIAL CHAR
	cinderFall,
	emeraldSustrai,
	mercuryBlack,
	neopolitan,
	romanTorchwick,
	tyrianCallows,
	arthurWatts,
	tock,
	hazelRainart,
	adamTaurus,
	siennaKhan,

	//STRQ
	summerRose,
	taiyangXiaoLong,
	ravenBranwen,
	qrowBranwen,

	//OC Contest1
	bolvaWolfmoth,
	janinaWolf,
	linAiren,
	zonBi,

	//OC Contest2
	ceriseHanBrunswick,
	lucineDhimme,
	mardiGras,
	sunsetPhoe,

	//OZMA SALEM
	ozma,
	youngSalem,

	//NOT LOOTABLE
	Zwei,
	RubyRoseBirthday,

	//SUMMER EDITION
	arslanAltanSummer,
	blakeBelladonnaSummer,
	cocoAdelSummer,
	dewGaylSummer,
	emeraldSustraiSummer,
	gwenDarcySummer,
	iliaAmitolaSummer,
	melanieMalachiteSummer,
	miltiadesMalachiteSummer,
	nebulaVioletteSummer,
	neonKattSummer,
	noraValkyrieSummer,
	octaviaEmberSummer,
	pennyPolendinaSummer,
	pyrrhaNikosSummer,
	reeseChlorisSummer,
	rubyRoseSummer,
	siennaKhanSummer,
	velvetScarlatinaSummer,
	weissSchneeSummer,
	winterSchneeSummer,
	yangXiaoLongSummer,

	//SUMMER EDITION V2
	amberSummer,
	anRenSummer,
	kaliBelladonnaSummer,
	neopolitanSummer,
	janinaWolSummer,
	blakeBelladonnaSummerV2,
	cielSoleilSummer,
	cinderFallSummer,
	fionaThymeSummer,
	glyndaGoodwitchSummer,
	pyrrhaNikosSummerV2,
	rubyRoseSummerV2,
	trifaSummer,
	velvetScarlatinaSummerV2,
	weissSchneeSummerV2,
	winterSchneeSummerV2,
	yangXiaoLongSummerV2,

	//HALLOWEEN EDITION
	blakeBelladonnaHalloween,
	janinaWolfHalloween,
	jauneArcHalloween,
	lieRenHalloween,
	neptuneVasiliasHalloween,
	noraValkyrieHalloween,
	pennyPolendinaHalloween,
	pyrrhaNikosHalloween,
	rubyRoseHalloween,
	sunWukongHalloween,
	weissSchneeHalloween,
	yangXiaoLongHalloween,

	//HALLOWEEN EDITION 2
	blakeBelladonnaHalloweenV2,
	neopolitanHalloweenV2,
	noraValkyrieHalloweenV2,
	pennyPolendinaHalloweenV2,
	pyrrhaNikosHalloweenV2,
	rubyRoseHalloweenV2,
	velvetScarlatinaHalloweenV2,
	weissSchneeHalloweenV2,
	yangXiaoLongHalloweenV2,

	//XMAS EDITION
	blakeBelladonnaXmas,
	janinaWolfXmas,
	lieRenXmas,
	melanieMalachiteXmas,
	miltiadesMalachiteXmas,
	neopolitanXmas,
	noraValkyrieXmas,
	pennyPolendinaXmas,
	pyrrhaNikosXmas,
	rubyRoseXmas,
	sunWukongXmas,
	velvetScarlatinaXmas,
	weissSchneeXmas,
	yangXiaoLongXmas,

	//Chinese New Year
	janinaChineseNewYear,
	rubyRoseChineseNewYear,
	cinderChineseNewYear,
	neopolitanChineseNewYear,
	blakeBelladonnaChineseNewYear,
	weissSchneeChineseNewYear,
	noraValkyrieChineseNewYear,
	pyrrhaNikosChineseNewYear,
	yangXiaoLongChineseNewYear,
	iliaAmitolaChineseNewYear,
	velvetScarlatinaChineseNewYear,
	pennyPolendinaChineseNewYear,

	//Valentine's Day
	pyrrhaNikosValentinesDay,
	rubyRoseValentinesDay,
	noraValkyrieValentinesDay,

	//Easter
	blakeBelladonnaEaster,
	cielSoleilEaster,
	cocoAdelEaster,
	janinaEaster,
	neopolitanEaster,
	noraValkyrieEaster,
	pennyPolendinaEaster,
	pyrrhaNikosEaster,
	rubyRoseEaster,
	velvetScarlatinaEaster,
	weissSchneeEaster,
	winterSchneeEaster,
	yangXiaoLongEaster,

	//OTHER EVENTS
	mariaCalavera,
	YangXiaoLongBirthday,
}
