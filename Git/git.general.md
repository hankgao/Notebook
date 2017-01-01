## git tag
[git tag介绍](http://xxw8393.blog.163.com/blog/static/37256834201010301043564/)
[Git Basics - Tagging](https://git-scm.com/book/en/v2/Git-Basics-Tagging)

## debug using liteIDE
[Debugging Go using LiteIDE on Ubuntu 14.04](https://www.leaseweb.com/labs/2014/06/debugging-golang-using-liteide-on-ubuntu-14-04/?utm_source=leaseweblabs.com&utm_medium=referral&utm_campaign=redirect)

### solve the "please check gdb is codesigned" issue
[refer to](http://stackoverflow.com/questions/18423124/please-check-gdb-is-codesigned-see-taskgated8-how-to-get-gdb-installed-w)

#### codesign gdb
This error occurs because OSX implements a pid access policy which requires a digital signature for binaries to access other processes pids. To enable gdb access to other processes, we must first code sign the binary. This signature depends on a particular certificate, which the user must create and register with the system.

To create a code signing certificate, open the Keychain Access application. Choose menu Keychain Access -> Certificate Assistant -> Create a Certificate…

Choose a name for the certificate (e.g., gdb-cert), set Identity Type to Self Signed Root, set Certificate Type to Code Signing and select the Let me override defaults. Click several times on Continue until you get to the Specify a Location For The Certificate screen, then set Keychain to System.

Double click on the certificate, open Trust section, and set Code Signing to Always Trust. Exit Keychain Access application.

Restart the taskgated service, and sign the binary.

$ sudo killall taskgated
$ codesign -fs gdb-cert "$(which gdb)"



## pull requests
reference: [Making a Pull Request](https://www.atlassian.com/git/tutorials/making-a-pull-request)

Pull requests are a feature that makes it easier for developers to collaborate using Bitbucket. They provide a user-friendly web interface for discussing proposed changes before integrating them into the official project.

In their simplest form, pull requests are a mechanism for a developer to notify team members that they have completed a feature. Once their feature branch is ready, the developer files a pull request via their Bitbucket account. This lets everybody involved know that they need to review the code and merge it into the master branch.

But, the pull request is more than just a notification—it’s a dedicated forum for discussing the proposed feature. If there are any problems with the changes, teammates can post feedback in the pull request and even tweak the feature by pushing follow-up commits. All of this activity is tracked directly inside of the pull request.

## Unnamed

List all files that have been changed in a commit
- `$ git diff-tree --no-commit-id --name-only -r bd61ad98 `
- `git show --pretty="" --name-only bd61ad98 `
    The --no-commit-id suppresses the commit ID output.
    The --pretty argument specifies an empty format string to avoid the cruft at the beginning.
    The --name-only argument shows only the file names that were affected (Thanks Hank).
    The -r argument is to recurse into sub-trees

List all fiels in a commit
- `git ls-tree --name-only -r <commit-ish>`

### list a specific person's commits 

This works for both git log and gitk - the 2 most common ways of viewing history. You don't need to use the whole name.

`git log --author="Jon"` will match a commit made by "Jonathan Smith"
`git log --author=Jon` and `git log --author=Smith`  would also work. The quotes are optional if you don't need any spaces.

Add --all if you intend to search all branches and not just the current commit's ancestors in your repo.

You can also easily match on multiple authors as regex is the underlying mechanism for this filter. So to list commits by Jonathan or Adam, you can do this:

`git log --author="\(Adam\)\|\(Jon\)"`
In order to exclude commits by a particular author or set of authors using regular expressions as noted in this question, you can use a negative lookahead in combination with the --perl-regexp switch:

`git log --author='^(?!Adam|Jon).*$' --perl-regexp`
Alternatively, you can exclude commits authored by Adam by using bash and piping:

`git log --format='%H %an' | 
  grep -v Adam | 
  cut -d ' ' -f1 | 
  xargs -n1 git log -1`
If you want to exclude commits commited (but not necessarily authored) by Adam, replace %an with %cn. More details about this are in my blog post here: http://dymitruk.com/blog/2012/07/18/filtering-by-author-name/

## Git submodule

[Git Tools - Submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules)

### add submodules to your project
Submodules allow you to keep a Git repository as a subdirectory of another Git repository. This lets you clone another repository into your project and keep your commits separate.

To add a new submodule you use the git submodule add command with the absolute or relative URL of the project you would like to start tracking. 

    $ git submodule add https://github.com/chaconinc/DbConnector
    Cloning into 'DbConnector'...
    remote: Counting objects: 11, done.
    remote: Compressing objects: 100% (10/10), done.
    remote: Total 11 (delta 0), reused 11 (delta 0)
    Unpacking objects: 100% (11/11), done.
    Checking connectivity... done.

By default, submodules will add the subproject into a directory named the same as the repository, in this case “DbConnector”. You can add a different path at the end of the command if you want it to go elsewhere.

.gitmodules file is the configuration file that stores the mapping between the project’s URL and the local subdirectory you’ve pulled it into:

    [submodule "DbConnector"]
      path = DbConnector
      url = https://github.com/chaconinc/DbConnector

If you have multiple submodules, you’ll have multiple entries in this file. It’s important to note that this file is version-controlled with your other files, like your .gitignore file. It’s pushed and pulled with the rest of your project. This is how other people who clone this project know where to get the submodule projects from.

Although DbConnector is a subdirectory in your working directory, Git sees it as a submodule and doesn’t track its contents when you’re not in that directory. Instead, Git sees it as a particular commit from that repository.

### Cloning a Project with Submodules
Here we’ll clone a project with a submodule in it. When you clone such a project, by default you get the directories that contain submodules, but none of the files within them yet:

You must run two commands: `git submodule init` to initialize your local configuration file, and `git submodule update` to fetch all the data from that project and check out the appropriate commit listed in your superproject:


There is another way to do this which is a little simpler, however. If you pass --recursive to the git clone command, it will automatically initialize and update each submodule in the repository.

### Working on a Project with Submodules
If you want to check for new work in a submodule, you can go into the directory and run git fetch and git merge the upstream branch to update the local code.

There is an easier way to do this as well, if you prefer to not manually fetch and merge in the subdirectory. If you run git submodule update --remote, Git will go into your submodules and fetch and update for you.

This command will by default assume that you want to update the checkout to the master branch of the submodule repository. You can, however, set this to something different if you want. For example, if you want to have the DbConnector submodule track that repository’s “stable” branch, you can set it in either your .gitmodules file (so everyone else also tracks it), or just in your local .git/config file. Let’s set it in the .gitmodules file:



## git configuration
- there are three configuraiton files 
  + one under each repository: .git/config
  > `git config -file user.name hankgao`
  + another one is located in your $HOME directory: gitconfig
  > `git config -global user.name hankgal`
  + The next config file is your system config file, gitconfig. This contains your system configuration settings used in running the git application.  You can add to your etc/gitconfig file with the --system flag like this: 
  `git config --system core.autocrlf true`

 Note that 
 > Git uses these in order when looking for parameters.  First the local repository .git/config, then the global repository .gitconfig, and then the system repository etc/gitconfig.

## 忽略文件

一般我们总会有些文件无需纳入 Git 的管理，也不希望它们总出现在未跟踪文件列表。通常都是些自动生成的文件，比如日志文件，或者编译过程中创建的临时文件等。 在这种情况下，我们可以创建一个名为 .gitignore 的文件，列出要忽略的文件模式。 来看一个实际的例子：

`$ cat .gitignore`
`*.[oa]`
`*~`
第一行告诉 Git 忽略所有以 .o 或 .a 结尾的文件。一般这类对象文件和存档文件都是编译过程中出现的。 第二行告诉 Git 忽略所有以波浪符（~）结尾的文件，许多文本编辑软件（比如 Emacs）都用这样的文件名保存副本。 此外，你可能还需要忽略 log，tmp 或者 pid 目录，以及自动生成的文档等等。 要养成一开始就设置好 .gitignore 文件的习惯，以免将来误提交这类无用的文件。

文件 .gitignore 的格式规范如下：

- 所有空行或者以 ＃ 开头的行都会被 Git 忽略。
- 可以使用标准的 glob 模式匹配。
- 匹配模式可以以（/）开头防止递归。
- 匹配模式可以以（/）结尾指定目录。
- 要忽略指定模式以外的文件或目录，可以在模式前加上惊叹号（!）取反。

所谓的 glob 模式是指 shell 所使用的简化了的正则表达式。 星号（*）匹配零个或多个任意字符；[abc] 匹配任何一个列在方括号中的字符（这个例子要么匹配一个 a，要么匹配一个 b，要么匹配一个 c）；问号（?）只匹配一个任意字符；如果在方括号中使用短划线分隔两个字符，表示所有在这两个字符范围内的都可以匹配（比如 [0-9] 表示匹配所有 0 到 9 的数字）。 使用两个星号（*) 表示匹配任意中间目录，比如a/**/z 可以匹配 a/z, a/b/z 或 a/b/c/z等。

我们再看一个 .gitignore 文件的例子：

    # no .a files
    *.a

    # but do track lib.a, even though you're ignoring .a files above
    !lib.a

    # only ignore the TODO file in the current directory, not subdir/TODO
    /TODO

    # ignore all files in the build/ directory
    build/

    # ignore doc/notes.txt, but not doc/server/arch.txt
    doc/*.txt

    # ignore all .pdf files in the doc/ directory
    doc/**/*.pdf

Tip

GitHub 有一个十分详细的针对数十种项目及语言的 .gitignore 文件列表，你可以在 https://github.com/github/gitignore 找到它.

## 移除文件
要从 Git 中移除某个文件，就必须要从已跟踪文件清单中移除（确切地说，是从暂存区域移除），然后提交。 可以用 git rm 命令完成此项工作，并连带从工作目录中删除指定的文件，这样以后就不会出现在未跟踪文件清单中了。

如果只是简单地从工作目录中手工删除文件，运行 git status 时就会在 “Changes not staged for commit” 部分（也就是 未暂存清单）看到：

    $ rm PROJECTS.md
    $ git status
    On branch master
    Your branch is up-to-date with 'origin/master'.
    Changes not staged for commit:
      (use "git add/rm &lt;file&gt;..." to update what will be committed)
      (use "git checkout -- &lt;file&gt;..." to discard changes in working directory)    

            deleted:    PROJECTS.md    

    no changes added to commit (use "git add" and/or "git commit -a")

