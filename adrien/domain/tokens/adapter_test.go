package tokens

import (
	"fmt"
	"testing"

	"github.com/deepvalue-network/software/adrien/domain/rules"
)

func TestAdapter_Success(t *testing.T) {
	rulesContent := `
        NAME: [a-zA-Z]+;
        DOMAIN: [a-zA-Z\_]{2,};
        EXTENSION: [a-z]{3,};
        COMMERCIAL_A: '@';
		DOT: '.';
        QUOTATION_MARK: '"';
        COMMA: ',';
        PIPE: '|';
        ----
        [a-zA-Z]+: a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z -> 1,;
        [a-zA-Z\_]{2,}: a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,_ -> 2,;
        [a-z]{3,}: a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z -> 3,;
    `

	rulesAdapter := rules.NewAdapter()
	rulesIns, err := rulesAdapter.ToRules(rulesContent)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	adapter, err := NewAdapterBuilder().Create().WithRules(rulesIns).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil")
		return
	}

	tokensContent := `

        sections
            : section+
            ;

        section
            : anyAmountExceptTenOrFifteen PIPE
            ;

        anyAmountExceptTenOrFifteen
            : emailList
            ---
            : emailList -> email{10} <-
            | emailList -> email{15} <-
            ;

        atMostTenEmails
            : emailList -> email{0,10} <-
            ;

        atLeastThreeEmails
            : emailList -> email{3,} <-
            ;

        threeToNineEmails
            : emailList -> email{3,9} <-
            ;

        fiveEmailsAndForCommas
            : emailList -> email{5} COMMA{4} <-
            ;

        emailList
            : email emailWithCommaPrefix*
            ;

        emailWithCommaPrefix
            : COMMA email
            ;

        email
            : NAME COMMERCIAL_A DOMAIN DOT EXTENSION
            | QUOTATION_MARK email QUOTATION_MARK
            ;
    `

	tokens, err := adapter.ToTokens(tokensContent)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%v\n", tokens)

}
