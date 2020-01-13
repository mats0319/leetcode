// Go_mod_update can update packages in mod to it's master version.
// Firstly, this tool is write for match tools repo: utils' update.
// I'm lazy to modify go.mod after each time I update utils repo ￣へ￣
// so, this tool comes.
package main

func main() {
	// get require slice. (only keep version is ok)
	// get require: only one line or several lines
	// trim space at the start and the end
	// split with space, slice[0] is package name, slice[1] is version
	//
	// update version.
	// backup version string
	// change it to 'master' and overwrite
	// run 'go lint ?' /* if not find, use 'go mod tidy' directly */
	//
	// get require slice once more. (need to keep both package name and new version)
	// compare the versions, if not equal, print it.
	//
	// if there is a package update, run 'go mod tidy', if it has ran before, skip this step.
}
