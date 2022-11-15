<template>
  <div class="home">
    <HomePage/>
    {{ text }}
  </div>
</template>

<script>
// @ is an alias to /src
import HomePage from '@/components/HomePage.vue'

export default {
  name: 'HomeView',
  data () {
    return{
      text: ""
    }
  },
  mounted() {
    this.axios.get('http://localhost:3000/hello').then((response)=>{
      console.log(response)
      this.text = response.data
    })
  },
  components: {
    HomePage
  }
}
</script>
