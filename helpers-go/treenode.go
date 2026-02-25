package helpersgo

import (
	"bytes"
	"encoding/json"
)

func (l *TreeNode) JSON() string {
	b, _ := json.Marshal(l)
	return string(b)
}

func (l *TreeNode) JSONIndent() string {
	var i bytes.Buffer
	json.Indent(&i, []byte(l.JSON()), "", "\t")
	return i.String()
}
