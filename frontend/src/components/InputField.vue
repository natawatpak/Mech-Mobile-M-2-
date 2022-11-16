<template>
  <div class="input__container">
    <div :class="'input__effect ' + filled">
      <v-text-field
  :placeholder="placeholder" 
  @input="$emit('input',$event.target.value)"
  :class="'effect ' + has__icon"
  @focus="show=!show"
  @blur="show=!show"
  />  
      <label :style="{color:color}">{{label}}</label>
      <span class="focus__border" :style="focus__border"></span>
    </div>
    <span class="input__hint" v-if="show">{{hint}}</span>
    <i class="material-icons input__icon">{{icon}}</i>
  </div>
</template>

<script>
export default{
  data:()=>({
    show: false
  }),
  computed:{
    filled(){
      if ( !this.show && this.value ) {
        return 'has__content'
      }
      return ''
    },
    has__icon(){
      if ( this.icon ){
        return 'input__has__icon'
      }
      return ''
    },
    focus__border(){
      return {
        "background-color" : this.color
      }
    }
  },
  props: {
    value : { type: String , required: false, default: ''},
    label: { type: String , required: false, default: ''},
    hint: { type: String, required:false , default:'hint'},
    icon: { type: String, required: false , default: '' },
    placeholder: { type: String, required: false, default: ''},
    color: { type: String , required: false , default: 'indigo' }
  },
}
</script>