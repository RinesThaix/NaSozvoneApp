package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

const (
	addonName = "SexyDKPSync"
	path      = "Interface" + string(os.PathSeparator) + "AddOns" + string(os.PathSeparator) + addonName
	luaFile   = path + string(os.PathSeparator) + addonName + ".lua"
	tocFile   = path + string(os.PathSeparator) + addonName + ".toc"
)

var prefix string

type (
	infoWriter struct{}
	warnWriter struct{}

	logger struct {
		i *log.Logger
		w *log.Logger
	}

	app struct {
		stopped chan struct{}
		logger  *logger
	}
)

func (writer infoWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("2006-01-02 15:04:05") + " [INFO] " + string(bytes))
}

func (writer warnWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("2006-01-02 15:04:05") + " [WARN] " + string(bytes))
}

func (logger *logger) info(data ...interface{}) {
	logger.i.Println(data...)
}

func (logger *logger) warn(data ...interface{}) {
	logger.w.Println(data...)
}

func initApp() *app {
	logger := &logger{
		i: &log.Logger{},
		w: &log.Logger{},
	}
	logger.i.SetFlags(0)
	logger.i.SetOutput(new(infoWriter))
	logger.w.SetFlags(0)
	logger.w.SetOutput(new(warnWriter))

	return &app{
		stopped: make(chan struct{}),
		logger:  logger,
	}
}

func Run() {
	var dir string
	var err error
	if runtime.GOOS == "windows" {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		dir, err = filepath.Abs(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(os.Args[0])))))
	}
	if err != nil {
		dir = "."
	}
	prefix = dir + string(os.PathSeparator)

	app := initApp()
	if err := testDirectory(); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		cancel()
	}()

	timer := time.NewTicker(time.Minute * 15)
	if err := rewrite(); err == nil {
		app.logger.info("Updated AddOn data on init. Sleeping for 15 minutes now.")
	} else {
		app.logger.warn(err)
	}
	go func() {
		for {
			select {
			case <-app.stopped:
				return
			case <-timer.C:
				if err := rewrite(); err == nil {
					app.logger.info("Updated AddOn data. Sleeping for 15 minutes now.")
				} else {
					app.logger.warn(err)
				}
			}
		}
	}()
	<-ctx.Done()
	app.logger.info("Stopping..")
	close(app.stopped)
}

func testDirectory() error {
	prerequirements := []string{"Interface", "Logs", "WTF"}
	for _, folder := range prerequirements {
		info, err := os.Stat(prefix + folder)
		if err != nil {
			return errors.New("there is no directory " + folder + ": " + err.Error())
		}
		if !info.IsDir() {
			return errors.New("is not directory: " + folder)
		}
	}
	if err := os.MkdirAll(prefix+path, 0777); err != nil && !os.IsExist(err) {
		return err
	}
	_, err := os.Stat(prefix + tocFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err = ioutil.WriteFile(prefix+tocFile, []byte("## Interface: 11305\n## Title: Sexy DKP Sync\n"+
			"## Notes: Аддон синхронизации внешнего мира и Sexy DKP\n"+
			"## Author: Konstantin Shandurenko\n"+
			"## Version: 1.0.0\n\n"+
			addonName+".lua\n"), 0777); err != nil {
			return err
		}
	}
	return nil
}

func rewrite() error {
	lua, err := constructLua()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(prefix+luaFile, []byte(lua), 0777)
}

func constructLua() (string, error) {
	raids, err := getRaidParties()
	if err != nil {
		return "", err
	}
	priorities, err := getLootPriorities()
	if err != nil {
		return "", err
	}
	warlocks, err := getWarlocks()
	if err != nil {
		return "", err
	}
	news, err := getNews()
	if err != nil {
		return "", err
	}
	table := mapToLua(map[string]string{
		"raids": raidPartiesToLua(raids),
		"loot":  lootPrioritiesToLua(priorities),
		"warlocks": warlocksToLua(warlocks),
		"news":  newsToLua(news),
		"ts":    strconv.Itoa(int(time.Now().Unix())),
	})
	return "SexyDKPAppData = " + table + "\n", nil
}
