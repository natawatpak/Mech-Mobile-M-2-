"use strict";(self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[]).push([[295],{5661:function(e,t,n){n.d(t,{_:function(){return F}});var o=n(3396),a=n(8434),l=n(320),r=n(9888);const i=(0,l.a)({name:"VCardActions",setup(e,t){let{slots:n}=t;return(0,a.AF)({VBtn:{variant:"text"}}),(0,r.L)((()=>{var e;return(0,o.Wm)("div",{class:"v-card-actions"},[null==(e=n.default)?void 0:e.call(n)])})),{}}});var s=n(652),c=n(1114);const u=(0,c.J)("v-card-subtitle");var d=n(6572),v=n(836),f=n(4960),p=n(9694);const m=(0,l.a)({name:"VCardItem",props:{appendAvatar:String,appendIcon:f.lE,prependAvatar:String,prependIcon:f.lE,subtitle:String,title:String,...(0,p.f)()},setup(e,t){let{slots:n}=t;return(0,r.L)((()=>{var t,a,l,r,i;const c=!!(e.prependAvatar||e.prependIcon||n.prepend),f=!!(e.appendAvatar||e.appendIcon||n.append),p=!(!e.title&&!n.title),m=!(!e.subtitle&&!n.subtitle);return(0,o.Wm)("div",{class:"v-card-item"},[c&&(0,o.Wm)(v.z,{key:"prepend",defaults:{VAvatar:{density:e.density,icon:e.prependIcon,image:e.prependAvatar},VIcon:{density:e.density,icon:e.prependIcon}}},{default:()=>[(0,o.Wm)("div",{class:"v-card-item__prepend"},[(null==(t=n.prepend)?void 0:t.call(n))??(0,o.Wm)(s.V,null,null)])]}),(0,o.Wm)("div",{class:"v-card-item__content"},[p&&(0,o.Wm)(d.E,{key:"title"},{default:()=>[(null==(a=n.title)?void 0:a.call(n))??e.title]}),m&&(0,o.Wm)(u,{key:"subtitle"},{default:()=>[(null==(l=n.subtitle)?void 0:l.call(n))??e.subtitle]}),null==(r=n.default)?void 0:r.call(n)]),f&&(0,o.Wm)(v.z,{key:"append",defaults:{VAvatar:{density:e.density,icon:e.appendIcon,image:e.appendAvatar},VIcon:{density:e.density,icon:e.appendIcon}}},{default:()=>[(0,o.Wm)("div",{class:"v-card-item__append"},[(null==(i=n.append)?void 0:i.call(n))??(0,o.Wm)(s.V,null,null)])]})])})),{}}});var y=n(1888),g=n(8694),h=n(3824),b=n(5221),x=n(1710),w=n(2718),E=n(4544),k=n(2465),O=n(5180),S=n(489),C=n(4231),A=n(6183),P=n(1138),B=n(7041);const F=(0,l.a)({name:"VCard",directives:{Ripple:h.H},props:{appendAvatar:String,appendIcon:f.lE,disabled:Boolean,flat:Boolean,hover:Boolean,image:String,link:{type:Boolean,default:void 0},prependAvatar:String,prependIcon:f.lE,ripple:Boolean,subtitle:String,text:String,title:String,...(0,B.x$)(),...(0,w.m)(),...(0,p.f)(),...(0,E.x)(),...(0,k.c)(),...(0,x.fF)(),...(0,O.y)(),...(0,S.F)(),...(0,C.I)(),...(0,A.GN)(),...(0,P.Q)(),...(0,b.bk)({variant:"elevated"})},setup(e,t){let{attrs:n,slots:a}=t;const{themeClasses:l}=(0,B.ER)(e),{borderClasses:s}=(0,w.P)(e),{colorClasses:c,colorStyles:u,variantClasses:d}=(0,b.c1)(e),{densityClasses:f}=(0,p.t)(e),{dimensionStyles:h}=(0,E.$)(e),{elevationClasses:P}=(0,k.Y)(e),{loaderClasses:F}=(0,x.U2)(e),{locationStyles:L}=(0,O.T)(e),{positionClasses:W}=(0,S.K)(e),{roundedClasses:R}=(0,C.b)(e),H=(0,A.nB)(e,n),_=(0,o.Fl)((()=>!1!==e.link&&H.isLink.value)),j=(0,o.Fl)((()=>!e.disabled&&!1!==e.link&&(e.link||H.isClickable.value)));return(0,r.L)((()=>{var t,n,r;const p=_.value?"a":e.tag,w=!(!a.title&&!e.title),E=!(!a.subtitle&&!e.subtitle),k=w||E,O=!!(a.append||e.appendAvatar||e.appendIcon),S=!!(a.prepend||e.prependAvatar||e.prependIcon),C=!(!a.image&&!e.image),A=k||S||O,B=!(!a.text&&!e.text);return(0,o.wy)((0,o.Wm)(p,{class:["v-card",{"v-card--disabled":e.disabled,"v-card--flat":e.flat,"v-card--hover":e.hover&&!(e.disabled||e.flat),"v-card--link":j.value},l.value,s.value,c.value,f.value,P.value,F.value,W.value,R.value,d.value],style:[u.value,h.value,L.value],href:H.href.value,onClick:j.value&&H.navigate,tabindex:e.disabled?-1:void 0},{default:()=>[C&&(0,o.Wm)(v.z,{key:"image",defaults:{VImg:{cover:!0,src:e.image}}},{default:()=>[(0,o.Wm)("div",{class:"v-card__image"},[(null==(t=a.image)?void 0:t.call(a))??(0,o.Wm)(g.f,null,null)])]}),(0,o.Wm)(x.rD,{name:"v-card",active:!!e.loading,color:"boolean"===typeof e.loading?void 0:e.loading},{default:a.loader}),A&&(0,o.Wm)(m,{key:"item",prependAvatar:e.prependAvatar,prependIcon:e.prependIcon,title:e.title,subtitle:e.subtitle,appendAvatar:e.appendAvatar,appendIcon:e.appendIcon},{default:a.item,prepend:a.prepend,title:a.title,subtitle:a.subtitle,append:a.append}),B&&(0,o.Wm)(y.Z,{key:"text"},{default:()=>[(null==(n=a.text)?void 0:n.call(a))??e.text]}),null==(r=a.default)?void 0:r.call(a),a.actions&&(0,o.Wm)(i,null,{default:a.actions}),(0,b.Ux)(j.value,"v-card")]}),[[(0,o.Q2)("ripple"),j.value]])})),{}}})},1888:function(e,t,n){n.d(t,{Z:function(){return a}});var o=n(1114);const a=(0,o.J)("v-card-text")},6572:function(e,t,n){n.d(t,{E:function(){return a}});var o=n(1114);const a=(0,o.J)("v-card-title")},3849:function(e,t,n){n.d(t,{B:function(){return Pe}});var o=n(3396),a=n(9242),l=n(320),r=n(3122),i=n(8587);const s=(0,l.a)({name:"VDialogTransition",props:{target:Object},setup(e,t){let{slots:n}=t;const l={onBeforeEnter(e){e.style.pointerEvents="none",e.style.visibility="hidden"},async onEnter(t,n){var o;await new Promise((e=>requestAnimationFrame(e))),await new Promise((e=>requestAnimationFrame(e))),t.style.visibility="";const{x:a,y:l,sx:s,sy:d,speed:v}=u(e.target,t),f=(0,r.j)(t,[{transform:`translate(${a}px, ${l}px) scale(${s}, ${d})`,opacity:0},{transform:""}],{duration:225*v,easing:i.uX});null==(o=c(t))||o.forEach((e=>{(0,r.j)(e,[{opacity:0},{opacity:0,offset:.33},{opacity:1}],{duration:450*v,easing:i.Ly})})),f.finished.then((()=>n()))},onAfterEnter(e){e.style.removeProperty("pointer-events")},onBeforeLeave(e){e.style.pointerEvents="none"},async onLeave(t,n){var o;await new Promise((e=>requestAnimationFrame(e)));const{x:a,y:l,sx:s,sy:d,speed:v}=u(e.target,t),f=(0,r.j)(t,[{transform:""},{transform:`translate(${a}px, ${l}px) scale(${s}, ${d})`,opacity:0}],{duration:125*v,easing:i.x0});f.finished.then((()=>n())),null==(o=c(t))||o.forEach((e=>{(0,r.j)(e,[{},{opacity:0,offset:.2},{opacity:0}],{duration:250*v,easing:i.Ly})}))},onAfterLeave(e){e.style.removeProperty("pointer-events")}};return()=>e.target?(0,o.Wm)(a.uT,(0,o.dG)({name:"dialog-transition"},l,{css:!1}),n):(0,o.Wm)(a.uT,{name:"dialog-transition"},n)}});function c(e){var t;const n=null==(t=e.querySelector(":scope > .v-card, :scope > .v-sheet, :scope > .v-list"))?void 0:t.children;return n&&[...n]}function u(e,t){const n=e.getBoundingClientRect(),o=(0,r.G)(t),[a,l]=getComputedStyle(t).transformOrigin.split(" ").map((e=>parseFloat(e))),[i,s]=getComputedStyle(t).getPropertyValue("--v-overlay-anchor-origin").split(" ");let c=n.left+n.width/2;"left"===i||"left"===s?c-=n.width/2:"right"!==i&&"right"!==s||(c+=n.width/2);let u=n.top+n.height/2;"top"===i||"top"===s?u-=n.height/2:"bottom"!==i&&"bottom"!==s||(u+=n.height/2);const d=n.width/o.width,v=n.height/o.height,f=Math.max(1,d,v),p=d/f,m=v/f,y=o.width*o.height/(window.innerWidth*window.innerHeight),g=y>.12?Math.min(1.5,10*(y-.12)+1):1;return{x:c-(a+o.left),y:u-(l+o.top),sx:p,sy:m,speed:g}}var d=n(836),v=n(3766),f=n(2385);const p=(0,v.U)({closeDelay:[Number,String],openDelay:[Number,String]},"delay");function m(e,t){const n={},o=o=>()=>{if(!f.BR)return Promise.resolve(!0);const a="openDelay"===o;return n.closeDelay&&window.clearTimeout(n.closeDelay),delete n.closeDelay,n.openDelay&&window.clearTimeout(n.openDelay),delete n.openDelay,new Promise((l=>{const r=parseInt(e[o]??0,10);n[o]=window.setTimeout((()=>{null==t||t(a),l(a)}),r)}))};return{runCloseDelay:o("closeDelay"),runOpenDelay:o("openDelay")}}const y=Symbol.for("vuetify:v-menu");var g=n(131),h=n(7514),b=n(4870);const x=(0,v.U)({activator:[String,Object],activatorProps:{type:Object,default:()=>({})},openOnClick:{type:Boolean,default:void 0},openOnHover:Boolean,openOnFocus:{type:Boolean,default:void 0},closeOnContentClick:Boolean,...p()},"v-overlay-activator");function w(e,t){let{isActive:n,isTop:a}=t;const l=(0,b.iH)();let r=!1,i=!1,s=!0;const c=(0,o.Fl)((()=>e.openOnFocus||null==e.openOnFocus&&e.openOnHover)),u=(0,o.Fl)((()=>e.openOnClick||null==e.openOnClick&&!e.openOnHover&&!c.value)),{runOpenDelay:d,runCloseDelay:v}=m(e,(t=>{t!==(e.openOnHover&&r||c.value&&i)||e.openOnHover&&n.value&&!a.value||(n.value!==t&&(s=!0),n.value=t)})),p={click:e=>{e.stopPropagation(),l.value=e.currentTarget||e.target,n.value=!n.value},mouseenter:e=>{r=!0,l.value=e.currentTarget||e.target,d()},mouseleave:e=>{r=!1,v()},focus:e=>{f.Z1&&!e.target.matches(":focus-visible")||(i=!0,e.stopPropagation(),l.value=e.currentTarget||e.target,d())},blur:e=>{i=!1,e.stopPropagation(),v()}},x=(0,o.Fl)((()=>{const t={};return u.value&&(t.click=p.click),e.openOnHover&&(t.mouseenter=p.mouseenter,t.mouseleave=p.mouseleave),c.value&&(t.focus=p.focus,t.blur=p.blur),t})),w=(0,o.Fl)((()=>{const t={};if(e.openOnHover&&(t.mouseenter=()=>{r=!0,d()},t.mouseleave=()=>{r=!1,v()}),e.closeOnContentClick){const e=(0,o.f3)(y,null);t.click=()=>{n.value=!1,null==e||e.closeParents()}}return t})),k=(0,o.Fl)((()=>{const t={};return e.openOnHover&&(t.mouseenter=()=>{s&&(r=!0,s=!1,d())},t.mouseleave=()=>{r=!1,v()}),t}));(0,o.YP)(a,(t=>{!t||(!e.openOnHover||r||c.value&&i)&&(!c.value||i||e.openOnHover&&r)||(n.value=!1)}));const O=(0,b.iH)();(0,o.m0)((()=>{O.value&&(0,o.Y3)((()=>{const e=O.value;l.value=(0,g.rU)(e)?e.$el:e}))}));const S=(0,h.FN)("useActivator");let C;return(0,o.YP)((()=>!!e.activator),(t=>{t&&f.BR?(C=(0,b.B)(),C.run((()=>{E(e,S,{activatorEl:l,activatorEvents:x})}))):C&&C.stop()}),{flush:"post",immediate:!0}),(0,b.EB)((()=>{var e;null==(e=C)||e.stop()})),{activatorEl:l,activatorRef:O,activatorEvents:x,contentEvents:w,scrimEvents:k}}function E(e,t,n){let{activatorEl:a,activatorEvents:l}=n;function r(){let t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:s(),n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:e.activatorProps;t&&(Object.entries(l.value).forEach((e=>{let[n,o]=e;t.addEventListener(n,o)})),Object.keys(n).forEach((e=>{null==n[e]?t.removeAttribute(e):t.setAttribute(e,n[e])})))}function i(){let t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:s(),n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:e.activatorProps;t&&(Object.entries(l.value).forEach((e=>{let[n,o]=e;t.removeEventListener(n,o)})),Object.keys(n).forEach((e=>{t.removeAttribute(e)})))}function s(){var n;let o,l=arguments.length>0&&void 0!==arguments[0]?arguments[0]:e.activator;if(l)if("parent"===l){var r,i;let e=null==t||null==(r=t.proxy)||null==(i=r.$el)?void 0:i.parentNode;while(e.hasAttribute("data-no-activator"))e=e.parentNode;o=e}else o="string"===typeof l?document.querySelector(l):"$el"in l?l.$el:l;return a.value=(null==(n=o)?void 0:n.nodeType)===Node.ELEMENT_NODE?o:null,a.value}(0,o.YP)((()=>e.activator),((e,t)=>{if(t&&e!==t){const e=s(t);e&&i(e)}e&&(0,o.Y3)((()=>r()))}),{immediate:!0}),(0,o.YP)((()=>e.activatorProps),(()=>{r()})),(0,b.EB)((()=>{i()}))}var k=n(4544);const O=(0,v.U)({eager:Boolean},"lazy");function S(e,t){const n=(0,b.iH)(!1),a=(0,o.Fl)((()=>n.value||e.eager||t.value));function l(){e.eager||(n.value=!1)}return(0,o.YP)(t,(()=>n.value=!0)),{isBooted:n,hasContent:a,onAfterLeave:l}}n(7658);function C(e){while(e){if("fixed"===window.getComputedStyle(e).position)return!0;e=e.offsetParent}return!1}var A=n(2879);function P(e){while(e){if(F(e))return e;e=e.parentElement}return document.scrollingElement}function B(e,t){const n=[];if(t&&e&&!t.contains(e))return n;while(e){if(F(e)&&n.push(e),e===t)break;e=e.parentElement}return n}function F(e){if(!e||e.nodeType!==Node.ELEMENT_NODE)return!1;const t=window.getComputedStyle(e);return"scroll"===t.overflowY||"auto"===t.overflowY&&e.scrollHeight>e.clientHeight}var L=n(6033),W=n(6309);function R(e,t){return{x:e.x+t.x,y:e.y+t.y}}function H(e,t){return{x:e.x-t.x,y:e.y-t.y}}function _(e,t){if("top"===e.side||"bottom"===e.side){const{side:n,align:o}=e,a="left"===o?0:"center"===o?t.width/2:"right"===o?t.width:o,l="top"===n?0:"bottom"===n?t.height:n;return R({x:a,y:l},t)}if("left"===e.side||"right"===e.side){const{side:n,align:o}=e,a="left"===n?0:"right"===n?t.width:n,l="top"===o?0:"center"===o?t.height/2:"bottom"===o?t.height:o;return R({x:a,y:l},t)}return R({x:t.width/2,y:t.height/2},t)}const j={static:V,connected:D},T=(0,v.U)({locationStrategy:{type:[String,Function],default:"static",validator:e=>"function"===typeof e||e in j},location:{type:String,default:"bottom"},origin:{type:String,default:"auto"},offset:[Number,String,Array]},"v-overlay-location-strategies");function I(e,t){const n=(0,b.iH)({}),a=(0,b.iH)();let l;function r(e){var t;null==(t=a.value)||t.call(a,e)}return(0,o.m0)((async()=>{var r;null==(r=l)||r.stop(),a.value=void 0,f.BR&&t.isActive.value&&e.locationStrategy&&(l=(0,b.B)(),await(0,o.Y3)(),l.run((()=>{var o,l;"function"===typeof e.locationStrategy?a.value=null==(o=e.locationStrategy(t,e,n))?void 0:o.updateLocation:a.value=null==(l=j[e.locationStrategy](t,e,n))?void 0:l.updateLocation})))})),f.BR&&window.addEventListener("resize",r,{passive:!0}),(0,b.EB)((()=>{var e;f.BR&&window.removeEventListener("resize",r),a.value=void 0,null==(e=l)||e.stop()})),{contentStyles:n,updateLocation:a}}function V(){}function $(e){const t=(0,r.G)(e);return t.x-=parseFloat(e.style.left||0),t.y-=parseFloat(e.style.top||0),t}function D(e,t,n){const a=C(e.activatorEl.value);a&&Object.assign(n.value,{position:"fixed"});const{preferredAnchor:l,preferredOrigin:r}=(0,g.S3)((()=>{const n=(0,A.wW)(t.location,e.isRtl.value),o="overlap"===t.origin?n:"auto"===t.origin?(0,A.tb)(n):(0,A.wW)(t.origin,e.isRtl.value);return n.side===o.side&&n.align===(0,A.aw)(o).align?{preferredAnchor:(0,A.Oe)(n),preferredOrigin:(0,A.Oe)(o)}:{preferredAnchor:n,preferredOrigin:o}})),[i,s,c,u]=["minWidth","minHeight","maxWidth","maxHeight"].map((e=>(0,o.Fl)((()=>{const n=parseFloat(t[e]);return isNaN(n)?1/0:n})))),d=(0,o.Fl)((()=>{if(Array.isArray(t.offset))return t.offset;if("string"===typeof t.offset){const e=t.offset.split(" ").map(parseFloat);return e.length<2&&e.push(0),e}return"number"===typeof t.offset?[t.offset,0]:[0,0]}));let v=!1;const f=new ResizeObserver((()=>{v&&p()}));function p(){if(v=!1,requestAnimationFrame((()=>{requestAnimationFrame((()=>v=!0))})),!e.activatorEl.value||!e.contentEl.value)return;const t=e.activatorEl.value.getBoundingClientRect(),o=$(e.contentEl.value),a=B(e.contentEl.value),f=12;a.length||(a.push(document.documentElement),e.contentEl.value.style.top&&e.contentEl.value.style.left||(o.x+=parseFloat(document.documentElement.style.getPropertyValue("--v-body-scroll-x")||0),o.y+=parseFloat(document.documentElement.style.getPropertyValue("--v-body-scroll-y")||0)));const p=a.reduce(((e,t)=>{const n=t.getBoundingClientRect(),o=new W.x({x:t===document.documentElement?0:n.x,y:t===document.documentElement?0:n.y,width:t.clientWidth,height:t.clientHeight});return e?new W.x({x:Math.max(e.left,o.left),y:Math.max(e.top,o.top),width:Math.min(e.right,o.right)-Math.max(e.left,o.left),height:Math.min(e.bottom,o.bottom)-Math.max(e.top,o.top)}):o}),void 0);p.x+=f,p.y+=f,p.width-=2*f,p.height-=2*f;let m={anchor:l.value,origin:r.value};function y(e){const n=new W.x(o),a=_(e.anchor,t),l=_(e.origin,n);let{x:r,y:i}=H(a,l);switch(e.anchor.side){case"top":i-=d.value[0];break;case"bottom":i+=d.value[0];break;case"left":r-=d.value[0];break;case"right":r+=d.value[0];break}switch(e.anchor.align){case"top":i-=d.value[1];break;case"bottom":i+=d.value[1];break;case"left":r-=d.value[1];break;case"right":r+=d.value[1];break}n.x+=r,n.y+=i,n.width=Math.min(n.width,c.value),n.height=Math.min(n.height,u.value);const s=(0,W.p)(n,p);return{overflows:s,x:r,y:i}}let h=0,b=0;const x={x:0,y:0},w={x:!1,y:!1};let E=-1;while(1){if(E++>10){(0,L.N6)("Infinite loop detected in connectedLocationStrategy");break}const{x:e,y:t,overflows:n}=y(m);h+=e,b+=t,o.x+=e,o.y+=t;{const e=(0,A.dd)(m.anchor),t=n.x.before||n.x.after,o=n.y.before||n.y.after;let a=!1;if(["x","y"].forEach((l=>{if("x"===l&&t&&!w.x||"y"===l&&o&&!w.y){const t={anchor:{...m.anchor},origin:{...m.origin}},o="x"===l?"y"===e?A.aw:A.tb:"y"===e?A.tb:A.aw;t.anchor=o(t.anchor),t.origin=o(t.origin);const{overflows:r}=y(t);(r[l].before<=n[l].before&&r[l].after<=n[l].after||r[l].before+r[l].after<(n[l].before+n[l].after)/2)&&(m=t,a=w[l]=!0)}})),a)continue}n.x.before&&(h+=n.x.before,o.x+=n.x.before),n.x.after&&(h-=n.x.after,o.x-=n.x.after),n.y.before&&(b+=n.y.before,o.y+=n.y.before),n.y.after&&(b-=n.y.after,o.y-=n.y.after);{const e=(0,W.p)(o,p);x.x=p.width-e.x.before-e.x.after,x.y=p.height-e.y.before-e.y.after,h+=e.x.before,o.x+=e.x.before,b+=e.y.before,o.y+=e.y.before}break}const k=(0,A.dd)(m.anchor);Object.assign(n.value,{"--v-overlay-anchor-origin":`${m.anchor.side} ${m.anchor.align}`,transformOrigin:`${m.origin.side} ${m.origin.align}`,top:(0,g.kb)(N(b)),left:(0,g.kb)(N(h)),minWidth:(0,g.kb)("y"===k?Math.min(i.value,t.width):i.value),maxWidth:(0,g.kb)(M((0,g.uZ)(x.x,i.value===1/0?0:i.value,c.value))),maxHeight:(0,g.kb)(M((0,g.uZ)(x.y,s.value===1/0?0:s.value,u.value)))})}return(0,o.YP)([e.activatorEl,e.contentEl],((e,t)=>{let[n,o]=e,[a,l]=t;a&&f.unobserve(a),n&&f.observe(n),l&&f.unobserve(l),o&&f.observe(o)}),{immediate:!0}),(0,b.EB)((()=>{f.disconnect()})),(0,o.YP)((()=>[l.value,r.value,t.offset,t.minWidth,t.minHeight,t.maxWidth,t.maxHeight]),(()=>p()),{immediate:!a}),a&&(0,o.Y3)((()=>p())),requestAnimationFrame((()=>{n.value.maxHeight&&p()})),{updateLocation:p}}function N(e){return Math.round(e*devicePixelRatio)/devicePixelRatio}function M(e){return Math.ceil(e*devicePixelRatio)/devicePixelRatio}let Y=!0;const z=[];function q(e){!Y||z.length?(z.push(e),U()):(Y=!1,e(),U())}let G=-1;function U(){cancelAnimationFrame(G),G=requestAnimationFrame((()=>{const e=z.shift();e&&e(),z.length?U():Y=!0}))}const J={none:null,close:K,block:X,reposition:ee},Z=(0,v.U)({scrollStrategy:{type:[String,Function],default:"block",validator:e=>"function"===typeof e||e in J}},"v-overlay-scroll-strategies");function Q(e,t){if(!f.BR)return;let n;(0,o.m0)((async()=>{var a;null==(a=n)||a.stop(),t.isActive.value&&e.scrollStrategy&&(n=(0,b.B)(),await(0,o.Y3)(),n.run((()=>{var n;"function"===typeof e.scrollStrategy?e.scrollStrategy(t,e):null==(n=J[e.scrollStrategy])||n.call(J,t,e)})))})),(0,b.EB)((()=>{var e;null==(e=n)||e.stop()}))}function K(e){function t(t){e.isActive.value=!1}te(e.activatorEl.value??e.contentEl.value,t)}function X(e,t){var n;const o=null==(n=e.root.value)?void 0:n.offsetParent,a=[...new Set([...B(e.activatorEl.value,t.contained?o:void 0),...B(e.contentEl.value,t.contained?o:void 0)])].filter((e=>!e.classList.contains("v-overlay-scroll-blocked"))),l=window.innerWidth-document.documentElement.offsetWidth,r=(e=>F(e)&&e)(o||document.documentElement);r&&e.root.value.classList.add("v-overlay--scroll-blocked"),a.forEach(((e,t)=>{e.style.setProperty("--v-body-scroll-x",(0,g.kb)(-e.scrollLeft)),e.style.setProperty("--v-body-scroll-y",(0,g.kb)(-e.scrollTop)),e.style.setProperty("--v-scrollbar-offset",(0,g.kb)(l)),e.classList.add("v-overlay-scroll-blocked")})),(0,b.EB)((()=>{a.forEach(((e,t)=>{const n=parseFloat(e.style.getPropertyValue("--v-body-scroll-x")),o=parseFloat(e.style.getPropertyValue("--v-body-scroll-y"));e.style.removeProperty("--v-body-scroll-x"),e.style.removeProperty("--v-body-scroll-y"),e.style.removeProperty("--v-scrollbar-offset"),e.classList.remove("v-overlay-scroll-blocked"),e.scrollLeft=-n,e.scrollTop=-o})),r&&e.root.value.classList.remove("v-overlay--scroll-blocked")}))}function ee(e){let t=!1,n=-1;function o(n){q((()=>{var o,a;const l=performance.now();null==(o=(a=e.updateLocation).value)||o.call(a,n);const r=performance.now()-l;t=r/(1e3/60)>2}))}te(e.activatorEl.value??e.contentEl.value,(e=>{t?(cancelAnimationFrame(n),n=requestAnimationFrame((()=>{n=requestAnimationFrame((()=>{o(e)}))}))):o(e)}))}function te(e,t){const n=[document,...B(e)];n.forEach((e=>{e.addEventListener("scroll",t,{passive:!0})})),(0,b.EB)((()=>{n.forEach((e=>{e.removeEventListener("scroll",t)}))}))}var ne=n(7041),oe=n(4906),ae=n(6183),le=n(2370),re=n(8717);function ie(){var e,t,n;if(!f.BR)return(0,b.iH)(!1);const a=(0,h.FN)("useHydration"),l=null==a||null==(e=a.root)||null==(t=e.appContext)||null==(n=t.app)?void 0:n._container,r=(0,b.iH)(!(null==l||!l.__vue_app__));return r.value||(0,o.bv)((()=>r.value=!0)),r}var se=n(1629),ce=n(4770);const ue=Symbol.for("vuetify:stack"),de=(0,b.qj)([]);function ve(e,t){const n=(0,h.FN)("useStack"),a=(0,o.f3)(ue,void 0),l=(0,b.qj)({activeChildren:new Set});(0,o.JJ)(ue,l);const r=(0,b.iH)(+t.value);(0,ce.U)(e,(()=>{var e;const o=null==(e=de.at(-1))?void 0:e[1];r.value=o?o+10:+t.value,de.push([n.uid,r.value]),null==a||a.activeChildren.add(n.uid),(0,b.EB)((()=>{const e=de.findIndex((e=>e[0]===n.uid));de.splice(e,1),null==a||a.activeChildren.delete(n.uid)}))}));const i=(0,b.iH)(!0);(0,o.m0)((()=>{var e;const t=(null==(e=de.at(-1))?void 0:e[0])===n.uid;setTimeout((()=>i.value=t))}));const s=(0,o.Fl)((()=>!l.activeChildren.size));return{globalTop:(0,b.OT)(i),localTop:s,stackStyles:(0,o.Fl)((()=>({zIndex:r.value})))}}function fe(e){const t=(0,o.Fl)((()=>{const t=e.value;if(!0===t||!f.BR)return;const n=!1===t?document.body:"string"===typeof t?document.querySelector(t):t;if(null!=n){if(!fe.cache.has(n)){const e=document.createElement("div");e.className="v-overlay-container",n.appendChild(e),fe.cache.set(n,e)}return fe.cache.get(n)}(0,o.ZK)(`Unable to locate target ${t}`)}));return{teleportTarget:t}}function pe(e){if("function"!==typeof e.getRootNode){while(e.parentNode)e=e.parentNode;return e!==document?null:document}const t=e.getRootNode();return t!==document&&t.getRootNode({composed:!0})!==document?null:t}function me(){return!0}function ye(e,t,n){if(!e||!1===ge(e,n))return!1;const o=pe(t);if("undefined"!==typeof ShadowRoot&&o instanceof ShadowRoot&&o.host===e.target)return!1;const a=("object"===typeof n.value&&n.value.include||(()=>[]))();return a.push(t),!a.some((t=>null==t?void 0:t.contains(e.target)))}function ge(e,t){const n="object"===typeof t.value&&t.value.closeConditional||me;return n(e)}function he(e,t,n){const o="function"===typeof n.value?n.value:n.value.handler;t._clickOutside.lastMousedownWasOutside&&ye(e,t,n)&&setTimeout((()=>{ge(e,n)&&o&&o(e)}),0)}function be(e,t){const n=pe(e);t(document),"undefined"!==typeof ShadowRoot&&n instanceof ShadowRoot&&t(n)}fe.cache=new WeakMap;const xe={mounted(e,t){const n=n=>he(n,e,t),o=n=>{e._clickOutside.lastMousedownWasOutside=ye(n,e,t)};be(e,(e=>{e.addEventListener("click",n,!0),e.addEventListener("mousedown",o,!0)})),e._clickOutside||(e._clickOutside={lastMousedownWasOutside:!0}),e._clickOutside[t.instance.$.uid]={onClick:n,onMousedown:o}},unmounted(e,t){e._clickOutside&&(be(e,(n=>{var o;if(!n||null==(o=e._clickOutside)||!o[t.instance.$.uid])return;const{onClick:a,onMousedown:l}=e._clickOutside[t.instance.$.uid];n.removeEventListener("click",a,!0),n.removeEventListener("mousedown",l,!0)})),delete e._clickOutside[t.instance.$.uid])}};var we=n(9888);function Ee(e){const{modelValue:t,color:n,...l}=e;return(0,o.Wm)(a.uT,{name:"fade-transition",appear:!0},{default:()=>[e.modelValue&&(0,o.Wm)("div",(0,o.dG)({class:["v-overlay__scrim",e.color.backgroundColorClasses.value],style:e.color.backgroundColorStyles.value},l),null)]})}const ke=(0,v.U)({absolute:Boolean,attach:[Boolean,String,Object],closeOnBack:{type:Boolean,default:!0},contained:Boolean,contentClass:null,contentProps:null,disabled:Boolean,noClickAnimation:Boolean,modelValue:Boolean,persistent:Boolean,scrim:{type:[String,Boolean],default:!0},zIndex:{type:[Number,String],default:2e3},...x(),...(0,k.x)(),...O(),...T(),...Z(),...(0,ne.x$)(),...(0,oe.X)()},"v-overlay"),Oe=(0,l.e)()({name:"VOverlay",directives:{ClickOutside:xe},inheritAttrs:!1,props:ke(),emits:{"click:outside":e=>!0,"update:modelValue":e=>!0,afterLeave:()=>!0},setup(e,t){let{slots:n,attrs:l,emit:s}=t;const c=(0,re.z)(e,"modelValue"),u=(0,o.Fl)({get:()=>c.value,set:t=>{t&&e.disabled||(c.value=t)}}),{teleportTarget:d}=fe((0,o.Fl)((()=>e.attach||e.contained))),{themeClasses:v}=(0,ne.ER)(e),{rtlClasses:p,isRtl:m}=(0,se.Vw)(),{hasContent:y,onAfterLeave:h}=S(e,u),x=(0,le.Y5)((0,o.Fl)((()=>"string"===typeof e.scrim?e.scrim:null))),{globalTop:E,localTop:O,stackStyles:C}=ve(u,(0,b.Vh)(e,"zIndex")),{activatorEl:A,activatorRef:B,activatorEvents:F,contentEvents:L,scrimEvents:W}=w(e,{isActive:u,isTop:O}),{dimensionStyles:R}=(0,k.$)(e),H=ie();(0,o.YP)((()=>e.disabled),(e=>{e&&(u.value=!1)}));const _=(0,b.iH)(),j=(0,b.iH)(),{contentStyles:T,updateLocation:V}=I(e,{isRtl:m,contentEl:j,activatorEl:A,isActive:u});function $(t){s("click:outside",t),e.persistent?z():u.value=!1}function D(){return u.value&&E.value}function N(t){"Escape"===t.key&&E.value&&(e.persistent?z():u.value=!1)}Q(e,{root:_,contentEl:j,activatorEl:A,isActive:u,updateLocation:V}),f.BR&&(0,o.YP)(u,(e=>{e?window.addEventListener("keydown",N):window.removeEventListener("keydown",N)}),{immediate:!0});const M=(0,ae.tv)();(0,ce.U)((()=>e.closeOnBack),(()=>{(0,ae.Kx)(M,(t=>{E.value&&u.value?(t(!1),e.persistent?z():u.value=!1):t()}))}));const Y=(0,b.iH)();function z(){e.noClickAnimation||j.value&&(0,r.j)(j.value,[{transformOrigin:"center"},{transform:"scale(1.03)"},{transformOrigin:"center"}],{duration:150,easing:i.Ly})}return(0,o.YP)((()=>u.value&&(e.absolute||e.contained)&&null==d.value),(e=>{if(e){const e=P(_.value);e&&e!==document.scrollingElement&&(Y.value=e.scrollTop)}})),(0,we.L)((()=>{var t,r;return(0,o.Wm)(o.HY,null,[null==(t=n.activator)?void 0:t.call(n,{isActive:u.value,props:(0,o.dG)({ref:B},(0,o.mx)(F.value),e.activatorProps)}),H.value&&(0,o.Wm)(o.lR,{disabled:!d.value,to:d.value},{default:()=>[y.value&&(0,o.Wm)("div",(0,o.dG)({class:["v-overlay",{"v-overlay--absolute":e.absolute||e.contained,"v-overlay--active":u.value,"v-overlay--contained":e.contained},v.value,p.value],style:[C.value,{top:(0,g.kb)(Y.value)}],ref:_},l),[(0,o.Wm)(Ee,(0,o.dG)({color:x,modelValue:u.value&&!!e.scrim},(0,o.mx)(W.value)),null),(0,o.Wm)(oe.J,{appear:!0,persisted:!0,transition:e.transition,target:A.value,onAfterLeave:()=>{h(),s("afterLeave")}},{default:()=>[(0,o.wy)((0,o.Wm)("div",(0,o.dG)({ref:j,class:["v-overlay__content",e.contentClass],style:[R.value,T.value]},(0,o.mx)(L.value),e.contentProps),[null==(r=n.default)?void 0:r.call(n,{isActive:u})]),[[a.F8,u.value],[(0,o.Q2)("click-outside"),{handler:$,closeConditional:D,include:()=>[A.value]}]])]})])]})])})),{activatorEl:A,animateClick:z,contentEl:j,globalTop:E,localTop:O,updateLocation:V}}});function Se(e){return(0,g.ei)(e,Object.keys(Oe.props))}function Ce(){const e=(0,h.FN)("useScopeId"),t=e.vnode.scopeId;return{scopeId:t?{[t]:""}:void 0}}var Ae=n(3185);const Pe=(0,l.e)()({name:"VDialog",props:{fullscreen:Boolean,retainFocus:{type:Boolean,default:!0},scrollable:Boolean,...ke({origin:"center center",scrollStrategy:"block",transition:{component:s},zIndex:2400})},emits:{"update:modelValue":e=>!0},setup(e,t){let{slots:n}=t;const a=(0,re.z)(e,"modelValue"),{scopeId:l}=Ce(),r=(0,b.iH)();function i(e){var t,n;const o=e.relatedTarget,a=e.target;if(o!==a&&null!=(t=r.value)&&t.contentEl&&null!=(n=r.value)&&n.globalTop&&![document,r.value.contentEl].includes(a)&&!r.value.contentEl.contains(a)){const e=[...r.value.contentEl.querySelectorAll('button, [href], input:not([type="hidden"]), select, textarea, [tabindex]:not([tabindex="-1"])')].filter((e=>!e.hasAttribute("disabled")&&!e.matches('[tabindex="-1"]')));if(!e.length)return;const t=e[0],n=e[e.length-1];o===t?n.focus():t.focus()}}return f.BR&&(0,o.YP)((()=>a.value&&e.retainFocus),(e=>{e?document.addEventListener("focusin",i):document.removeEventListener("focusin",i)}),{immediate:!0}),(0,o.YP)(a,(async e=>{var t,n;(await(0,o.Y3)(),e)?null==(t=r.value.contentEl)||t.focus({preventScroll:!0}):null==(n=r.value.activatorEl)||n.focus({preventScroll:!0})})),(0,we.L)((()=>{const[t]=Se(e);return(0,o.Wm)(Oe,(0,o.dG)({ref:r,class:["v-dialog",{"v-dialog--fullscreen":e.fullscreen,"v-dialog--scrollable":e.scrollable}]},t,{modelValue:a.value,"onUpdate:modelValue":e=>a.value=e,"aria-role":"dialog","aria-modal":"true",activatorProps:(0,o.dG)({"aria-haspopup":"dialog","aria-expanded":String(a.value)},e.activatorProps)},l),{activator:n.activator,default:function(){for(var e,t=arguments.length,a=new Array(t),l=0;l<t;l++)a[l]=arguments[l];return(0,o.Wm)(d.z,{root:!0},{default:()=>[null==(e=n.default)?void 0:e.call(n,...a)]})}})})),(0,Ae.F)({},r)}})},5310:function(e,t,n){n.d(t,{o:function(){return k}});n(7658);var o=n(1138),a=n(7139),l=n(3396),r=n(320);const i=["sm","md","lg","xl","xxl"],s=["start","end","center"],c=["space-between","space-around","space-evenly"];function u(e,t){return i.reduce(((n,o)=>(n[e+(0,a.kC)(o)]=t(),n)),{})}const d=[...s,"baseline","stretch"],v=e=>d.includes(e),f=u("align",(()=>({type:String,default:null,validator:v}))),p=[...s,...c],m=e=>p.includes(e),y=u("justify",(()=>({type:String,default:null,validator:m}))),g=[...s,...c,"stretch"],h=e=>g.includes(e),b=u("alignContent",(()=>({type:String,default:null,validator:h}))),x={align:Object.keys(f),justify:Object.keys(y),alignContent:Object.keys(b)},w={align:"align",justify:"justify",alignContent:"align-content"};function E(e,t,n){let o=w[e];if(null!=n){if(t){const n=t.replace(e,"");o+=`-${n}`}return o+=`-${n}`,o.toLowerCase()}}const k=(0,r.a)({name:"VRow",props:{dense:Boolean,noGutters:Boolean,align:{type:String,default:null,validator:v},...f,justify:{type:String,default:null,validator:m},...y,alignContent:{type:String,default:null,validator:h},...b,...(0,o.Q)()},setup(e,t){let{slots:n}=t;const o=(0,l.Fl)((()=>{const t=[];let n;for(n in x)x[n].forEach((o=>{const a=e[o],l=E(n,o,a);l&&t.push(l)}));return t.push({"v-row--no-gutters":e.noGutters,"v-row--dense":e.dense,[`align-${e.align}`]:e.align,[`justify-${e.justify}`]:e.justify,[`align-content-${e.alignContent}`]:e.alignContent}),t}));return()=>{var t;return(0,l.h)(e.tag,{class:["v-row",o.value]},null==(t=n.default)?void 0:t.call(n))}}})},3185:function(e,t,n){n.d(t,{F:function(){return a}});n(7658);const o=Symbol("Forwarded refs");function a(e){for(var t=arguments.length,n=new Array(t>1?t-1:0),a=1;a<t;a++)n[a-1]=arguments[a];return e[o]=n,new Proxy(e,{get(e,t){if(Reflect.has(e,t))return Reflect.get(e,t);for(const o of n)if(o.value&&Reflect.has(o.value,t)){const e=Reflect.get(o.value,t);return"function"===typeof e?e.bind(o.value):e}},getOwnPropertyDescriptor(e,t){const a=Reflect.getOwnPropertyDescriptor(e,t);if(a)return a;if("symbol"!==typeof t&&!t.startsWith("__")){for(const e of n){if(!e.value)continue;const n=Reflect.getOwnPropertyDescriptor(e.value,t);if(n)return n;if("_"in e.value&&"setupState"in e.value._){const n=Reflect.getOwnPropertyDescriptor(e.value._.setupState,t);if(n)return n}}for(const e of n){let n=e.value&&Object.getPrototypeOf(e.value);while(n){const e=Reflect.getOwnPropertyDescriptor(n,t);if(e)return e;n=Object.getPrototypeOf(n)}}for(const e of n){const n=e.value&&e.value[o];if(!n)continue;const a=n.slice();while(a.length){const e=a.shift(),n=Reflect.getOwnPropertyDescriptor(e.value,t);if(n)return n;const l=e.value&&e.value[o];l&&a.push(...l)}}}}})}},3122:function(e,t,n){n.d(t,{G:function(){return a},j:function(){return l}});var o=n(6309);function a(e){const t=e.getBoundingClientRect(),n=getComputedStyle(e),a=n.transform;if(a){let l,r,i,s,c;if(a.startsWith("matrix3d("))l=a.slice(9,-1).split(/, /),r=+l[0],i=+l[5],s=+l[12],c=+l[13];else{if(!a.startsWith("matrix("))return new o.x(t);l=a.slice(7,-1).split(/, /),r=+l[0],i=+l[3],s=+l[4],c=+l[5]}const u=n.transformOrigin,d=t.x-s-(1-r)*parseFloat(u),v=t.y-c-(1-i)*parseFloat(u.slice(u.indexOf(" ")+1)),f=r?t.width/r:e.offsetWidth+1,p=i?t.height/i:e.offsetHeight+1;return new o.x({x:d,y:v,width:f,height:p})}return new o.x(t)}function l(e,t,n){if("undefined"===typeof e.animate)return{finished:Promise.resolve()};const o=e.animate(t,n);return"undefined"===typeof o.finished&&(o.finished=new Promise((e=>{o.onfinish=()=>{e(o)}}))),o}},6309:function(e,t,n){n.d(t,{p:function(){return a},x:function(){return o}});class o{constructor(e){let{x:t,y:n,width:o,height:a}=e;this.x=t,this.y=n,this.width=o,this.height=a}get top(){return this.y}get bottom(){return this.y+this.height}get left(){return this.x}get right(){return this.x+this.width}}function a(e,t){return{x:{before:Math.max(0,t.left-e.left),after:Math.max(0,e.right-t.right)},y:{before:Math.max(0,t.top-e.top),after:Math.max(0,e.bottom-t.bottom)}}}},8587:function(e,t,n){n.d(t,{Ly:function(){return o},uX:function(){return a},x0:function(){return l}});const o="cubic-bezier(0.4, 0, 0.2, 1)",a="cubic-bezier(0.0, 0, 0.2, 1)",l="cubic-bezier(0.4, 0, 1, 1)"}}]);
//# sourceMappingURL=295.66befcfa.js.map