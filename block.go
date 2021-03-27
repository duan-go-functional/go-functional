package go_functional

func Block(FuncList ...interface{}) Func {
	return func(...interface{}) interface{} {
		var arg interface{}
		for _, f := range FuncList {
			if arg == nil {
				arg = f
			} else {
				switch t := f.(type) {
				case []Func:
					arg = t
				case Func:
					var argsArr []interface{}
					switch args := arg.(type) {
					case []interface{}:
						argsArr = args
					case []Func:
						argsArr = make([]interface{}, len(args))
						for i, v := range args {
							argsArr[i] = v
						}
					case Func:
						argsArr = []interface{}{
							args,
						}
					default:
						argsArr = []interface{}{
							args,
						}
					}
					// 执行，并将结果传递到下一层
					arg = t(argsArr...)
				default:
					arg = t
				}
			}
		}

		switch t := arg.(type) {
		case []interface{}:
			for _, resOne := range t {
				switch tt := resOne.(type) {
				case Func:
					return tt()
				default:
					return tt
				}
			}
			return nil
		case Func:
			return t()
		default:
			// 如果顶层不是Func，直接返回结果
			return arg
		}
	}
}
