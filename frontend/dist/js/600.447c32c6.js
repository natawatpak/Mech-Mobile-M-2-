"use strict";(self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[]).push([[600],{7144:function(e,l,t){t.r(l),t.d(l,{default:function(){return ve}});var a=t(3396);const n={class:"home"};function o(e,l,t,o,r,i){const s=(0,a.up)("CallMech");return(0,a.wg)(),(0,a.iD)("div",n,[(0,a.Wm)(s)])}var r=t(7139),i=t(702),s=t(5661),u=t(6572),d=t(1888),c=t(8302),p=t(6233),m=t(4960),v=t(8717),f=t(3766),g=t(320),h=t(9888),b=t(131);const x=(0,f.U)({indeterminate:Boolean,indeterminateIcon:{type:m.lE,default:"$checkboxIndeterminate"},...(0,p.$9)({falseIcon:"$checkboxOff",trueIcon:"$checkboxOn"})},"v-checkbox-btn"),C=(0,g.a)({name:"VCheckboxBtn",props:x(),emits:{"update:modelValue":e=>!0,"update:indeterminate":e=>!0},setup(e,l){let{slots:t}=l;const n=(0,v.z)(e,"indeterminate"),o=(0,v.z)(e,"modelValue");function r(e){n.value&&(n.value=!1)}const i=(0,a.Fl)((()=>e.indeterminate?e.indeterminateIcon:e.falseIcon)),s=(0,a.Fl)((()=>e.indeterminate?e.indeterminateIcon:e.trueIcon));return(0,h.L)((()=>(0,a.Wm)(p.g5,(0,a.dG)(e,{modelValue:o.value,"onUpdate:modelValue":[e=>o.value=e,r],class:"v-checkbox-btn",type:"checkbox",inline:!0,falseIcon:i.value,trueIcon:s.value,"aria-checked":e.indeterminate?"mixed":void 0}),t))),{}}});function w(e){return(0,b.ei)(e,Object.keys(C.props))}var k=t(8969),y=t(7514);const W=(0,g.a)({name:"VCheckbox",inheritAttrs:!1,props:{...(0,c.co)(),...x()},emits:{"update:focused":e=>!0},setup(e,l){let{attrs:t,slots:n}=l;const{isFocused:o,focus:r,blur:i}=(0,k.K)(e),s=(0,y.sq)(),u=(0,a.Fl)((()=>e.id||`checkbox-${s}`));return(0,h.L)((()=>{const[l,s]=(0,b.An)(t),[d,p]=(0,c.PE)(e),[m,v]=w(e);return(0,a.Wm)(c.q8,(0,a.dG)({class:"v-checkbox"},l,d,{id:u.value,focused:o.value}),{...n,default:e=>{let{id:l,isDisabled:t,isReadonly:o}=e;return(0,a.Wm)(C,(0,a.dG)(m,{id:l.value,disabled:t.value,readonly:o.value},s,{onFocus:r,onBlur:i}),n)}})})),{}}});var _=t(3849),V=t(4075),I=t(4357),F=t(3185),S=t(4870);const U=(0,g.a)({name:"VForm",props:{...(0,I.vC)()},emits:{"update:modelValue":e=>!0,submit:e=>!0},setup(e,l){let{slots:t,emit:n}=l;const o=(0,I.Np)(e),r=(0,S.iH)();function i(e){e.preventDefault(),o.reset()}function s(e){const l=e,t=o.validate();l.then=t.then.bind(t),l.catch=t.catch.bind(t),l.finally=t.finally.bind(t),n("submit",l),l.defaultPrevented||t.then((e=>{let{valid:l}=e;var t;l&&(null==(t=r.value)||t.submit())})),l.preventDefault()}return(0,h.L)((()=>{var e;return(0,a.Wm)("form",{ref:r,class:"v-form",novalidate:!0,onReset:i,onSubmit:s},[null==(e=t.default)?void 0:e.call(t,o)])})),(0,F.F)(o,r)}});var B=t(9234),P=t(5310),L=t(3289),H=t(4647),D=t(8777),z=t(9242),$=t(8952),A=t(6308),Y=t(7302);const E=(0,g.a)({name:"VFieldLabel",props:{floating:Boolean},setup(e,l){let{slots:t}=l;return(0,h.L)((()=>(0,a.Wm)(Y.J,{class:["v-field-label",{"v-field-label--floating":e.floating}],"aria-hidden":e.floating||void 0},t))),{}}});var M=t(1710),R=t(7041),G=t(2370),N=t(3122),Z=t(8587);const j=["underlined","outlined","filled","solo","plain"],T=(0,f.U)({appendInnerIcon:m.lE,bgColor:String,clearable:Boolean,clearIcon:{type:m.lE,default:"$clear"},active:Boolean,color:String,dirty:Boolean,disabled:Boolean,error:Boolean,label:String,persistentClear:Boolean,prependInnerIcon:m.lE,reverse:Boolean,singleLine:Boolean,variant:{type:String,default:"filled",validator:e=>j.includes(e)},"onClick:clear":b.as,"onClick:appendInner":b.as,"onClick:prependInner":b.as,...(0,R.x$)(),...(0,M.fF)()},"v-field"),q=(0,g.e)()({name:"VField",inheritAttrs:!1,props:{id:String,...(0,k.B)(),...T()},emits:{"click:control":e=>!0,"update:focused":e=>!0,"update:modelValue":e=>!0},setup(e,l){let{attrs:t,emit:n,slots:o}=l;const{themeClasses:r}=(0,R.ER)(e),{loaderClasses:i}=(0,M.U2)(e),{focusClasses:s,isFocused:u,focus:d,blur:c}=(0,k.K)(e),{InputIcon:p}=(0,A.v)(e),m=(0,a.Fl)((()=>e.dirty||e.active)),v=(0,a.Fl)((()=>!e.singleLine&&!(!e.label&&!o.label))),f=(0,y.sq)(),g=(0,a.Fl)((()=>e.id||`input-${f}`)),x=(0,S.iH)(),C=(0,S.iH)(),w=(0,S.iH)(),{backgroundColorClasses:W,backgroundColorStyles:_}=(0,G.Y5)((0,S.Vh)(e,"bgColor")),{textColorClasses:V,textColorStyles:I}=(0,G.rY)((0,a.Fl)((()=>m.value&&u.value&&!e.error&&!e.disabled?e.color:void 0)));(0,a.YP)(m,(e=>{if(v.value){const l=x.value.$el,t=C.value.$el,a=(0,N.G)(l),n=t.getBoundingClientRect(),o=n.x-a.x,r=n.y-a.y-(a.height/2-n.height/2),i=n.width/.75,s=Math.abs(i-a.width)>1?{maxWidth:(0,b.kb)(i)}:void 0,u=getComputedStyle(l),d=getComputedStyle(t),c=1e3*parseFloat(u.transitionDuration)||150,p=parseFloat(d.getPropertyValue("--v-field-label-scale")),m=d.getPropertyValue("color");l.style.visibility="visible",t.style.visibility="hidden",(0,N.j)(l,{transform:`translate(${o}px, ${r}px) scale(${p})`,color:m,...s},{duration:c,easing:Z.Ly,direction:e?"normal":"reverse"}).finished.then((()=>{l.style.removeProperty("visibility"),t.style.removeProperty("visibility")}))}}),{flush:"post"});const F=(0,a.Fl)((()=>({isActive:m,isFocused:u,controlRef:w,blur:c,focus:d})));function U(e){e.target!==document.activeElement&&e.preventDefault(),n("click:control",e)}return(0,h.L)((()=>{var l,n,u;const f="outlined"===e.variant,h=o["prepend-inner"]||e.prependInnerIcon,b=!(!e.clearable&&!o.clear),w=!!(o["append-inner"]||e.appendInnerIcon||b),k=o.label?o.label({label:e.label,props:{for:g.value}}):e.label;return(0,a.Wm)("div",(0,a.dG)({class:["v-field",{"v-field--active":m.value,"v-field--appended":w,"v-field--disabled":e.disabled,"v-field--dirty":e.dirty,"v-field--error":e.error,"v-field--has-background":!!e.bgColor,"v-field--persistent-clear":e.persistentClear,"v-field--prepended":h,"v-field--reverse":e.reverse,"v-field--single-line":e.singleLine,"v-field--no-label":!k,[`v-field--variant-${e.variant}`]:!0},r.value,W.value,s.value,i.value],style:[_.value,I.value],onClick:U},t),[(0,a.Wm)("div",{class:"v-field__overlay"},null),(0,a.Wm)(M.rD,{name:"v-field",active:e.loading,color:e.error?"error":e.color},{default:o.loader}),h&&(0,a.Wm)("div",{key:"prepend",class:"v-field__prepend-inner"},[e.prependInnerIcon&&(0,a.Wm)(p,{key:"prepend-icon",name:"prependInner"},null),null==(l=o["prepend-inner"])?void 0:l.call(o,F.value)]),(0,a.Wm)("div",{class:"v-field__field","data-no-activator":""},[["solo","filled"].includes(e.variant)&&v.value&&(0,a.Wm)(E,{key:"floating-label",ref:C,class:[V.value],floating:!0,for:g.value},{default:()=>[k]}),(0,a.Wm)(E,{ref:x,for:g.value},{default:()=>[k]}),null==(n=o.default)?void 0:n.call(o,{...F.value,props:{id:g.value,class:"v-field__input"},focus:d,blur:c})]),b&&(0,a.Wm)($.Zq,{key:"clear"},{default:()=>[(0,a.wy)((0,a.Wm)("div",{class:"v-field__clearable"},[o.clear?o.clear():(0,a.Wm)(p,{name:"clear"},null)]),[[z.F8,e.dirty]])]}),w&&(0,a.Wm)("div",{key:"append",class:"v-field__append-inner"},[null==(u=o["append-inner"])?void 0:u.call(o,F.value),e.appendInnerIcon&&(0,a.Wm)(p,{key:"append-icon",name:"appendInner"},null)]),(0,a.Wm)("div",{class:["v-field__outline",V.value]},[f&&(0,a.Wm)(a.HY,null,[(0,a.Wm)("div",{class:"v-field__outline__start"},null),v.value&&(0,a.Wm)("div",{class:"v-field__outline__notch"},[(0,a.Wm)(E,{ref:C,floating:!0,for:g.value},{default:()=>[k]})]),(0,a.Wm)("div",{class:"v-field__outline__end"},null)]),["plain","underlined"].includes(e.variant)&&v.value&&(0,a.Wm)(E,{ref:C,floating:!0,for:g.value},{default:()=>[k]})])])})),{controlRef:w}}});function K(e){const l=Object.keys(q.props).filter((e=>!(0,b.F7)(e)));return(0,b.ei)(e,l)}var O=t(4906);const J=(0,g.a)({name:"VCounter",functional:!0,props:{active:Boolean,max:[Number,String],value:{type:[Number,String],default:0},...(0,O.X)({transition:{component:$.cu}})},setup(e,l){let{slots:t}=l;const n=(0,a.Fl)((()=>e.max?`${e.value} / ${e.max}`:String(e.value)));return(0,h.L)((()=>(0,a.Wm)(O.J,{transition:e.transition},{default:()=>[(0,a.wy)((0,a.Wm)("div",{class:"v-counter"},[t.default?t.default({counter:n.value,max:e.max,value:e.value}):n.value]),[[z.F8,e.active]])]}))),{}}});var Q=t(7052);const X=["color","file","time","date","datetime-local","week","month"],ee=(0,f.U)({autofocus:Boolean,counter:[Boolean,Number,String],counterValue:Function,hint:String,persistentHint:Boolean,prefix:String,placeholder:String,persistentPlaceholder:Boolean,persistentCounter:Boolean,suffix:String,type:{type:String,default:"text"},...(0,c.co)(),...T()},"v-text-field"),le=(0,g.e)()({name:"VTextField",directives:{Intersect:Q.Z},inheritAttrs:!1,props:ee(),emits:{"click:control":e=>!0,"click:input":e=>!0,"update:focused":e=>!0,"update:modelValue":e=>!0},setup(e,l){let{attrs:t,emit:n,slots:o}=l;const r=(0,v.z)(e,"modelValue"),{isFocused:i,focus:s,blur:u}=(0,k.K)(e),d=(0,a.Fl)((()=>"function"===typeof e.counterValue?e.counterValue(r.value):(r.value??"").toString().length)),p=(0,a.Fl)((()=>t.maxlength?t.maxlength:!e.counter||"number"!==typeof e.counter&&"string"!==typeof e.counter?void 0:e.counter));function m(l,t){var a,n;e.autofocus&&l&&(null==(a=t[0].target)||null==(n=a.focus)||n.call(a))}const f=(0,S.iH)(),g=(0,S.iH)(),x=(0,S.iH)(),C=(0,a.Fl)((()=>X.includes(e.type)||e.persistentPlaceholder||i.value)),w=(0,a.Fl)((()=>e.messages.length?e.messages:i.value||e.persistentHint?e.hint:""));function y(){var e;x.value!==document.activeElement&&(null==(e=x.value)||e.focus());i.value||s()}function W(e){y(),n("click:control",e)}function _(l){l.stopPropagation(),y(),(0,a.Y3)((()=>{r.value=null,(0,b.dr)(e["onClick:clear"],l)}))}function V(e){r.value=e.target.value}return(0,h.L)((()=>{const l=!!(o.counter||e.counter||e.counterValue),s=!(!l&&!o.details),[v,h]=(0,b.An)(t),[{modelValue:k,...I}]=(0,c.PE)(e),[F]=K(e);return(0,a.Wm)(c.q8,(0,a.dG)({ref:f,modelValue:r.value,"onUpdate:modelValue":e=>r.value=e,class:["v-text-field",{"v-text-field--prefixed":e.prefix,"v-text-field--suffixed":e.suffix,"v-text-field--flush-details":["plain","underlined"].includes(e.variant)}],"onClick:prepend":e["onClick:prepend"],"onClick:append":e["onClick:append"]},v,I,{focused:i.value,messages:w.value}),{...o,default:l=>{let{id:t,isDisabled:s,isDirty:d,isReadonly:c,isValid:p}=l;return(0,a.Wm)(q,(0,a.dG)({ref:g,onMousedown:e=>{e.target!==x.value&&e.preventDefault()},"onClick:control":W,"onClick:clear":_,"onClick:prependInner":e["onClick:prependInner"],"onClick:appendInner":e["onClick:appendInner"],role:"textbox"},F,{id:t.value,active:C.value||d.value,dirty:d.value||e.dirty,focused:i.value,error:!1===p.value}),{...o,default:l=>{let{props:{class:t,...i}}=l;const d=(0,a.wy)((0,a.Wm)("input",(0,a.dG)({ref:x,value:r.value,onInput:V,autofocus:e.autofocus,readonly:c.value,disabled:s.value,name:e.name,placeholder:e.placeholder,size:1,type:e.type,onFocus:y,onBlur:u},i,h),null),[[(0,a.Q2)("intersect"),{handler:m},null,{once:!0}]]);return(0,a.Wm)(a.HY,null,[e.prefix&&(0,a.Wm)("span",{class:"v-text-field__prefix"},[e.prefix]),o.default?(0,a.Wm)("div",{class:t,onClick:e=>n("click:input",e),"data-no-activator":""},[o.default(),d]):(0,a.Ho)(d,{class:t}),e.suffix&&(0,a.Wm)("span",{class:"v-text-field__suffix"},[e.suffix])])}})},details:s?t=>{var n;return(0,a.Wm)(a.HY,null,[null==(n=o.details)?void 0:n.call(o,t),l&&(0,a.Wm)(a.HY,null,[(0,a.Wm)("span",null,null),(0,a.Wm)(J,{active:e.persistentCounter||i.value,value:d.value,max:p.value},o.counter)])])}:void 0})})),(0,F.F)({},f,g,x)}});const te=(0,g.a)({name:"VTextarea",directives:{Intersect:Q.Z},inheritAttrs:!1,props:{autoGrow:Boolean,autofocus:Boolean,counter:[Boolean,Number,String],counterValue:Function,hint:String,persistentHint:Boolean,prefix:String,placeholder:String,persistentPlaceholder:Boolean,persistentCounter:Boolean,noResize:Boolean,rows:{type:[Number,String],default:5,validator:e=>!isNaN(parseFloat(e))},maxRows:{type:[Number,String],validator:e=>!isNaN(parseFloat(e))},suffix:String,...(0,c.co)(),...T()},emits:{"click:control":e=>!0,"update:focused":e=>!0,"update:modelValue":e=>!0},setup(e,l){let{attrs:t,emit:n,slots:o}=l;const r=(0,v.z)(e,"modelValue"),{isFocused:i,focus:s,blur:u}=(0,k.K)(e),d=(0,a.Fl)((()=>"function"===typeof e.counterValue?e.counterValue(r.value):(r.value||"").toString().length)),p=(0,a.Fl)((()=>t.maxlength?t.maxlength:!e.counter||"number"!==typeof e.counter&&"string"!==typeof e.counter?void 0:e.counter));function m(l,t){var a,n;e.autofocus&&l&&(null==(a=t[0].target)||null==(n=a.focus)||n.call(a))}const f=(0,S.iH)(),g=(0,S.iH)(),x=(0,S.iH)(""),C=(0,S.iH)(),w=(0,a.Fl)((()=>i.value||e.persistentPlaceholder)),y=(0,a.Fl)((()=>e.messages.length?e.messages:w.value||e.persistentHint?e.hint:""));function W(){var e;C.value!==document.activeElement&&(null==(e=C.value)||e.focus());i.value||s()}function _(e){W(),n("click:control",e)}function V(l){l.stopPropagation(),W(),(0,a.Y3)((()=>{r.value="",(0,b.dr)(e["onClick:clear"],l)}))}function I(e){r.value=e.target.value}const U=(0,S.iH)();function B(){e.autoGrow&&(0,a.Y3)((()=>{if(!U.value||!g.value)return;const l=getComputedStyle(U.value),t=getComputedStyle(g.value.$el),a=parseFloat(l.getPropertyValue("--v-field-padding-top"))+parseFloat(l.getPropertyValue("--v-input-padding-top"))+parseFloat(l.getPropertyValue("--v-field-padding-bottom")),n=U.value.scrollHeight,o=parseFloat(l.lineHeight),r=Math.max(parseFloat(e.rows)*o+a,parseFloat(t.getPropertyValue("--v-input-control-height"))),i=parseFloat(e.maxRows)*o+a||1/0;x.value=(0,b.kb)((0,b.uZ)(n??0,r,i))}))}let P;return(0,a.bv)(B),(0,a.YP)(r,B),(0,a.YP)((()=>e.rows),B),(0,a.YP)((()=>e.maxRows),B),(0,a.YP)((()=>e.density),B),(0,a.YP)(U,(e=>{var l;e?(P=new ResizeObserver(B),P.observe(U.value)):null==(l=P)||l.disconnect()})),(0,a.Jd)((()=>{var e;null==(e=P)||e.disconnect()})),(0,h.L)((()=>{const l=!!(o.counter||e.counter||e.counterValue),n=!(!l&&!o.details),[s,v]=(0,b.An)(t),[{modelValue:h,...k}]=(0,c.PE)(e),[F]=K(e);return(0,a.Wm)(c.q8,(0,a.dG)({ref:f,modelValue:r.value,"onUpdate:modelValue":e=>r.value=e,class:["v-textarea v-text-field",{"v-textarea--prefixed":e.prefix,"v-textarea--suffixed":e.suffix,"v-text-field--prefixed":e.prefix,"v-text-field--suffixed":e.suffix,"v-textarea--auto-grow":e.autoGrow,"v-textarea--no-resize":e.noResize||e.autoGrow,"v-text-field--flush-details":["plain","underlined"].includes(e.variant)}],"onClick:prepend":e["onClick:prepend"],"onClick:append":e["onClick:append"]},s,k,{focused:i.value,messages:y.value}),{...o,default:l=>{let{isDisabled:t,isDirty:n,isReadonly:s,isValid:d}=l;return(0,a.Wm)(q,(0,a.dG)({ref:g,style:{"--v-textarea-control-height":x.value},"onClick:control":_,"onClick:clear":V,"onClick:prependInner":e["onClick:prependInner"],"onClick:appendInner":e["onClick:appendInner"],role:"textbox"},F,{active:w.value||n.value,dirty:n.value||e.dirty,focused:i.value,error:!1===d.value}),{...o,default:l=>{let{props:{class:n,...o}}=l;return(0,a.Wm)(a.HY,null,[e.prefix&&(0,a.Wm)("span",{class:"v-text-field__prefix"},[e.prefix]),(0,a.wy)((0,a.Wm)("textarea",(0,a.dG)({ref:C,class:n,value:r.value,onInput:I,autofocus:e.autofocus,readonly:s.value,disabled:t.value,placeholder:e.placeholder,rows:e.rows,name:e.name,onFocus:W,onBlur:u},o,v),null),[[(0,a.Q2)("intersect"),{handler:m},null,{once:!0}]]),e.autoGrow&&(0,a.wy)((0,a.Wm)("textarea",{class:[n,"v-textarea__sizer"],"onUpdate:modelValue":e=>r.value=e,ref:U,readonly:!0,"aria-hidden":"true"},null),[[z.nr,r.value]]),e.suffix&&(0,a.Wm)("span",{class:"v-text-field__suffix"},[e.suffix])])}})},details:n?t=>{var n;return(0,a.Wm)(a.HY,null,[null==(n=o.details)?void 0:n.call(o,t),l&&(0,a.Wm)(a.HY,null,[(0,a.Wm)("span",null,null),(0,a.Wm)(J,{active:e.persistentCounter||i.value,value:d.value,max:p.value},o.counter)])])}:void 0})})),(0,F.F)({},f,g,C)}}),ae={class:"pa-5"},ne=(0,a._)("div",{style:{width:"auto",height:"15rem"},id:"map"},null,-1),oe=(0,a._)("br",null,null,-1),re={class:"text-center"};function ie(e,l,t,n,o,c){const p=(0,a.up)("v-card-action");return(0,a.wg)(),(0,a.iD)("div",ae,[(0,a.Wm)(s._,{"text-left":"","pa-4":"","d-flex":"","justify-left":"","align-center":""},{default:(0,a.w5)((()=>[ne])),_:1}),(0,a.Wm)(s._,{variant:"tonal",onClick:l[0]||(l[0]=e=>c.locatorButtonPressed()),class:"text-left pa-4 d-flex justify-left align-center"},{default:(0,a.w5)((()=>[(0,a.Wm)(L.t,{icon:"mdi-pin"}),(0,a._)("section",null,[(0,a.Wm)(u.E,null,{default:(0,a.w5)((()=>[(0,a.Uk)("Current Location")])),_:1}),(0,a.Wm)(d.Z,null,{default:(0,a.w5)((()=>[(0,a.Uk)((0,r.zw)(o.currentLocation.lat+", "+o.currentLocation.lng),1)])),_:1})])])),_:1}),(0,a.Wm)(B.C,{class:"my-5"}),(0,a.Wm)(s._,{variant:"tonal",class:"text-left pa-4"},{default:(0,a.w5)((()=>[(0,a.Wm)(P.o,{class:"pl-5"},{default:(0,a.w5)((()=>[(0,a.Wm)(u.E,null,{default:(0,a.w5)((()=>[(0,a.Uk)("Car detail")])),_:1}),(0,a.Wm)(B.C),(0,a.Wm)(i.T,{class:"mx-1",onClick:l[1]||(l[1]=e=>o.selectCarModal=!0)},{default:(0,a.w5)((()=>[(0,a.Uk)("choose from preset")])),_:1})])),_:1}),(0,a._)("section",null,[(0,a.Wm)(d.Z,null,{default:(0,a.w5)((()=>[(0,a.Uk)(" Type "+(0,r.zw)(o.car.type)+", Brand "+(0,r.zw)(o.car.brand)+" ",1),oe,(0,a.Uk)(" License plate: "+(0,r.zw)(o.car.plate),1)])),_:1})])])),_:1}),(0,a.Wm)(_.B,{modelValue:o.selectCarModal,"onUpdate:modelValue":l[3]||(l[3]=e=>o.selectCarModal=e),width:"800"},{default:(0,a.w5)((()=>[(0,a.Wm)(s._,null,{default:(0,a.w5)((()=>[(0,a.Wm)(u.E,null,{default:(0,a.w5)((()=>[(0,a.Uk)("Car preset")])),_:1}),((0,a.wg)(!0),(0,a.iD)(a.HY,null,(0,a.Ko)(o.cars,((e,l)=>((0,a.wg)(),(0,a.j4)(d.Z,{key:l,onClick:l=>c.selectCar(e)},{default:(0,a.w5)((()=>[(0,a.Uk)((0,r.zw)(e.brand)+" : "+(0,r.zw)(e.plate),1)])),_:2},1032,["onClick"])))),128)),(0,a.Wm)(V.J),(0,a.Wm)(p,null,{default:(0,a.w5)((()=>[(0,a.Wm)(i.T,{block:"",class:"mx-1","prepend-icon":"mdi-plus",onClick:l[2]||(l[2]=e=>o.dialog2=!0)},{default:(0,a.w5)((()=>[(0,a.Uk)("Add new preset")])),_:1})])),_:1})])),_:1})])),_:1},8,["modelValue"]),(0,a.Wm)(_.B,{modelValue:o.dialog2,"onUpdate:modelValue":l[10]||(l[10]=e=>o.dialog2=e),width:"800"},{default:(0,a.w5)((()=>[(0,a.Wm)(s._,null,{default:(0,a.w5)((()=>[(0,a.Wm)(u.E,null,{default:(0,a.w5)((()=>[(0,a.Uk)("New preset")])),_:1}),(0,a.Wm)(U,{ref:"form",modelValue:e.valid,"onUpdate:modelValue":l[7]||(l[7]=l=>e.valid=l)},{default:(0,a.w5)((()=>[(0,a.Wm)(le,{modelValue:o.newCar.type,"onUpdate:modelValue":l[4]||(l[4]=e=>o.newCar.type=e),class:"px-4",label:"Car type",required:""},null,8,["modelValue"]),(0,a.Wm)(le,{modelValue:o.newCar.brand,"onUpdate:modelValue":l[5]||(l[5]=e=>o.newCar.brand=e),class:"px-4",label:"Brand",required:""},null,8,["modelValue"]),(0,a.Wm)(le,{modelValue:o.newCar.plate,"onUpdate:modelValue":l[6]||(l[6]=e=>o.newCar.plate=e),class:"px-4",label:"License Plate",required:""},null,8,["modelValue"])])),_:1},8,["modelValue"]),(0,a.Wm)(P.o,{justify:"end",class:"pa-6"},{default:(0,a.w5)((()=>[(0,a.Wm)(p,null,{default:(0,a.w5)((()=>[(0,a.Wm)(i.T,{class:"mx-1",variant:"tonal",color:"error",onClick:l[8]||(l[8]=e=>o.dialog2=!1)},{default:(0,a.w5)((()=>[(0,a.Uk)(" Cancel ")])),_:1}),(0,a.Wm)(i.T,{class:"mx-1",color:"blue",variant:"tonal",onClick:l[9]||(l[9]=e=>(c.addNewCar(),o.dialog2=!1))},{default:(0,a.w5)((()=>[(0,a.Uk)(" Save ")])),_:1})])),_:1})])),_:1})])),_:1})])),_:1},8,["modelValue"]),(0,a.Wm)(B.C,{class:"my-5"}),(0,a.Wm)(s._,{variant:"tonal",class:"text-left pa-4"},{default:(0,a.w5)((()=>[(0,a.Wm)(u.E,null,{default:(0,a.w5)((()=>[(0,a.Uk)("Problems")])),_:1}),(0,a.Wm)(d.Z,{class:"d-flex"},{default:(0,a.w5)((()=>[(0,a.Wm)(H.i,null,{default:(0,a.w5)((()=>[(0,a.Wm)(D.l,{class:"pa-0 ma-0"},{default:(0,a.w5)((()=>[(0,a.Wm)(W,{modelValue:o.problem,"onUpdate:modelValue":l[11]||(l[11]=e=>o.problem=e),value:"1",label:"Problem1"},null,8,["modelValue"])])),_:1}),(0,a.Wm)(D.l,{class:"pa-0 ma-0"},{default:(0,a.w5)((()=>[(0,a.Wm)(W,{modelValue:o.problem,"onUpdate:modelValue":l[12]||(l[12]=e=>o.problem=e),value:"3",label:"Problem3"},null,8,["modelValue"])])),_:1})])),_:1}),(0,a.Wm)(H.i,null,{default:(0,a.w5)((()=>[(0,a.Wm)(D.l,{class:"pa-0 ma-0"},{default:(0,a.w5)((()=>[(0,a.Wm)(W,{modelValue:o.problem,"onUpdate:modelValue":l[13]||(l[13]=e=>o.problem=e),value:"2",label:"Problem2"},null,8,["modelValue"])])),_:1}),(0,a.Wm)(D.l,{class:"pa-0 ma-0"},{default:(0,a.w5)((()=>[(0,a.Wm)(W,{modelValue:o.problem,"onUpdate:modelValue":l[14]||(l[14]=e=>o.problem=e),value:"4",label:"Problem4"},null,8,["modelValue"])])),_:1})])),_:1})])),_:1})])),_:1}),(0,a.Wm)(B.C,{class:"my-5"}),(0,a.Wm)(s._,{variant:"tonal",class:"text-left pa-4"},{default:(0,a.w5)((()=>[(0,a.Wm)(u.E,null,{default:(0,a.w5)((()=>[(0,a.Uk)("Details")])),_:1}),(0,a.Uk)(" "+(0,r.zw)(o.description)+" ",1),(0,a.Wm)(te,{outlined:"",name:"description",modelValue:o.description,"onUpdate:modelValue":l[15]||(l[15]=e=>o.description=e),rows:"5",label:"Description","bg-color":"white"},null,8,["modelValue"])])),_:1}),(0,a.Wm)(B.C,{class:"my-5"}),(0,a.Wm)(B.C,{class:"my-5"}),(0,a._)("section",re,[(0,a._)("div",{class:"text-decoration-none",onClick:l[16]||(l[16]=e=>c.addTicket())},[(0,a.Wm)(i.T,{class:"mx-1",variant:"tonal",color:"blue"},{default:(0,a.w5)((()=>[(0,a.Uk)("Find Service")])),_:1})])])])}t(7658);var se={data(){return{map:null,currentLocation:{lat:0,lng:0},car:{id:"",type:"",brand:"",plate:""},cars:[],selectCarModal:!1,dialog2:!1,newCar:{type:"",brand:"",plate:""},problem:[],description:"",files:[]}},mounted(){this.initMap(),this.locatorButtonPressed(),this.setMarker(this.mapCenter,"A"),this.getCarList()},methods:{locatorButtonPressed(){const e=e=>{const l=e.coords.latitude,t=e.coords.longitude;console.log(l+", "+t),this.currentLocation.lat=l,this.currentLocation.lng=t,this.setMarker(this.currentLocation),this.map.setCenter(this.currentLocation)},l=e=>{console.log(e)};navigator.geolocation.getCurrentPosition(e,l)},initMap(){this.map=new google.maps.Map(document.getElementById("map"),{center:this.currentLocation,zoom:15,maxZoom:20,minZoom:3,streetViewControl:!1,mapTypeControl:!1,fullscreenControl:!1,zoomControl:!0})},setMarker(e){new google.maps.Marker({position:e,map:this.map})},selectCar(e){this.car=e,this.selectCarModal=!1},getCarList(){const e=new URLSearchParams({cusID:sessionStorage.getItem("cusID")});console.log(sessionStorage.getItem("cusID")),this.axios.post(this.$backendApi+"customer/get-car-list",e,{headers:{Authorization:sessionStorage.getItem("jwt")}}).then((e=>{console.log(e.data),this.cars=e.data})),console.log(sessionStorage.getItem("jwt"))},addTicket(){const e=new URLSearchParams({cusID:sessionStorage.getItem("cusID"),carID:this.car.id,problem:this.problem,status:"Active",lng:this.currentLocation.lng,lat:this.currentLocation.lat,description:this.description});this.axios.post(this.$backendApi+"customer/add-ticket",e,{headers:{Authorization:sessionStorage.getItem("jwt")}}).then((e=>{console.log(e.data),sessionStorage.setItem("ticketID",e.data.ticketID),this.$router.push("/progress")}))},addNewCar(){console.log("clicked");const e=new URLSearchParams({cusID:sessionStorage.getItem("cusID"),plateNum:this.newCar.plate,issuedAt:this.newCar.issuedAt,color:this.newCar.color,type:this.newCar.type,brand:this.newCar.brand,build:this.newCar.build});this.axios.post(this.$backendApi+"customer/add-car",e,{headers:{Authorization:sessionStorage.getItem("jwt")}}).then((e=>{console.log(e.data),this.getCarList()}))}}},ue=t(89);const de=(0,ue.Z)(se,[["render",ie]]);var ce=de,pe={name:"CallMechView",components:{CallMech:ce}};const me=(0,ue.Z)(pe,[["render",o]]);var ve=me}}]);
//# sourceMappingURL=600.447c32c6.js.map