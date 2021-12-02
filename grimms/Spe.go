package grimms

//CallMain will properly call the main function for the semblance
func (s *GrimmSpe) CallMain(stats *GrimmStatsStruct) GrimmSpeUsed {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	r := s.Main(stats, s)
	return r
}

//CallAttacked will properly call the Attacked function for the semblance
func (s *GrimmSpe) CallAttacked(stats *GrimmStatsStruct, damageDealt int) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	s.Attacked(stats, damageDealt, s)
}

//CallGotAttacked will properly call the GotAttacked function for the semblance
func (s *GrimmSpe) CallGotAttacked(stats *GrimmStatsStruct, damageReceived int) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	s.GotAttacked(stats, damageReceived, s)
}

//CallPassive will properly call the Passive function for the semblance
func (s *GrimmSpe) CallPassive(stats *GrimmStatsStruct) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	s.Passive(stats, s)
}
