package system

import (
	"back-end/model/system"
	"testing"
)

func TestMenuService_getChildrenList(t *testing.T) {
	type args struct {
		menu    *system.MenuTree
		treeMap map[uint][]system.MenuTree
	}
	tests := []struct {
		name        string
		userService *MenuService
		args        args
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.userService.getChildrenList(tt.args.menu, tt.args.treeMap); (err != nil) != tt.wantErr {
				t.Errorf("MenuService.getChildrenList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
