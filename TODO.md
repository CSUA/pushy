# Short Term

* Revert privilege-dropping code; rely on whatever daemon wrappers are available to handle it instead
	* Will make it more portable
* Add pre-update, post-update hooks
* Add service wrapper scripts for Debian, Red Hat
* Add LICENSE file; remember to attribute Go for privilege-lookup code

# Long Term

* Consider using libgit2 to manipulate repositories
	* PRO: easier Windows deployment?
	* CON: external library dependency, and is Windows really that important?
* Add configurable repository update strategies
	* Fetch + detached HEAD
	* Bare repository + working copy
	* Vanilla git pull

# Cool Ideas

* Implement git commit hook that produces a valid push message
	* Should probably read .git/config values for push destination
* Add support for a heartbeat ping
* Add support for stats monitoring?
