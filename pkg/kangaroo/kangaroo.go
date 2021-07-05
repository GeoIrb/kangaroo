package kangaroo

var (
	allCoordinate = new(int)
)

type Kangaroo struct {
	utils utils
	x     int
	v     int
}

// IsIntersect returns true if kangaroo intersects with other kangaroo
func (k *Kangaroo) IsIntersect(other *Kangaroo) bool {
	step, coordinate := k.getIntersectionPoint(other)
	if step == nil {
		return coordinate == allCoordinate
	}

	return k.alongWay(*coordinate) && other.alongWay(*coordinate)
}

func (k *Kangaroo) getIntersectionPoint(other *Kangaroo) (*int, *int) {
	if k.v == other.v {
		if k.x != other.x {
			return nil, nil
		}
		return nil, allCoordinate
	}

	stepTmp := float64(k.x-other.x) / float64(other.v-k.v)
	if !k.utils.IsFloatInt(stepTmp) || stepTmp < 0 {
		return nil, nil
	}

	step := int(stepTmp)
	coordinate := k.x + k.v*step
	return &step, &coordinate
}

func (k *Kangaroo) alongWay(coordinate int) bool {
	return (coordinate == k.x) || (coordinate < k.x && k.v < 0) || (coordinate > k.x && k.v > 0)
}
