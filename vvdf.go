package vvdf

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

var tokenRegexp = regexp.MustCompile(`"([^"]*)"|(\{|\}|\S+)`)

type VDFS struct {
	KVBuffer   [][]string
	BraceStack []int
	ParentKeys []string
	ParentDept int
	InnerKey   string
}

func NewVDFS() *VDFS {
	return &VDFS{
		KVBuffer:   [][]string{},
		BraceStack: []int{},
		ParentKeys: []string{},
		ParentDept: -1,
	}
}

// 沿 ParentKeys 链路找到当前层级的 target map
func (vdfs *VDFS) getMapAtDept(root map[string]any, dept int) map[string]any {
	tempMap := root
	for i := 0; i <= dept; i++ {
		pk := vdfs.ParentKeys[i]
		if pk == "" {
			continue
		}
		t, ok := tempMap[pk].(map[string]any)
		if !ok {
			t = make(map[string]any)
			tempMap[pk] = t
		}
		tempMap = t
	}
	return tempMap
}

func (vdfs *VDFS) enterBlock(root map[string]any) {
	key := strings.ToLower(vdfs.InnerKey)
	vdfs.InnerKey = ""

	if vdfs.ParentDept == -1 {
		if key == "" {
			vdfs.ParentKeys = append(vdfs.ParentKeys, "")
		} else {
			root[key] = make(map[string]any)
			vdfs.ParentKeys = append(vdfs.ParentKeys, key)
		}
	} else {
		if key == "" {
			key = "unnamed"
		}

		parentMap := vdfs.getMapAtDept(root, vdfs.ParentDept)
		newMap := make(map[string]any)
		parentMap[key] = newMap
		vdfs.ParentKeys = append(vdfs.ParentKeys, key)
	}

	vdfs.ParentDept++
	vdfs.BraceStack = append(vdfs.BraceStack, len(vdfs.KVBuffer))
}

func (vdfs *VDFS) leaveBlock(root map[string]any) {
	if vdfs.ParentDept < 0 {
		return
	}

	targetMap := vdfs.getMapAtDept(root, vdfs.ParentDept)
	startIdx := vdfs.BraceStack[vdfs.ParentDept]

	for _, el := range vdfs.KVBuffer[startIdx:] {
		targetMap[strings.ToLower(el[0])] = el[1]
	}

	vdfs.KVBuffer = vdfs.KVBuffer[:startIdx]
	vdfs.ParentKeys = vdfs.ParentKeys[:vdfs.ParentDept]
	vdfs.BraceStack = vdfs.BraceStack[:vdfs.ParentDept]
	vdfs.ParentDept--
}

func StringToMap(text string) (map[string]any, error) {
	if text == "" {
		return map[string]any{}, nil
	}

	start := strings.Index(text, "{")
	if start == -1 {
		return map[string]any{}, fmt.Errorf("invalid VDF: missing '{'")
	}

	end := strings.LastIndex(text, "}")
	if end == -1 || end < start {
		end = len(text) - 1
	}

	str := strings.TrimSpace(text[start : end+1])
	return ParseVDFSinglePass(str)
}

func ParseVDFSinglePass(text string) (map[string]any, error) {
	vdfs := NewVDFS()
	root := make(map[string]any)
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := scanner.Text()

		inQuote := false
		for i := 0; i < len(line); i++ {
			if line[i] == '"' {
				inQuote = !inQuote
			} else if !inQuote && i+1 < len(line) && line[i] == '/' && line[i+1] == '/' {
				line = line[:i]
				break
			}
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		switch trimmed {
		case "{":
			vdfs.enterBlock(root)
		case "}":
			vdfs.leaveBlock(root)
		default:
			tokenSlice := extractTokensWithRegexp(trimmed)

			switch len(tokenSlice) {
			case 1:
				vdfs.InnerKey = tokenSlice[0]
			case 2:
				vdfs.KVBuffer = append(vdfs.KVBuffer, tokenSlice)
			}
		}
	}

	if len(vdfs.KVBuffer) > 0 {
		for _, el := range vdfs.KVBuffer {
			root[strings.ToLower(el[0])] = el[1]
		}
	}

	return root, nil
}

func extractTokensWithRegexp(line string) []string {
	matches := tokenRegexp.FindAllStringSubmatch(line, -1)
	tokens := make([]string, 0, len(matches))

	for _, m := range matches {
		var tok string
		if strings.HasPrefix(m[0], `"`) {
			tok = m[1]
		} else {
			tok = m[2]
		}

		tok = strings.TrimSpace(tok)
		if tok != "" || strings.HasPrefix(m[0], `"`) {
			tokens = append(tokens, tok)
		}
	}
	return tokens
}
