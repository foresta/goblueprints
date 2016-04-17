package trace

// Tracerはコード内での出来事を記録できるオブジェクトを表すインターフェース
type Tracer interface {
	Trace(...interface{})
}
