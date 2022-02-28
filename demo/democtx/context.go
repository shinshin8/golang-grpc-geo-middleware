package democtx

import "context"

const (
	demoKey = "demo-key"
)

func SetDemo(parent context.Context, d string) context.Context {
	return context.WithValue(parent, demoKey, d)
}

func GetDemo(ctx context.Context) string {
	v := ctx.Value(demoKey)

	demo, ok := v.(string)
	if !ok {
		return ""
	}
	return demo
}
