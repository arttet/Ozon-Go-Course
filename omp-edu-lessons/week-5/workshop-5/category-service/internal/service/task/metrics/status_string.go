// Code generated by "stringer -linecomment -type=status"; DO NOT EDIT.

package task_metrics

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusAll-1]
	_ = x[StatusNonStarted-2]
	_ = x[StatusStarted-3]
}

const _status_name = "allnon_startedstarted"

var _status_index = [...]uint8{0, 3, 14, 21}

func (i status) String() string {
	i -= 1
	if i >= status(len(_status_index)-1) {
		return "status(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _status_name[_status_index[i]:_status_index[i+1]]
}
