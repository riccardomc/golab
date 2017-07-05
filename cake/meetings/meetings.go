package meetings

import "fmt"

//Meeting represents a meeting
type Meeting struct {
	start int
	end   int
}

//MergeMeetingsBrute returns merged meetings
func MergeMeetingsBrute(meetings []Meeting) []Meeting {
	calendar := []Meeting{}
	for i := range meetings {
		calendar = insertMeeting(&calendar, meetings[i])
	}
	return calendar
}

func insertMeeting(calendar *[]Meeting, meeting Meeting) []Meeting {
	existing := false
	for i, m := range *calendar {
		overlap, first, last := overlap(meeting, m)
		if first.start == last.start && first.end == last.end {
			existing = true
			continue
		}
		if overlap {
			fmt.Println(first, last)
			(*calendar)[i].start = first.start
			(*calendar)[i].end = max(first.end, last.end)
			return insertMeeting(calendar, (*calendar)[i])
		}
	}
	if !existing {
		*calendar = append(*calendar, meeting)
	}
	return *calendar
}

func overlap(m1, m2 Meeting) (bool, Meeting, Meeting) {
	first := m1
	last := m2

	if m1.start > m2.start {
		first = m2
		last = m1
	}

	if first.end > last.start {
		return true, first, last
	}

	return false, first, last
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
