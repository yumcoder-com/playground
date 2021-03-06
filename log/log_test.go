// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package log

import (
	"errors"
	"flag"
	"github.com/golang/glog"
	"log"
	"testing"
)

func Benchmark_Log(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Printf("log")
	}
}

func Benchmark_GLog(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		glog.Infoln("glog")
	}
}

// Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
// example: I0617 15:53:33.466007   11580 log_test.go:35] glog info message
func Test_GLog(t *testing.T) {
	//glog.CopyStandardLogTo("")
	flag.Set("alsologtostderr", "false")
	//flag.Set("stderrthreshold", "FATAL")
	flag.Set("log_dir", "/home/yumcoder")

	glog.Infoln("glog info message")
	glog.Error("glog error message")
	defer glog.Flush()
}

func f1(i int) (int32, error) {
	returnErr := func(err error) (int32, error) {
		glog.Error(err)
		return 0, err
	}

	if i > 10 {
		return returnErr(errors.New("i>10"))
	}

	glog.Info("info...")
	return 100, nil
}

func Test_log_inline_fun(t *testing.T) {
	f1(300)
}
