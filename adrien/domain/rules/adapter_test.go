package rules

import (
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	content := `
        FIRST: [a-zA-Z]+ '@' [0-9]{1};
        SECOND: '-->' [0-9]{2,3};
		THIRD: [0-9]{2,3};
		FOURTH: 'yay';
        ----
        [a-zA-Z]+: a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z -> 1,;
        [0-9]{2,3}: 0,1,2,3,4,5,6,7,8,9 -> 2,3;
		[0-9]{1}: 0,1,2,3,4,5,6,7,8,9 -> 1;
    `

	adapter := NewAdapter()
	rules, err := adapter.ToRules(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := rules.All()
	if len(list) != 4 {
		t.Errorf("%d rules were expected, %d returned", 4, len(list))
		return
	}
}
