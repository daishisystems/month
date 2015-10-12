<a href="http://insidethecpu.com">![Image of insidethecpu](https://dl.dropboxusercontent.com/u/26042707/Daishi%20Systems%20Icon%20with%20Text%20%28really%20tiny%20with%20photo%29.png)</a>
# Go Month Package

[![Join the chat at https://gitter.im/daishisystems/month](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/daishisystems/month?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://travis-ci.org/daishisystems/month.svg)](https://travis-ci.org/daishisystems/month)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/daishisystems/month)

*Package month provides functionality pertaining to months of the year. The package is designed with developer productivity in mind, encapsulating some of the more verbose features of underlying Go packages, and exposing them intuitively. Examples of this include the ability to return the last day of any given month.*

![Month Icon](https://dl.dropboxusercontent.com/u/26042707/gomonth%28medium%29.jpg)
## Installation
go get github.com/daishisystems/month
## Sample Code
```go
	// The last numeric day of January is 31
	m := month.January
	fmt.Printf("The last numeric day of %s is %d\n", m, m.LastDay(2015))

	// The last numeric day of February is 28
	m = month.February
	fmt.Printf("The last numeric day of %s is %d\n", m, month.February.LastDay(2015))

	// The last numeric day of February is 29
	m = month.February
	fmt.Printf("The last numeric day of %s is %d\n", m, m.LastDay(2008))

	// The last numeric day of July is 31
	m = month.July
	fmt.Printf("The last numeric day of %s is %d\n", m, m.LastDay(2015))
```
## Contact the Developer
Please reach out and contact me for questions, suggestions, or to just talk tech in general.


<a href="http://insidethecpu.com/feed/">![RSS](https://dl.dropboxusercontent.com/u/26042707/rss.png)</a><a href="https://twitter.com/daishisystems">![Twitter](https://dl.dropboxusercontent.com/u/26042707/twitter.png)</a><a href="https://www.linkedin.com/in/daishisystems">![LinkedIn](https://dl.dropboxusercontent.com/u/26042707/linkedin.png)</a><a href="https://plus.google.com/102806071104797194504/posts">![Google+](https://dl.dropboxusercontent.com/u/26042707/g.png)</a><a href="https://www.youtube.com/user/daishisystems">![YouTube](https://dl.dropboxusercontent.com/u/26042707/youtube.png)</a>