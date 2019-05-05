package fsmchannelwait

import "text/template"

var tpl = template.Must(template.New("fsmchannelwait").Parse(`
func {{.StructName}}ChannelWait(ctx context.Context, ch ffsm.Channel) (chan {{.StructName}}) {
	done := make(chan {{.StructName}}, 1)
	if ctx.Err() != nil {
		close(done)
		return done
	}

	go func() {
		select {
		case <-ctx.Done():
			close(done)
			return
		case msg := <-ch:
			switch msg := msg.(type) {
			case {{.StructName}}:
				done <- msg
			default:
				close(done)
			}
		}
	}()
	return done
}`))
