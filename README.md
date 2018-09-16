# Assignments for COS 418

### Environment Setup

Please follow <a href="setup.md">these instructions</a> for setting up for Go environment for assignments, as well as pointers to some necessary/useful tools.

### Coding Style

<p>All of the code you turn in for this course should have good style.
Make sure that your code has proper indentation, descriptive comments,
and a comment header at the beginning of each file, which includes
your name, userid, and a description of the file.</p>

<p>A portion of credit for each assignment is determined by code
quality tests, using the standard tools <tt>gofmt</tt> and <tt>go
vet</tt>.  You will receive full credit for this portion if all files
submitted conform to the style standards set by <tt>gofmt</tt> and the
report from <tt>go vet</tt> is clean (that is, produces no errors).
If your code does not pass the <tt>gofmt</tt> test, you should
reformat your code using the tool. You can also use the <a
href="https://github.com/qiniu/checkstyle">Go Checkstyle</a> tool for
advice to improve your code's style, if applicable.  Additionally,
though not part of the graded cheks, it would also be advisable to
produce code that complies with <a
href="https://github.com/golang/lint">Golint</a> where possible. </p>

<h3>How do I git?</h3>
This page has some useful tutorials on git: <a href="https://www.atlassian.com/git/tutorial">Git Tutorials</a></br>
In particular, tutorials 1 and 5 will be most useful for this course.</p>

<p>The basic git workflow in the shell (assuming you already have a repo set up):</br>
<ul>
<li>git pull</li>
<li>do some work</li>
<li>git status (shows what has changed)</li>
<li>git add <i>all files you want to commit</i></li>
<li>git commit -m "brief message on your update"</li>
<li>git push</li>
</ul>
</p>

<p>Finally, <a href="https://confluence.atlassian.com/display/BITBUCKET/Bitbucket+101">Bitbucket 101</a> is another good resource.</p>


<p> All programming assignments, require Git for submission. <p> We are using Github for distributing and collecting your assignments. At the time of seeing this, you should have already joined the [COS418F18](https://github.com/orgs/COS418F18) organization on Github and forked your private repository. You will need to develop in a *nix environment, i.e., Linux or OS X. Your Github page should have a link. Normally, you only need to clone the repository once, and you will have everything you need for all the assignments in this class.

```bash
$ git clone https://github.com/COS418F18/assignments-myusername.git 418
$ cd 418
$ ls
assignment1-1  assignment1-2  assignment1-3  assignment2  assignment3  assignment4  assignment5  README.md  setup.md
$ 
```

Now, you have everything you need for doing all assignments, i.e., instructions and starter code. Git allows you to keep track of the changes you make to the code. For example, if you want to checkpoint your progress, you can <emph>commit</emph> your changes by running:

```bash
$ git commit -am 'partial solution to assignment 1-1'
$ 
```

You should do this early and often!  You can _push_ your changes to Github after you commit with:

```bash
$ git push origin master
$ 
```

Please let us know that you've gotten this far in the assignment, by pushing a tag to Github.

```bash
$ git tag -a -m "i got git and cloned the assignments" gotgit
$ git push origin gotgit
$
```

As you complete parts of the assignments (and begin future assignments) we'll ask you push tags. You should also be committing and pushing your progress regularly.

### Steping into Assignment 1-1

Now it's time to go to assignment 1-1 folder to begin your adventure!