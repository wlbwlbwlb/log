package log

import "go.uber.org/zap"

type Wrap struct {
	*zap.SugaredLogger
}

func (p *Wrap) Info(args ...interface{}) {
	p.SugaredLogger.Info(args...)
}

func (p *Wrap) Warn(args ...interface{}) {
	p.SugaredLogger.Warn(args...)
}

func (p *Wrap) Error(args ...interface{}) {
	p.SugaredLogger.Error(args...)
}

func (p *Wrap) Panic(args ...interface{}) {
	p.SugaredLogger.Panic(args...)
}

func (p *Wrap) Infof(template string, args ...interface{}) {
	p.SugaredLogger.Infof(template, args...)
}

func (p *Wrap) Warnf(template string, args ...interface{}) {
	p.SugaredLogger.Warnf(template, args...)
}

func (p *Wrap) Errorf(template string, args ...interface{}) {
	p.SugaredLogger.Errorf(template, args...)
}

func (p *Wrap) Panicf(template string, args ...interface{}) {
	p.SugaredLogger.Panicf(template, args...)
}

func (p *Wrap) With(args ...interface{}) *Wrap {
	return &Wrap{
		SugaredLogger: p.SugaredLogger.With(args...),
	}
}
