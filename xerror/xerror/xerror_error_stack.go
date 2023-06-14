package xerror

import (
	"bytes"
	"container/list"
	"fmt"
	"runtime"
	"strings"
)

// stackInfo 管理特定错误的堆栈信息。
type stackInfo struct {
	Index   int        // Index 整个错误堆栈中当前错误的索引。
	Message string     // Error 信息字符串。
	Lines   *list.List // Lines 按顺序包含当前错误堆栈的所有错误堆栈行。
}

// stackLine 管理堆栈的每一行信息。
type stackLine struct {
	Function string // Function 包含其完整包路径。
	FileLine string // FileLine 函数的源文件名及其行号。
}

// Stack 以字符串形式返回错误堆栈信息。
func (err *Error) Stack() string {
	if err == nil {
		return ""
	}
	var (
		loop  = err
		index = 1
		infos []*stackInfo
	)
	for loop != nil {
		info := &stackInfo{
			Index:   index,
			Message: fmt.Sprintf("%-v", loop),
		}
		index++
		infos = append(infos, info)
		loopLinesOfStackInfo(loop.stack, info)
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				loop = e
			} else {
				infos = append(infos, &stackInfo{
					Index:   index,
					Message: loop.error.Error(),
				})
				index++
				break
			}
		} else {
			break
		}
	}
	filterLinesOfStackInfos(infos)
	return formatStackInfos(infos)
}

// filterLinesOfStackInfos 从顶部错误中删除存在于后续堆栈中的重复行。
func filterLinesOfStackInfos(infos []*stackInfo) {
	var (
		ok      bool
		set     = make(map[string]struct{})
		info    *stackInfo
		line    *stackLine
		removes []*list.Element
	)
	for i := len(infos) - 1; i >= 0; i-- {
		info = infos[i]
		if info.Lines == nil {
			continue
		}
		for n, e := 0, info.Lines.Front(); n < info.Lines.Len(); n, e = n+1, e.Next() {
			line = e.Value.(*stackLine)
			if _, ok = set[line.FileLine]; ok {
				removes = append(removes, e)
			} else {
				set[line.FileLine] = struct{}{}
			}
		}
		if len(removes) > 0 {
			for _, e := range removes {
				info.Lines.Remove(e)
			}
		}
		removes = removes[:0]
	}
}

// formatStackInfos 格式化并以字符串形式返回错误堆栈信息。
func formatStackInfos(infos []*stackInfo) string {
	var buffer = bytes.NewBuffer(nil)
	for i, info := range infos {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, info.Message))
		if info.Lines != nil && info.Lines.Len() > 0 {
			formatStackLines(buffer, info.Lines)
		}
	}
	return buffer.String()
}

// formatStackLines 格式化并将错误堆栈行作为字符串返回。
func formatStackLines(buffer *bytes.Buffer, lines *list.List) string {
	var (
		line   *stackLine
		space  = "  "
		length = lines.Len()
	)
	for i, e := 0, lines.Front(); i < length; i, e = i+1, e.Next() {
		line = e.Value.(*stackLine)
		// Graceful indent.
		if i >= 9 {
			space = " "
		}
		buffer.WriteString(fmt.Sprintf(
			"   %d).%s%s\n        %s\n",
			i+1, space, line.Function, line.FileLine,
		))
	}
	return buffer.String()
}

// loopLinesOfStackInfo 迭代堆栈信息行并生成堆栈行信息。
func loopLinesOfStackInfo(st stack, info *stackInfo) {
	if st == nil {
		return
	}
	for _, p := range st {
		if fn := runtime.FuncForPC(p - 1); fn != nil {
			file, line := fn.FileLine(p - 1)
			if IsUsingBriefStack {
				// 过滤整个 G 包堆栈路径。
				if strings.Contains(file, stackFilterKeyForG) {
					continue
				}
			} else {
				// 包路径堆栈过滤。
				if strings.Contains(file, stackFilterKeyLocal) {
					continue
				}
			}
			// 避免堆栈字符串如 `autogenerated`
			if strings.Contains(file, "<") {
				continue
			}
			// 忽略 GO ROOT 路径。
			if goRootForFilter != "" &&
				len(file) >= len(goRootForFilter) &&
				file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if info.Lines == nil {
				info.Lines = list.New()
			}
			info.Lines.PushBack(&stackLine{
				Function: fn.Name(),
				FileLine: fmt.Sprintf(`%s:%d`, file, line),
			})
		}
	}
}
