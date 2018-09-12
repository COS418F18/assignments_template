# COS418 Assignment 1 (Part 1): Intro to Go

<h2>Introduction</h2>
<p>
  In this assignment you will solve two short problems as a way to familiarize
  yourself with the Go programming language. We expect you to already have a
  basic knowledge of the language. If you're starting from nothing, we highly
  recommend going through the <a href="http://tour.golang.org/list">Golang tour</a>
  before you begin this assignment. Get started by
  <a href="https://golang.org/doc/install">installing Go</a> on your machine.
</p>
  <!-- In this class we will use the latest stable version 1.9.  -->
<h2>Software</h2>
<p>
  You will find the code the same directory. The two problems that you need to solve are in <tt>q1.go</tt>
  and <tt>q2.go</tt>. You should only add code to places that say <tt>TODO: implement me</tt>. 
  Do not change any of the function signatures as our testing framework uses them.
</p>

<p>
  <b>Q1 - Top K words:</b> The task is to find the <tt>K</tt> most common words in a
  given document. To exclude common words such as "a" and "the", the user of your program
  should be able to specify the minimum character threshold for a word. Word matching is
  case insensitive and punctuations should be removed. You can find more details on what
  qualifies as a word in the comments in the code.
</p>

<p>
  <b>Q2 - Parallel sum:</b> The task is to implement a function that sums a list of
  numbers in a file in parallel. For this problem you are required to use goroutines (the
  <tt>go</tt> keyword) and channels to pass messages across the goroutines. While it is
  possible to just sum all the numbers sequentially, the point of this problem is to
  familiarize yourself with the synchronization mechanisms in Go.
</p>

<h3>Testing</h3>

<p>
  Our grading uses the tests in <tt>q1_test.go</tt> and <tt>q2_test.go</tt> provided to you.
  To test the correctness of your code, run the following:
</p>
<pre>
  $ cd assignment1-1
  $ go test
</pre>
<p>
  If all tests pass, you should see the following output:
</p>
<pre>
  $ go test
  PASS
  ok      /path/to/assignment1-1   0.009s
</pre>




### Submitting Assignment
<p> Now you need to submit your assignment. Commit your change and push it to the remote repository by doing the following: 

```bash
$ git commit -am "[you fill me in]"
$ git tag -a -m "i finished assignment 1-1" a11-handin
$ git push origin master
$ git push origin a11-handin
```

