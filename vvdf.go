package vvdf

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

var tokenRegexp = regexp.MustCompile(`"([^"]*)"|(\{|\}|\S+)`)

type VDFS struct {
	// KVBuffer stores temporary [key, value] pairs
	KVBuffer [][]string
	// LastBraceKVIdx stores the KVBuffer index where the last '{' occurred
	LastBraceKVIdx int
	// BraceStack records historical KVBuffer indices for each '{' scope level
	BraceStack []int
	// ParentKeys stores the chain of block keys waiting to be mapped
	ParentKeys []string
	// ParentDepth tracks the active index pointer for ParentKeys (-1 for root)
	ParentDept int
	// void fake key
	InnerKey string
}

func NewVDFS() *VDFS {
	return &VDFS{
		KVBuffer:       [][]string{},
		LastBraceKVIdx: -1,
		BraceStack:     []int{},
		ParentKeys:     []string{},
		ParentDept:     -1,
	}
}

func (vdfs *VDFS) newSubMap(root map[string]any) {
	if vdfs.ParentDept >= 0 && vdfs.ParentKeys[vdfs.ParentDept] == vdfs.InnerKey {
		vdfs.ParentDept += 1
		vdfs.ParentKeys = append(vdfs.ParentKeys, vdfs.InnerKey)
		vdfs.BraceStack = append(vdfs.BraceStack, vdfs.LastBraceKVIdx)
	}

	vdfs.InnerKey = ""

	tempMap := root

	if vdfs.ParentDept >= 0 {

		for i := 0; i <= vdfs.ParentDept; i++ {

			t, ok := tempMap[vdfs.ParentKeys[i]].(map[string]any)

			if !ok {
				panic("newSubMap")
			}
			tempMap = t
		}
	}

	if vdfs.ParentDept == -1 {
		for _, el := range vdfs.KVBuffer {
			// fmt.Println("-1 ", el)
			tempMap[strings.ToLower(el[0])] = el[1]
		}
	}

	if vdfs.ParentDept >= 0 {
		for _, el := range vdfs.KVBuffer[vdfs.BraceStack[vdfs.ParentDept]:] {
			// fmt.Println(el)
			tempMap[strings.ToLower(el[0])] = el[1]
		}

		vdfs.KVBuffer = vdfs.KVBuffer[:vdfs.LastBraceKVIdx]

		vdfs.ParentKeys = vdfs.ParentKeys[:vdfs.ParentDept]
		vdfs.ParentDept -= 1
	}
}

func (vdfs *VDFS) newMapKey(root map[string]any) {
	if vdfs.ParentDept == 0 && vdfs.LastBraceKVIdx >= 0 {
		root[vdfs.ParentKeys[0]] = make(map[string]any)
		return
	}

	tempMap := root

	if vdfs.ParentDept > 0 {
		for i := 0; i < vdfs.ParentDept; i++ {

			t, ok := tempMap[vdfs.ParentKeys[i]].(map[string]any)
			if !ok {
				panic("newMapKey ")
			}
			tempMap = t
		}
		tempMap[vdfs.ParentKeys[vdfs.ParentDept]] = make(map[string]any)
		vdfs.BraceStack = append(vdfs.BraceStack, vdfs.LastBraceKVIdx)
	}
}

// func StringToMap(text string) (map[string]any, error) {
// 	if text == "" {
// 		return map[string]any{}, nil
// 	}

// 	start := strings.Index(text, "{")
// 	end := strings.LastIndex(text, "}")

// 	fmt.Println(end)
// 	if end == -1 {
// 		// end = len(strings.TrimSpace(text[start:]))
// 		end = len(text[start:])
// 		fmt.Println(text[start])
// 	}

// 	// if start != -1 && end != -1 && start < end {
// 	if start != -1 && end != -1 && start < end {
// 		str := strings.TrimSpace(text[start : end+1])
// 		return ParseVDFSinglePass(str)
// 	}

//		return map[string]any{}, fmt.Errorf("xxxxxxxxxx")
//	}
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
	// vdfs := VDFS{}
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

		switch strings.TrimSpace(line) {
		case "{":
			if vdfs.ParentDept >= 0 {
				vdfs.newMapKey(root)
			}
		case "}":
			// vdfs.ParentKeys = vdfs.ParentKeys[:vdfs.ParentDept-1]
			vdfs.newSubMap(root)
		default:
			// fmt.Println(line)
			st := strings.TrimSpace(line)
			tokenSlice := extractTokensWithRegexp(st)

			switch len(tokenSlice) {
			case 1:
				vdfs.InnerKey = tokenSlice[0]

				// vdfs.ParentDept += 1
				// vdfs.ParentKeys = append(vdfs.ParentKeys, tokenSlice[0])
				// vdfs.BraceStack = append(vdfs.BraceStack, vdfs.LastBraceKVIdx)
			case 2:
				vdfs.KVBuffer = append(vdfs.KVBuffer, tokenSlice)
				vdfs.LastBraceKVIdx += 1
			}
		}
	}

	if len(vdfs.KVBuffer) > 0 {
		// fmt.Printf("%+v", vdfs.KVBuffer)
		for _, el := range vdfs.KVBuffer {
			root[strings.ToLower(el[0])] = el[1]
		}
	}
	// fmt.Printf("root ---------> %+v", root)
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
