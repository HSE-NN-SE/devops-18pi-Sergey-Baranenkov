(this["webpackJsonpmy-app"]=this["webpackJsonpmy-app"]||[]).push([[10],{106:function(e,t,a){"use strict";a.r(t);var n=a(5),r=a.n(n),c=a(25),s=a(15),i=a(19),o=a(9),l=a(10),p=a(17),u=a(12),m=a(11),d=a(0),h=a.n(d),f=(a(58),a(88),a(26)),_=(a(57),a(3)),b=a(21),E=a(6),v=a(22),y=a(20),g=a(36),O=a(23),j=function(e){Object(u.a)(a,e);var t=Object(m.a)(a);function a(){var e;Object(o.a)(this,a);for(var n=arguments.length,r=new Array(n),c=0;c<n;c++)r[c]=arguments[c];return(e=t.call.apply(t,[this].concat(r))).state={error:null},e.handleError=y.c.bind(Object(p.a)(e)),e.handleClose=y.b.bind(Object(p.a)(e)),e}return Object(l.a)(a,[{key:"render",value:function(){var e=this;return h.a.createElement("div",{className:"page__container"},h.a.createElement("div",{className:"page__header background_pic__city background_pic__city_green"},h.a.createElement("h1",{style:{color:"white"}},"\u0421\u0432\u044f\u0437\u0438"),h.a.createElement("p",{style:{color:"white"}},"\u0417\u0434\u0435\u0441\u044c \u043d\u0430\u0445\u043e\u0434\u0438\u0442\u0441\u044f \u0441\u043f\u0438\u0441\u043e\u043a \u0441\u0432\u044f\u0437\u0435\u0439"),h.a.createElement("div",{className:"friends_foreground_pic default_img"})),h.a.createElement("div",{className:"friends_classifier__container"},h.a.createElement(w,{to:"/\u0441\u0432\u044f\u0437\u0438/".concat(this.props.match.params.id,"/\u0434\u0440\u0443\u0437\u044c\u044f"),label:"\u0414\u0440\u0443\u0437\u044c\u044f"}),h.a.createElement(w,{to:"/\u0441\u0432\u044f\u0437\u0438/".concat(this.props.match.params.id,"/\u043f\u043e\u0434\u043f\u0438\u0441\u0447\u0438\u043a\u0438"),label:"\u041f\u043e\u0434\u043f\u0438\u0441\u0447\u0438\u043a\u0438"}),h.a.createElement(w,{to:"/\u0441\u0432\u044f\u0437\u0438/".concat(this.props.match.params.id,"/\u043f\u043e\u0434\u043f\u0438\u0441\u043a\u0438"),label:"\u041f\u043e\u0434\u043f\u0438\u0441\u043a\u0438"})),h.a.createElement(b.d,null,h.a.createElement(b.b,{path:"/\u0441\u0432\u044f\u0437\u0438/:id/\u0434\u0440\u0443\u0437\u044c\u044f",render:function(t){return h.a.createElement(k,Object.assign({type:3,handleError:e.handleError},t))}}),h.a.createElement(b.b,{path:"/\u0441\u0432\u044f\u0437\u0438/:id/\u043f\u043e\u0434\u043f\u0438\u0441\u0447\u0438\u043a\u0438",render:function(t){return h.a.createElement(k,Object.assign({type:2,handleError:e.handleError},t))}}),h.a.createElement(b.b,{path:"/\u0441\u0432\u044f\u0437\u0438/:id/\u043f\u043e\u0434\u043f\u0438\u0441\u043a\u0438",render:function(t){return h.a.createElement(k,Object.assign({type:1,handleError:e.handleError},t))}})),this.state.error&&h.a.createElement(y.a,{text:this.state.error,handleClose:this.handleClose}))}}]),a}(h.a.Component);function w(e){var t=e.label,a=e.to,n=e.activeOnlyWhenExact,r={color:"white"};return Object(b.h)({path:a,exact:n})&&(r.backgroundColor="orange"),h.a.createElement(E.b,{style:r,className:"default_link friends_classifier__item",to:a},t)}var k=function(e){Object(u.a)(a,e);var t=Object(m.a)(a);function a(){var e;Object(o.a)(this,a);for(var n=arguments.length,l=new Array(n),p=0;p<n;p++)l[p]=arguments[p];return(e=t.call.apply(t,[this].concat(l))).state={people:[],offset:0,isFetching:!1,Done:!1},e.limit=15,e.myId=+Object(O.a)("userId"),e.isPageMine=+e.props.match.params.id===e.myId,e.handleScrollThrottled=Object(g.a)((function(){Math.abs(window.scrollY+window.innerHeight-document.documentElement.scrollHeight)<10&&!e.state.Done&&!e.state.isFetching&&e.fetchData()}),1e3),e.fetchData=Object(i.a)(r.a.mark((function t(){var a,n,i,o;return r.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return e.setState({isFetching:!0}),t.next=3,Object(v.a)(_.b+_.a+"/relations/get_relations",{userId:e.props.match.params.id,mode:e.props.type,limit:e.limit,offset:e.state.offset});case 3:a=t.sent,n=Object(s.a)(a,2),i=n[0],o=n[1],null===i?e.setState((function(t){return{people:[].concat(Object(c.a)(t.people),Object(c.a)(o.People)),offset:t.offset+e.limit,Done:o.Done}})):e.props.handleError("\u043d\u0435\u0432\u043e\u0437\u043c\u043e\u0436\u043d\u043e \u043f\u043e\u043b\u0443\u0447\u0438\u0442\u044c \u0434\u0430\u043d\u043d\u044b\u0435 \u0441 \u0441\u0435\u0440\u0432\u0435\u0440\u0430"),e.setState({isFetching:!1});case 9:case"end":return t.stop()}}),t)}))),e.onAction=function(){var t=Object(i.a)(r.a.mark((function t(a){var n,c,i;return r.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return n=_.b+_.a+"/relations/update_relationship",t.next=3,Object(v.a)(n,{prevRelType:e.props.type,userId:a},"POST","text");case 3:c=t.sent,i=Object(s.a)(c,1),null===i[0]?e.setState((function(e){return{people:e.people.filter((function(e){return e.user_id!==a})),offset:e.offset-1}})):e.props.handleError("\u043d\u0435\u0432\u043e\u0437\u043c\u043e\u0436\u043d\u043e \u043e\u0431\u043d\u043e\u0432\u0438\u0442\u044c \u0441\u0432\u044f\u0437\u044c");case 7:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}(),e}return Object(l.a)(a,[{key:"componentDidUpdate",value:function(e,t,a){var n=this;e.type===this.props.type&&this.props.match.params.id===e.match.params.id||(this.isPageMine=+this.props.match.params.id===this.myId,this.setState({people:[],offset:0,Done:!1},(function(){n.fetchData()})))}},{key:"componentDidMount",value:function(){this.fetchData(),window.addEventListener("scroll",this.handleScrollThrottled,!0)}},{key:"componentWillUnmount",value:function(){window.removeEventListener("scroll",this.handleScrollThrottled)}},{key:"render",value:function(){var e=this;return h.a.createElement(h.a.Fragment,null,this.state.people.map((function(t){return h.a.createElement(N,Object.assign({key:t.user_id},t,{type:e.props.type,onAction:e.onAction,isPageMine:e.isPageMine}))})),this.state.isFetching&&h.a.createElement("span",null,"\u0417\u0430\u0433\u0440\u0443\u0437\u043a\u0430..."))}}]),a}(h.a.Component);function N(e){var t=Object(b.g)(),a=function(){e.onAction(e.user_id)};return h.a.createElement("div",{className:"rel__container"},h.a.createElement("img",{className:"default_img rel_user__avatar",alt:" ",src:_.b+_.a+"/profile_bgs".concat(Object(f.a)(e.user_id),"/profile_avatar.jpg")}),h.a.createElement("div",{className:"rel_user__vertical"},h.a.createElement("span",{className:"rel_user__fullname",style:{cursor:"pointer"},onClick:function(){t.push("/\u043f\u0440\u043e\u0444\u0438\u043b\u044c/"+e.user_id)}},"".concat(e.first_name," ").concat(e.last_name)),h.a.createElement(E.b,{className:"rel_user__send_mes",to:{pathname:"/\u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u044f/",user_id:e.user_id,first_name:e.first_name,last_name:e.last_name}},"\u041d\u0430\u043f\u0438\u0441\u0430\u0442\u044c \u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u0435")),h.a.createElement("div",{className:"rel__functional"},e.isPageMine&&(3===e.type||1===e.type?h.a.createElement("button",{className:"rel__button delete_friend__button",onClick:a},"\u274c"):h.a.createElement("button",{className:"rel__button add_friend__button",onClick:a},"+"))))}t.default=Object(b.i)(j)},88:function(e,t,a){}}]);
//# sourceMappingURL=10.bbb0f016.chunk.js.map