// Code generated by "stringer -type=functionCode"; DO NOT EDIT.

package protocol

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[fcNil-0]
	_ = x[fcDDL-1]
	_ = x[fcInsert-2]
	_ = x[fcUpdate-3]
	_ = x[fcDelete-4]
	_ = x[fcSelect-5]
	_ = x[fcSelectForUpdate-6]
	_ = x[fcExplain-7]
	_ = x[fcDBProcedureCall-8]
	_ = x[fcDBProcedureCallWithResult-9]
	_ = x[fcFetch-10]
	_ = x[fcCommit-11]
	_ = x[fcRollback-12]
	_ = x[fcSavepoint-13]
	_ = x[fcConnect-14]
	_ = x[fcWriteLob-15]
	_ = x[fcReadLob-16]
	_ = x[fcPing-17]
	_ = x[fcDisconnect-18]
	_ = x[fcCloseCursor-19]
	_ = x[fcFindLob-20]
	_ = x[fcAbapStream-21]
	_ = x[fcXAStart-22]
	_ = x[fcXAJoin-23]
}

const _functionCode_name = "fcNilfcDDLfcInsertfcUpdatefcDeletefcSelectfcSelectForUpdatefcExplainfcDBProcedureCallfcDBProcedureCallWithResultfcFetchfcCommitfcRollbackfcSavepointfcConnectfcWriteLobfcReadLobfcPingfcDisconnectfcCloseCursorfcFindLobfcAbapStreamfcXAStartfcXAJoin"

var _functionCode_index = [...]uint8{0, 5, 10, 18, 26, 34, 42, 59, 68, 85, 112, 119, 127, 137, 148, 157, 167, 176, 182, 194, 207, 216, 228, 237, 245}

func (i functionCode) String() string {
	if i < 0 || i >= functionCode(len(_functionCode_index)-1) {
		return "functionCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _functionCode_name[_functionCode_index[i]:_functionCode_index[i+1]]
}
