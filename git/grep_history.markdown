
From [jeet](http://stackoverflow.com/users/268330/jeet) on [stackoverflow](http://stackoverflow.com/questions/2928584/how-to-grep-search-committed-code-in-the-git-history)

To search for commit content  

    git grep <regexp> $(git rev-list --all)

Search working tree  

    git grep <regexp>
    git grep -e <regexp1> [--or] -e <regexp2>
    git grep -e <regexp1> --and -e <regexp2>

Search working tree for files that have lines of text matching regular expression

    git grep -l --all-match -e <regexp1> -e <regexp2>

Search all revisions for text matching regular expression regexp:

    git grep <regexp> $(git rev-list --all)

Search all revisions between rev1 and rev2 for text matching regular expression

    git grep <regexp> $(git rev-list <rev1>..<rev2>)

