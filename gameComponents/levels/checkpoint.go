package levels

type Checkpoint struct {
	X, Y                 float64
	Index                int
	StartPoint, EndPoint bool
}

func (lvl *Level) GetCheckpointXY(checkpointIndex int) (float64, float64) {
	var activeCheckpoint Checkpoint
	for _, checkpoint := range lvl.Checkpoints {
		if checkpointIndex == 0 {
			if checkpoint.StartPoint {
				activeCheckpoint = *checkpoint
				break
			}
		} else if checkpointIndex == 9999 {
			if checkpoint.EndPoint {
				activeCheckpoint = *checkpoint
				break
			}
		} else if checkpointIndex == checkpoint.Index {
			activeCheckpoint = *checkpoint
			break
		}
	}
	return activeCheckpoint.X, activeCheckpoint.Y
}
