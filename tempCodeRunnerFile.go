	var i InfoMessage
	i.TrainingType = s.TrainingType
	i.Duration = s.Duration
	i.Distance = s.distance()
	i.Speed = s.meanSpeed()
	i.Calories = s.Calories()
	return i