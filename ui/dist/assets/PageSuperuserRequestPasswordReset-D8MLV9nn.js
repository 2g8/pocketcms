import{S as M,i as T,s as z,F as A,d as E,t as w,a as y,m as H,c as L,h as g,C as B,D,l as k,n as d,E as G,G as I,v,u as m,w as p,p as C,H as F,I as N,A as h,f as j,k as P,o as R,q as J,z as S}from"./index-DP4U8Zk5.js";function K(u){let e,s,n,l,t,r,c,_,i,a,b,f;return l=new j({props:{class:"form-field required",name:"email",$$slots:{default:[Q,({uniqueId:o})=>({5:o}),({uniqueId:o})=>o?32:0]},$$scope:{ctx:u}}}),{c(){e=m("form"),s=m("div"),s.innerHTML='<h4 class="m-b-xs">Forgotten superuser password</h4> <p>Enter the email associated with your account and we’ll send you a recovery link:</p>',n=v(),L(l.$$.fragment),t=v(),r=m("button"),c=m("i"),_=v(),i=m("span"),i.textContent="Send recovery link",p(s,"class","content txt-center m-b-sm"),p(c,"class","ri-mail-send-line"),p(i,"class","txt"),p(r,"type","submit"),p(r,"class","btn btn-lg btn-block"),r.disabled=u[1],P(r,"btn-loading",u[1]),p(e,"class","m-b-base")},m(o,$){k(o,e,$),d(e,s),d(e,n),H(l,e,null),d(e,t),d(e,r),d(r,c),d(r,_),d(r,i),a=!0,b||(f=R(e,"submit",J(u[3])),b=!0)},p(o,$){const q={};$&97&&(q.$$scope={dirty:$,ctx:o}),l.$set(q),(!a||$&2)&&(r.disabled=o[1]),(!a||$&2)&&P(r,"btn-loading",o[1])},i(o){a||(y(l.$$.fragment,o),a=!0)},o(o){w(l.$$.fragment,o),a=!1},d(o){o&&g(e),E(l),b=!1,f()}}}function O(u){let e,s,n,l,t,r,c,_,i;return{c(){e=m("div"),s=m("div"),s.innerHTML='<i class="ri-checkbox-circle-line"></i>',n=v(),l=m("div"),t=m("p"),r=h("Check "),c=m("strong"),_=h(u[0]),i=h(" for the recovery link."),p(s,"class","icon"),p(c,"class","txt-nowrap"),p(l,"class","content"),p(e,"class","alert alert-success")},m(a,b){k(a,e,b),d(e,s),d(e,n),d(e,l),d(l,t),d(t,r),d(t,c),d(c,_),d(t,i)},p(a,b){b&1&&N(_,a[0])},i:F,o:F,d(a){a&&g(e)}}}function Q(u){let e,s,n,l,t,r,c,_;return{c(){e=m("label"),s=h("Email"),l=v(),t=m("input"),p(e,"for",n=u[5]),p(t,"type","email"),p(t,"id",r=u[5]),t.required=!0,t.autofocus=!0},m(i,a){k(i,e,a),d(e,s),k(i,l,a),k(i,t,a),S(t,u[0]),t.focus(),c||(_=R(t,"input",u[4]),c=!0)},p(i,a){a&32&&n!==(n=i[5])&&p(e,"for",n),a&32&&r!==(r=i[5])&&p(t,"id",r),a&1&&t.value!==i[0]&&S(t,i[0])},d(i){i&&(g(e),g(l),g(t)),c=!1,_()}}}function U(u){let e,s,n,l,t,r,c,_;const i=[O,K],a=[];function b(f,o){return f[2]?0:1}return e=b(u),s=a[e]=i[e](u),{c(){s.c(),n=v(),l=m("div"),t=m("a"),t.textContent="Back to login",p(t,"href","/login"),p(t,"class","link-hint"),p(l,"class","content txt-center")},m(f,o){a[e].m(f,o),k(f,n,o),k(f,l,o),d(l,t),r=!0,c||(_=G(I.call(null,t)),c=!0)},p(f,o){let $=e;e=b(f),e===$?a[e].p(f,o):(B(),w(a[$],1,1,()=>{a[$]=null}),D(),s=a[e],s?s.p(f,o):(s=a[e]=i[e](f),s.c()),y(s,1),s.m(n.parentNode,n))},i(f){r||(y(s),r=!0)},o(f){w(s),r=!1},d(f){f&&(g(n),g(l)),a[e].d(f),c=!1,_()}}}function V(u){let e,s;return e=new A({props:{$$slots:{default:[U]},$$scope:{ctx:u}}}),{c(){L(e.$$.fragment)},m(n,l){H(e,n,l),s=!0},p(n,[l]){const t={};l&71&&(t.$$scope={dirty:l,ctx:n}),e.$set(t)},i(n){s||(y(e.$$.fragment,n),s=!0)},o(n){w(e.$$.fragment,n),s=!1},d(n){E(e,n)}}}function W(u,e,s){let n="",l=!1,t=!1;async function r(){if(!l){s(1,l=!0);try{await C.collection("_superusers").requestPasswordReset(n),s(2,t=!0)}catch(_){C.error(_)}s(1,l=!1)}}function c(){n=this.value,s(0,n)}return[n,l,t,r,c]}class Y extends M{constructor(e){super(),T(this,e,W,V,z,{})}}export{Y as default};
