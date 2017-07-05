package okeros

//Rectangle is a rectangle...
type Rectangle struct {
	leftX   int
	bottomY int
	width   int
	height  int
}

func Intersection(r1, r2 Rectangle) *Rectangle {
	firstX := r1
	lastX := r2
	if r2.leftX < r1.leftX {
		firstX = r2
		lastX = r1
	}

	firstY := r1
	lastY := r2
	if r2.bottomY < r1.bottomY {
		firstY = r2
		lastY = r1
	}

	overlapX := false
	if firstX.leftX+firstX.width >= lastX.leftX {
		overlapX = true
	}

	overlapY := false
	if firstX.bottomY+firstY.height >= lastX.bottomY {
		overlapY = true
	}

	if !overlapY || !overlapX {
		return nil
	}

	i := Rectangle{}
	i.leftX = lastX.leftX
	i.bottomY = lastX.bottomY
	i.width = firstX.width - (lastX.leftX - firstX.leftX)
	i.height = firstY.height - (lastY.bottomY - firstY.bottomY)

	return &i
}
