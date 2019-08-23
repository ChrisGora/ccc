// dotWriter partially adapted from https://github.com/hirokidaichi/goviz

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type dotGenerator struct {
	nextFrame int
	mutex     sync.Mutex
}

type position struct {
	x, y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func escape(s string) string {
	return strconv.Quote(s)
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// export takes the state of all accounts and puts it in a graphviz dot file
// It should ONLY ever be called inside a locked bank.mutex !!!
func (gen *dotGenerator) export(bank *bank) {
	gen.mutex.Lock()
	//fmt.Println("DOT-A")
	file, ioError := os.Create(fmt.Sprintf("out/%010d.gv", gen.nextFrame))
	gen.nextFrame++
	check(ioError)
	//fmt.Println("DOT-B")
	_, _ = fmt.Fprintf(file, "digraph bank {\n")
	//fmt.Println("DOT-C")
	accounts := len(bank.accounts)
	//fmt.Println("DOT-D")
	var positions []position
	//fmt.Println("DOT-E")
	if accounts == 6 {
		positions = []position{
			{1, 0},
			{0, 1},
			{1, 2},
			{2, 2},
			{3, 1},
			{2, 0},
		}
	} else {
		cols := 2
		rows := accounts / cols
		for i := 0; i < cols; i++ {
			for j := 0; j < rows; j++ {
				positions = append(positions, position{i, j})
			}
		}

	}

	//fmt.Println("DOT-F")
	for i, p := range positions {
		accountNumber := i
		a := bank.accounts[accountNumber]
		_, _ = fmt.Fprintln(file, "\t", strconv.Itoa(accountNumber), "[label=", escape(a.name+"\n"+a.lockedBy), ", pos=", escape(strconv.Itoa(p.x)+","+strconv.Itoa(p.y)+"!"), ", shape=circle, fixedsize=true, width=0.5", "]")
		//fmt.Println("DOT-G")
	}
	element := bank.transactionsInProgress.Front()
	for element != nil {
		//fmt.Println("DOT-H")
		t := element.Value.(transaction)
		//fmt.Println("DOT-I", t)
		_, _ = fmt.Fprintln(file, "\t", escape(strconv.Itoa(t.from)), "->", escape(strconv.Itoa(t.to)), "[label=", escape(t.handledBy), "]")
		//fmt.Println("DOT-J")
		element = element.Next()
		//fmt.Println("DOT-K-", element)
	}

	_, _ = fmt.Fprintf(file, "}\n")

	err := file.Sync()
	check(err)
	err = file.Close()
	check(err)

	gen.mutex.Unlock()
}

func newGenerator() *dotGenerator {
	_ = os.Mkdir("out", os.ModePerm)
	_ = RemoveContents("out")
	return &dotGenerator{}
}
