package loger

import (
	//"strings"
	"fmt"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Loger struct {
	id      string
	loginId int
	service string
}

func NewLogController() *Loger {
	logControllerPointer := &Loger{
		id: uuid.New().String(),
	}
	return logControllerPointer
}

func Initialize(loginId int, service string) *Loger {
	logControllerPointer := &Loger{
		id:      uuid.New().String(),
		loginId: loginId,
		service: service,
	}
	return logControllerPointer
}

func (l Loger) Info(message string) (err error) {
	log.Info(message)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: InfoLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	err = createLog(log)
	if err != nil {

		return
	}
	return
}

func (l Loger) Debug(message string) (err error) {
	log.Debug(message)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: DebugLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) Error(message string) (err error) {
	log.Error(message)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: ErrorLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) Infof(format string, args interface{}) (err error) {
	log.Infof(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: InfoLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) Debugf(format string, args interface{}) (err error) {
	log.Debugf(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: DebugLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) Errorf(format string, args interface{}) (err error) {
	log.Errorf(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: ErrorLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) LogInfo(message string, loginId int) (err error) {
	log.Info(message)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: InfoLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	if loginId != 0 {
		log.LoginId = loginId
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) LogDebug(message string, loginId int) (err error) {
	log.Debug(message)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: DebugLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	if loginId != 0 {
		log.LoginId = loginId
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) LogError(message string, loginId int) (err error) {
	log.Error(message)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: ErrorLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	if loginId != 0 {
		log.LoginId = loginId
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) LogInfof(format string, args interface{}, loginId int) (err error) {
	log.Infof(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: InfoLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	if loginId != 0 {
		log.LoginId = loginId
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) LogDebugf(format string, args interface{}, loginId int) (err error) {
	log.Debugf(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: DebugLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	if loginId != 0 {
		log.LoginId = loginId
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) LogDebugMoreArgsf(loginId int, format string, args ...interface{}) (err error) {
	log.Debugf(format, args...)
	message := fmt.Sprintf(format, args...)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: DebugLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	if loginId != 0 {
		log.LoginId = loginId
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}

func (l Loger) LogErrorf(format string, args interface{}, loginId int) (err error) {
	log.Errorf(format, args)
	message := fmt.Sprintf(format, args)
	var log = logModel{
		SeriesId: l.id,
		LoginId:  l.loginId,
		LogLevel: ErrorLevel,
		Service:  l.service,
		Message:  message,
		DateTime: time.Now(),
	}
	if loginId != 0 {
		log.LoginId = loginId
	}
	err = createLog(log)
	if err != nil {
		return
	}
	return
}
