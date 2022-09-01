package chars

//CallMain will properly call the main function for the semblance
func (s *CharacterSemblance) CallMain(stats *CharacterStatsStruct) CharacterSemblanceUsed {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	r := s.Main(stats, s)
	return r
}

//CallAttacked will properly call the Attacked function for the semblance
func (s *CharacterSemblance) CallAttacked(stats *CharacterStatsStruct, damageDealt int) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	s.Attacked(stats, damageDealt, s)
}

//CallGotAttacked will properly call the GotAttacked function for the semblance
func (s *CharacterSemblance) CallGotAttacked(stats *CharacterStatsStruct, damageReceived int) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	s.GotAttacked(stats, damageReceived, s)
}

//CallPassive will properly call the Passive function for the semblance
func (s *CharacterSemblance) CallPassive(stats *CharacterStatsStruct) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	s.Passive(stats, s)
}
