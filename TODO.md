# Short Term

* Make privilege-dropping more robust (Don't drop if nothing specified)
* Add pre-update hooks
* Add post-update hooks
* Add service wrapper scripts for Debian, Red Hat
* Add LICENSE file; remember to attribute Go for privilege-lookup code

# Long Term

* Consider using libgit2 to manipulate repositories
* Improve repository update mechanism
* Support separate repositories + working copies (i.e. a bare repository + staging area)

# Cool Ideas

* Implement git commit hook that produces a valid push message
	* Should probably read .git/config values for push destination
