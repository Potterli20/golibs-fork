package netutil_test

import (
	"fmt"

	"github.com/AdguardTeam/golibs/netutil"
)

func ExampleJoinHostPort() {
	fmt.Println(netutil.JoinHostPort("example.com", 12345))

	// Output:
	//
	// example.com:12345
}

func ExampleSplitHostPort() {
	host, port, err := netutil.SplitHostPort("example.com:12345")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T(%[1]v)\n", host)
	fmt.Printf("%T(%[1]v)\n", port)

	// Output:
	//
	// string(example.com)
	// int(12345)
}

func ExampleSplitHost() {
	host, err := netutil.SplitHost("example.com:12345")
	if err != nil {
		panic(err)
	}

	fmt.Println(host)

	host, err = netutil.SplitHost("example.org")
	if err != nil {
		panic(err)
	}

	fmt.Println(host)

	_, err = netutil.SplitHost("[BAD:!")
	fmt.Println(err)

	// Output:
	//
	// example.com
	// example.org
	// address [BAD:!: missing ']' in address
}

func ExampleSubdomains() {
	fmt.Printf("%#v\n", netutil.Subdomains("subsub.sub.domain.tld"))

	fmt.Println()

	fmt.Printf("%#v\n", netutil.Subdomains(""))

	// Output:
	//
	// []string{"subsub.sub.domain.tld", "sub.domain.tld", "domain.tld", "tld"}
	//
	// []string(nil)
}

func ExampleIsSubdomain() {
	fmt.Printf("%-14s: %5t\n", "same domain", netutil.IsSubdomain("sub.example.com", "example.com"))
	fmt.Printf("%-14s: %5t\n", "not immediate", netutil.IsSubdomain("subsub.sub.example.com", "example.com"))

	fmt.Printf("%-14s: %5t\n", "empty", netutil.IsSubdomain("", ""))
	fmt.Printf("%-14s: %5t\n", "same", netutil.IsSubdomain("example.com", "example.com"))
	fmt.Printf("%-14s: %5t\n", "dot only", netutil.IsSubdomain(".example.com", "example.com"))
	fmt.Printf("%-14s: %5t\n", "backwards", netutil.IsSubdomain("example.com", "sub.example.com"))
	fmt.Printf("%-14s: %5t\n", "other domain", netutil.IsSubdomain("sub.example.com", "example.org"))
	fmt.Printf("%-14s: %5t\n", "similar 1", netutil.IsSubdomain("sub.myexample.com", "example.org"))
	fmt.Printf("%-14s: %5t\n", "similar 2", netutil.IsSubdomain("sub.example.com", "myexample.org"))

	// Output:
	// same domain   :  true
	// not immediate :  true
	// empty         : false
	// same          : false
	// dot only      : false
	// backwards     : false
	// other domain  : false
	// similar 1     : false
	// similar 2     : false
}

func ExampleIsImmediateSubdomain() {
	fmt.Printf("%-14s: %5t\n", "same domain", netutil.IsImmediateSubdomain("sub.example.com", "example.com"))

	fmt.Printf("%-14s: %5t\n", "empty", netutil.IsImmediateSubdomain("", ""))
	fmt.Printf("%-14s: %5t\n", "same", netutil.IsImmediateSubdomain("example.com", "example.com"))
	fmt.Printf("%-14s: %5t\n", "dot only", netutil.IsImmediateSubdomain(".example.com", "example.com"))
	fmt.Printf("%-14s: %5t\n", "backwards", netutil.IsImmediateSubdomain("example.com", "sub.example.com"))
	fmt.Printf("%-14s: %5t\n", "other domain", netutil.IsImmediateSubdomain("sub.example.com", "example.org"))
	fmt.Printf("%-14s: %5t\n", "not immediate", netutil.IsImmediateSubdomain("subsub.sub.example.com", "example.com"))
	fmt.Printf("%-14s: %5t\n", "similar 1", netutil.IsSubdomain("sub.myexample.com", "example.org"))
	fmt.Printf("%-14s: %5t\n", "similar 2", netutil.IsSubdomain("sub.example.com", "myexample.org"))

	// Output:
	// same domain   :  true
	// empty         : false
	// same          : false
	// dot only      : false
	// backwards     : false
	// other domain  : false
	// not immediate : false
	// similar 1     : false
	// similar 2     : false
}
