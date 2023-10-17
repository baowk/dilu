package utils

import "strings"

func NewSB() *StringBuilder {
	return &StringBuilder{}
}

type StringBuilder struct {
	build strings.Builder
}

func (s *StringBuilder) Append(str string) *StringBuilder {
	s.build.WriteString(str)
	return s
}

func (s *StringBuilder) AppendBytes(bs []byte) *StringBuilder {
	s.build.Write(bs)
	return s
}

func (s *StringBuilder) AppendByte(b byte) *StringBuilder {
	s.build.WriteByte(b)
	return s
}

func (s *StringBuilder) String() string {
	return s.build.String()
}
